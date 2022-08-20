// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/contrib/entgql/internal/todo/ent/category"
	"entgo.io/contrib/entgql/internal/todo/ent/todo"
	"entgo.io/contrib/entgql/internal/todo/ent/verysecret"
	"entgo.io/ent/dialect/sql"
)

// Todo is the model entity for the Todo schema.
type Todo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Status holds the value of the "status" field.
	Status todo.Status `json:"status,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int `json:"priority,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Blob holds the value of the "blob" field.
	Blob []byte `json:"blob,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID int `json:"category_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TodoQuery when eager-loading is set.
	Edges         TodoEdges `json:"edges"`
	todo_children *int
	todo_secret   *int
}

// TodoEdges holds the relations/edges for other nodes in the graph.
type TodoEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Todo `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Todo `json:"children,omitempty"`
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// Secret holds the value of the secret edge.
	Secret *VerySecret `json:"secret,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedChildren map[string][]*Todo
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TodoEdges) ParentOrErr() (*Todo, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: todo.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e TodoEdges) ChildrenOrErr() ([]*Todo, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TodoEdges) CategoryOrErr() (*Category, error) {
	if e.loadedTypes[2] {
		if e.Category == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// SecretOrErr returns the Secret value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TodoEdges) SecretOrErr() (*VerySecret, error) {
	if e.loadedTypes[3] {
		if e.Secret == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: verysecret.Label}
		}
		return e.Secret, nil
	}
	return nil, &NotLoadedError{edge: "secret"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case todo.FieldBlob:
			values[i] = new([]byte)
		case todo.FieldID, todo.FieldPriority, todo.FieldCategoryID:
			values[i] = new(sql.NullInt64)
		case todo.FieldStatus, todo.FieldText:
			values[i] = new(sql.NullString)
		case todo.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case todo.ForeignKeys[0]: // todo_children
			values[i] = new(sql.NullInt64)
		case todo.ForeignKeys[1]: // todo_secret
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Todo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todo fields.
func (t *Todo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case todo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case todo.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = todo.Status(value.String)
			}
		case todo.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				t.Priority = int(value.Int64)
			}
		case todo.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				t.Text = value.String
			}
		case todo.FieldBlob:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field blob", values[i])
			} else if value != nil {
				t.Blob = *value
			}
		case todo.FieldCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value.Valid {
				t.CategoryID = int(value.Int64)
			}
		case todo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field todo_children", value)
			} else if value.Valid {
				t.todo_children = new(int)
				*t.todo_children = int(value.Int64)
			}
		case todo.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field todo_secret", value)
			} else if value.Valid {
				t.todo_secret = new(int)
				*t.todo_secret = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the Todo entity.
func (t *Todo) QueryParent() *TodoQuery {
	return (&TodoClient{config: t.config}).QueryParent(t)
}

// QueryChildren queries the "children" edge of the Todo entity.
func (t *Todo) QueryChildren() *TodoQuery {
	return (&TodoClient{config: t.config}).QueryChildren(t)
}

// QueryCategory queries the "category" edge of the Todo entity.
func (t *Todo) QueryCategory() *CategoryQuery {
	return (&TodoClient{config: t.config}).QueryCategory(t)
}

// QuerySecret queries the "secret" edge of the Todo entity.
func (t *Todo) QuerySecret() *VerySecretQuery {
	return (&TodoClient{config: t.config}).QuerySecret(t)
}

// Update returns a builder for updating this Todo.
// Note that you need to call Todo.Unwrap() before calling this method if this Todo
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todo) Update() *TodoUpdateOne {
	return (&TodoClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Todo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Todo) Unwrap() *Todo {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todo is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todo) String() string {
	var builder strings.Builder
	builder.WriteString("Todo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", t.Priority))
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(t.Text)
	builder.WriteString(", ")
	builder.WriteString("blob=")
	builder.WriteString(fmt.Sprintf("%v", t.Blob))
	builder.WriteString(", ")
	builder.WriteString("category_id=")
	builder.WriteString(fmt.Sprintf("%v", t.CategoryID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedChildren returns the Children named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Todo) NamedChildren(name string) ([]*Todo, error) {
	if t.Edges.namedChildren == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedChildren[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Todo) appendNamedChildren(name string, edges ...*Todo) {
	if t.Edges.namedChildren == nil {
		t.Edges.namedChildren = make(map[string][]*Todo)
	}
	if len(edges) == 0 {
		t.Edges.namedChildren[name] = []*Todo{}
	} else {
		t.Edges.namedChildren[name] = append(t.Edges.namedChildren[name], edges...)
	}
}

// Todos is a parsable slice of Todo.
type Todos []*Todo

func (t Todos) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
