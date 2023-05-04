package dblogstore

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/direktiv/direktiv/pkg/refactor/logstore"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ns  = "namespace"
	wf  = "workflow"
	srv = "server"
	ins = "instance"
	mir = "mirror"
)

var _ logstore.LogStore = &SQLLogStore{} // Ensures SQLLogStore struct conforms to logstore.LogStore interface.

func NewSQLLogStore(db *gorm.DB) logstore.LogStore {
	return &SQLLogStore{
		db: db,
	}
}

type SQLLogStore struct {
	db *gorm.DB
}

// Append implements logstore.LogStore.
func (sl *SQLLogStore) Append(ctx context.Context, timestamp time.Time, msg string, keysAndValues ...interface{}) error {
	fields, err := mapKeysAndValues(keysAndValues...)
	cols := make([]string, 0, len(keysAndValues))
	vals := make([]interface{}, 0, len(keysAndValues))
	if err != nil {
		return err
	}
	lvl, err := getLevel(fields)
	if err != nil {
		return err
	}
	cols = append(cols, "t", "msg", "level")
	vals = append(vals, timestamp, msg, lvl)
	recipientType, err := getRecipientType(fields)
	if err != nil {
		return err
	}
	switch recipientType {
	case ins:
		fieldValues, err := getFields(fields, "instance-id", "callpath", "rootInstanceID")
		if err != nil {
			return err
		}
		cols = append(cols, "instance_logs", "log_instance_call_path", "root_instance_id")
		vals = append(vals, fieldValues...)
	case wf:
		fieldValues, err := getFields(fields, "workflow-id")
		if err != nil {
			return err
		}
		cols = append(cols, "workflow_id")
		vals = append(vals, fieldValues...)
	case ns:
		fieldValues, err := getFields(fields, "namespace-id")
		if err != nil {
			return err
		}
		cols = append(cols, "namespace_logs")
		vals = append(vals, fieldValues...)
	case mir:
		fieldValues, err := getFields(fields, "mirror-id")
		if err != nil {
			return err
		}
		cols = append(cols, "mirror_activity_id")
		vals = append(vals, fieldValues...)
	case srv:
		// do nothing
	}
	delete(fields, "level")
	delete(fields, "rootInstanceID")
	delete(fields, "callpath")
	delete(fields, "recipientType")
	b, err := json.Marshal(fields)
	if err != nil {
		return err
	}
	cols = append(cols, "tags")
	vals = append(vals, b)
	q := "INSERT INTO log_msgs ("
	qTail := "VALUES ("
	for i := range vals {
		q += fmt.Sprintf("'%s'", cols[i])
		qTail += fmt.Sprintf("$%d", i+1)
		if i != len(vals)-1 {
			q += ", "
			qTail += ", "
		}
	}
	q = q + ") " + qTail + ");"
	tx := sl.db.WithContext(ctx).Exec(q, vals...)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func getFields(fields map[string]interface{}, fieldNames ...string) ([]interface{}, error) {
	res := []interface{}{}
	for _, v := range fieldNames {
		field, ok := fields[v]
		if !ok {
			return nil, fmt.Errorf("%s was missing as argument in the keysAndValues pair", v)
		}
		res = append(res, field)
	}

	return res, nil
}

// Get implements logstore.LogStore.
func (sl *SQLLogStore) Get(ctx context.Context, keysAndValues ...interface{}) ([]*logstore.LogEntry, error) {
	fields, err := mapKeysAndValues(keysAndValues...)
	if err != nil {
		return nil, err
	}
	query, err := buildQuery(fields)
	if err != nil {
		return nil, err
	}
	resultList := make([]*gormLogMsg, 0)
	tx := sl.db.WithContext(ctx).Raw(query).Scan(&resultList)
	if tx.Error != nil {
		return nil, tx.Error
	}
	convertedList := make([]*logstore.LogEntry, 0, len(resultList))
	for _, e := range resultList {
		m := make(map[string]interface{})
		for k, e := range e.Tags {
			m[k] = e
		}
		m["level"] = e.Level
		convertedList = append(convertedList, &logstore.LogEntry{
			T:      e.T,
			Msg:    e.Msg,
			Fields: m,
		})
	}

	return convertedList, nil
}

func buildQuery(fields map[string]interface{}) (string, error) {
	wEq := []string{}
	recipientType, err := getRecipientType(fields)
	if err != nil {
		return "", err
	}
	level, ok := fields["level"]
	if ok {
		levelValue, ok := level.(string)
		if !ok {
			return "", fmt.Errorf("level should be a string")
		}
		wEq = setMinumLogLevel(levelValue, wEq)
	}

	if recipientType == srv {
		wEq = append(wEq, "workflow_id IS NULL")
		wEq = append(wEq, "namespace_logs IS NULL")
		wEq = append(wEq, "instance_logs IS NULL")
	}
	if recipientType == wf {
		id, ok := fields["workflow-id"]
		if !ok {
			return "", fmt.Errorf("workflow-id is missing")
		}
		wEq = append(wEq, fmt.Sprintf("workflow_id = '%s'", id))
	}
	if recipientType == ns {
		id, ok := fields["namespace-id"]
		if !ok {
			return "", fmt.Errorf("namespace-id is missing")
		}
		wEq = append(wEq, fmt.Sprintf("namespace_logs = '%s'", id))
	}
	if recipientType == ins {
		var err error
		wEq, err = addInstanceValuesToQuery(wEq, fields)
		if err != nil {
			return "", err
		}
	}
	if recipientType == mir {
		id, ok := fields["mirror-activity-id"]
		if !ok {
			return "", fmt.Errorf("mirror-activity-id is missing")
		}
		wEq = append(wEq, fmt.Sprintf("mirror_activity_id = '%s'", id))
	}
	limit, ok := fields["limit"]
	var limitValue int
	if ok {
		limitValue, ok = limit.(int)
		if !ok {
			return "", fmt.Errorf("limit should be an int")
		}
	}
	offset, ok := fields["offset"]
	var offsetValue int
	if ok {
		offsetValue, ok = offset.(int)
		if !ok {
			return "", fmt.Errorf("offset should be an int")
		}
	}
	q := composeQuery(limitValue, offsetValue, wEq)

	return q, nil
}

func composeQuery(limit, offset int, wEq []string) string {
	q := `SELECT oid, t, msg, level, root_instance_id, log_instance_call_path, tags, workflow_id, mirror_activity_id, instance_logs, namespace_logs
	FROM log_msgs `
	q += "WHERE "
	for i, e := range wEq {
		q += e
		if i+1 < len(wEq) {
			q += " AND "
		}
	}
	q += " ORDER BY t ASC"
	if limit > 0 {
		q += fmt.Sprintf(" LIMIT %d", limit)
	}
	if offset > 0 {
		q += fmt.Sprintf(" OFFSET %d", offset)
	}

	return q
}

func getLevel(fields map[string]interface{}) (string, error) {
	lvl, ok := fields["level"]
	if !ok {
		return "", fmt.Errorf("level was missing as argument in the keysAndValues pair")
	}
	switch lvl {
	case "debug", "info", "error", "panic":
		return fmt.Sprintf("%s", lvl), nil
	}

	return "", fmt.Errorf("field level provided in keysAndValues has an invalid value %s", lvl)
}

func getRecipientType(fields map[string]interface{}) (string, error) {
	recipientType, ok := fields["recipientType"]
	if !ok {
		return "", fmt.Errorf("recipientType was missing as argument in the keysAndValues pair")
	}
	switch recipientType {
	case srv, ins, mir, ns, wf:
		return fmt.Sprintf("%s", recipientType), nil
	}

	return "", fmt.Errorf("invalid recipientType %s", recipientType)
}

func addInstanceValuesToQuery(wEq []string, fields map[string]interface{}) ([]string, error) {
	values, err := getFields(fields, "callpath", "rootInstanceID", "isCallerRoot")
	if err != nil {
		return nil, err
	}
	prefix := values[0]
	rootInstanceID := values[1]
	callerIsRootValue, ok := values[2].(bool)
	if !ok {
		return nil, fmt.Errorf("callerIsRoot should be an bool")
	}
	wEq = append(wEq, fmt.Sprintf("root_instance_id = '%s'", rootInstanceID))
	if !callerIsRootValue {
		wEq = append(wEq, fmt.Sprintf("log_instance_call_path like '%s%%'", prefix))
	}

	return wEq, nil
}

func mapKeysAndValues(keysAndValues ...interface{}) (map[string]interface{}, error) {
	fields := make(map[string]interface{})
	if len(keysAndValues) == 0 || len(keysAndValues)%2 != 0 {
		return nil, fmt.Errorf("keysAndValues was not a list of key value pairs")
	}
	for i := 0; i < len(keysAndValues); i += 2 {
		key := fmt.Sprintf("%s", keysAndValues[i])
		fields[key] = keysAndValues[i+1]
	}

	return fields, nil
}

type gormLogMsg struct {
	Oid                 uuid.UUID
	T                   time.Time
	Msg                 string
	Level               string
	Tags                map[string]interface{}
	WorkflowID          string
	MirrorActivityID    string
	InstanceLogs        string
	NamespaceLogs       string
	RootInstanceID      string
	LogInstanceCallPath string
}

func setMinumLogLevel(loglevel string, wEq []string) []string {
	levels := []string{"debug", "info", "error", "panic"}
	switch loglevel {
	case "debug":
		return wEq
	case "info":
		levels = levels[1:]
	case "error":
		levels = levels[2:]
	case "panic":
		levels = levels[3:]
	}
	q := "( "
	for i, e := range levels {
		q += fmt.Sprintf("level = '%s' ", e)
		if i < len(levels)-1 {
			q += " OR "
		}
	}
	q += " )"

	return append(wEq, q)
}
