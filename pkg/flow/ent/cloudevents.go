// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/direktiv/direktiv/pkg/flow/ent/cloudevents"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/google/uuid"
)

// CloudEvents is the model entity for the CloudEvents schema.
type CloudEvents struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id"`
	// EventId holds the value of the "eventId" field.
	EventId string `json:"eventId,omitempty"`
	// Event holds the value of the "event" field.
	Event event.Event `json:"event,omitempty"`
	// Fire holds the value of the "fire" field.
	Fire time.Time `json:"fire,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Processed holds the value of the "processed" field.
	Processed bool `json:"processed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CloudEventsQuery when eager-loading is set.
	Edges                 CloudEventsEdges `json:"edges"`
	namespace_cloudevents *uuid.UUID
}

// CloudEventsEdges holds the relations/edges for other nodes in the graph.
type CloudEventsEdges struct {
	// Namespace holds the value of the namespace edge.
	Namespace *Namespace `json:"namespace,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// NamespaceOrErr returns the Namespace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CloudEventsEdges) NamespaceOrErr() (*Namespace, error) {
	if e.loadedTypes[0] {
		if e.Namespace == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: namespace.Label}
		}
		return e.Namespace, nil
	}
	return nil, &NotLoadedError{edge: "namespace"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CloudEvents) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cloudevents.FieldEvent:
			values[i] = new([]byte)
		case cloudevents.FieldProcessed:
			values[i] = new(sql.NullBool)
		case cloudevents.FieldEventId:
			values[i] = new(sql.NullString)
		case cloudevents.FieldFire, cloudevents.FieldCreated:
			values[i] = new(sql.NullTime)
		case cloudevents.FieldID:
			values[i] = new(uuid.UUID)
		case cloudevents.ForeignKeys[0]: // namespace_cloudevents
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type CloudEvents", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CloudEvents fields.
func (ce *CloudEvents) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cloudevents.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ce.ID = *value
			}
		case cloudevents.FieldEventId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field eventId", values[i])
			} else if value.Valid {
				ce.EventId = value.String
			}
		case cloudevents.FieldEvent:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field event", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ce.Event); err != nil {
					return fmt.Errorf("unmarshal field event: %w", err)
				}
			}
		case cloudevents.FieldFire:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field fire", values[i])
			} else if value.Valid {
				ce.Fire = value.Time
			}
		case cloudevents.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				ce.Created = value.Time
			}
		case cloudevents.FieldProcessed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field processed", values[i])
			} else if value.Valid {
				ce.Processed = value.Bool
			}
		case cloudevents.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field namespace_cloudevents", values[i])
			} else if value.Valid {
				ce.namespace_cloudevents = new(uuid.UUID)
				*ce.namespace_cloudevents = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryNamespace queries the "namespace" edge of the CloudEvents entity.
func (ce *CloudEvents) QueryNamespace() *NamespaceQuery {
	return (&CloudEventsClient{config: ce.config}).QueryNamespace(ce)
}

// Update returns a builder for updating this CloudEvents.
// Note that you need to call CloudEvents.Unwrap() before calling this method if this CloudEvents
// was returned from a transaction, and the transaction was committed or rolled back.
func (ce *CloudEvents) Update() *CloudEventsUpdateOne {
	return (&CloudEventsClient{config: ce.config}).UpdateOne(ce)
}

// Unwrap unwraps the CloudEvents entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ce *CloudEvents) Unwrap() *CloudEvents {
	_tx, ok := ce.config.driver.(*txDriver)
	if !ok {
		panic("ent: CloudEvents is not a transactional entity")
	}
	ce.config.driver = _tx.drv
	return ce
}

// String implements the fmt.Stringer.
func (ce *CloudEvents) String() string {
	var builder strings.Builder
	builder.WriteString("CloudEvents(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ce.ID))
	builder.WriteString("eventId=")
	builder.WriteString(ce.EventId)
	builder.WriteString(", ")
	builder.WriteString("event=")
	builder.WriteString(fmt.Sprintf("%v", ce.Event))
	builder.WriteString(", ")
	builder.WriteString("fire=")
	builder.WriteString(ce.Fire.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(ce.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("processed=")
	builder.WriteString(fmt.Sprintf("%v", ce.Processed))
	builder.WriteByte(')')
	return builder.String()
}

// CloudEventsSlice is a parsable slice of CloudEvents.
type CloudEventsSlice []*CloudEvents

func (ce CloudEventsSlice) config(cfg config) {
	for _i := range ce {
		ce[_i].config = cfg
	}
}
