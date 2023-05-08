package sql

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/direktiv/direktiv/pkg/refactor/logengine"
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

var _ logengine.LogStore = &sqlLogStore{} // Ensures SQLLogStore struct conforms to logengine.LogStore interface.

type sqlLogStore struct {
	db *gorm.DB
}

// Append implements logengine.LogStore.
// - For instance-logs following Key Value pairs SHOULD be present: instance-id, callpath, root-instance-id
// - For namespace-logs following Key Value pairs SHOULD be present: namespace-id
// - For mirror-logs following Key Value pairs SHOULD be present: mirror-id
// - For workflow-logs following Key Value pairs SHOULD be present: workflow-id
// - All passed keysAndValues pair will be stored attached to the log-entry.
func (sl *sqlLogStore) Append(ctx context.Context, timestamp time.Time, msg string, keysAndValues ...interface{}) error {
	fields, err := mapKeysAndValues(keysAndValues...)
	cols := make([]string, 0, len(keysAndValues))
	vals := make([]interface{}, 0, len(keysAndValues))
	if err != nil {
		return err
	}
	msg = strings.ReplaceAll(msg, "\u0000", "") // postgres will return an error if a string contains "\u0000"
	cols = append(cols, "oid", "t", "msg")
	vals = append(vals, uuid.New(), timestamp, msg)
	lvl, ok := fields["level"]
	if ok {
		cols = append(cols, "level")
		vals = append(vals, lvl)
		delete(fields, "level")
	}
	value, ok := fields["instance-id"]
	if ok {
		cols = append(cols, "instance_logs")
		vals = append(vals, value)
	}
	value, ok = fields["callpath"]
	if ok {
		cols = append(cols, "log_instance_call_path")
		vals = append(vals, value)
		delete(fields, "callpath")
	}
	value, ok = fields["root-instance-id"]
	if ok {
		cols = append(cols, "root_instance_id")
		vals = append(vals, value)
		delete(fields, "root-instance-id")
	}
	value, ok = fields["workflow-id"]
	if ok {
		cols = append(cols, "workflow_id")
		vals = append(vals, value)
	}
	value, ok = fields["namespace-id"]
	if ok {
		cols = append(cols, "namespace_logs")
		vals = append(vals, value)
	}
	value, ok = fields["mirror-id"]
	if ok {
		cols = append(cols, "mirror_activity_id")
		vals = append(vals, value)
	}
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

// Get implements logengine.LogStore.
// - To query server-logs pass: "recipientType", "server" via keysAndValues
// - For pagination pass limit and | or offset. They MUST be passed as integer.
// - level MUST be passed as a string. Valid values are "debug", "info", "error", "panic".
// - This method will search for any of followings keys and query all matching logs:
// level, workflow-id, namespace-id, callpath, root-instance-id, mirror-id
// Any other not mentioned passed key value pair will be ignored.
// Returned log-entries will have same or higher level as the passed one.
// - Passing a callpath will return all logs which have a callpath with the prefix as the passed callpath value.
// when passing callpath the root-instance-id SHOULD be passed to optimize the performance of the query.
func (sl *sqlLogStore) Get(ctx context.Context, keysAndValues ...interface{}) ([]*logengine.LogEntry, error) {
	fields, err := mapKeysAndValues(keysAndValues...)
	if err != nil {
		return nil, err
	}
	wEq := []string{}
	if fields["recipientType"] == srv {
		wEq = append(wEq, "workflow_id IS NULL")
		wEq = append(wEq, "namespace_logs IS NULL")
		wEq = append(wEq, "instance_logs IS NULL")
	}
	level, ok := fields["level"]
	if ok {
		levelValue, ok := level.(string)
		if !ok {
			return nil, fmt.Errorf("level should be a string")
		}
		wEq = setMinumLogLevel(levelValue, wEq)
	}
	id, ok := fields["workflow-id"]
	if ok {
		wEq = append(wEq, fmt.Sprintf("workflow_id = '%s'", id))
	}
	id, ok = fields["namespace-id"]
	if ok {
		wEq = append(wEq, fmt.Sprintf("namespace_logs = '%s'", id))
	}
	prefix, ok := fields["callpath"]
	if ok {
		wEq = append(wEq, fmt.Sprintf("log_instance_call_path like '%s%%'", prefix))
	}
	id, ok = fields["root-instance-id"]
	if ok {
		wEq = append(wEq, fmt.Sprintf("root_instance_id = '%s'", id))
	}
	id, ok = fields["mirror-id"]
	if ok {
		wEq = append(wEq, fmt.Sprintf("mirror_activity_id = '%s'", id))
	}
	limit, ok := fields["limit"]
	limitValue := -1
	if ok {
		limitValue, ok = limit.(int)
		if !ok {
			return nil, fmt.Errorf("limit should be an int")
		}
	}
	offset, ok := fields["offset"]
	offsetValue := -1
	if ok {
		offsetValue, ok = offset.(int)
		if !ok {
			return nil, fmt.Errorf("offset should be an int")
		}
	}
	query := composeQuery(limitValue, offsetValue, wEq)

	resultList := make([]*gormLogMsg, 0)
	tx := sl.db.WithContext(ctx).Raw(query).Scan(&resultList)
	if tx.Error != nil {
		return nil, tx.Error
	}
	convertedList := make([]*logengine.LogEntry, 0, len(resultList))
	for _, e := range resultList {
		m := make(map[string]interface{})
		for k, e := range e.Tags {
			m[k] = e
		}
		m["level"] = e.Level
		convertedList = append(convertedList, &logengine.LogEntry{
			T:      e.T,
			Msg:    e.Msg,
			Fields: m,
		})
	}

	return convertedList, nil
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
	if limit >= 0 {
		q += fmt.Sprintf(" LIMIT %d ", limit)
	}
	if offset >= 0 {
		q += fmt.Sprintf(" OFFSET %d ", offset)
	}

	return q + ";"
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