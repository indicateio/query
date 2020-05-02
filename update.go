package query

import (
    "fmt"
    "strings"
)

type UpdateQuery struct {
    tables      []string
    ignore      bool
    assignments []string
    where       Predicate
    placeholder Placeholder
}

func NewUpdate(t ...string) UpdateQuery {
    return UpdateQuery{
        tables:      t,
        placeholder: PlaceholderQMark,
    }
}

func (q UpdateQuery) String() string {
    builder := strings.Builder{}
    builder.WriteString("UPDATE ")
    if q.ignore {
        builder.WriteString("IGNORE ")
    }
    builder.WriteString(fmt.Sprintf("%s ", strings.Join(q.tables, ", ")))
    builder.WriteString(fmt.Sprintf("SET %s ", strings.Join(q.assignments, ", ")))

    if w := q.where.String(); w != "" {
        builder.WriteString(fmt.Sprintf("WHERE %s ", w))
    }
    return strings.Trim(builder.String(), " ")
}

func (q *UpdateQuery) Where(exp Predicate) {
    q.where = exp
}

func (q *UpdateQuery) Ignore(d bool) {
    q.ignore = d
}

func (q *UpdateQuery) SetPrepared(a bool, fields ...string) {
    if !a {
        q.assignments = make([]string, 0, len(fields))
    }
    for _, f := range fields {
        q.assignments = append(q.assignments, q.Assign(f, q.placeholder.String()))
    }
}
func (q *UpdateQuery) Set(a bool, assignments ...string) {
    if a {
        q.assignments = append(q.assignments, assignments...)
        return
    }
    q.assignments = assignments
}

func (q *UpdateQuery) Assign(field string, value interface{}) string {
    return fmt.Sprintf("%s = %v", field, value)
}

func (q *UpdateQuery) Increment(field string) string {
    return q.Add(field, 1)
}

func (q *UpdateQuery) Decrement(field string) string {
    return q.Subtract(field, 1)
}

func (q *UpdateQuery) Add(field string, value interface{}) string {
    return fmt.Sprintf("%s = %s + %v", field, field, value)
}

func (q *UpdateQuery) Subtract(field string, value interface{}) string {
    return fmt.Sprintf("%s = %s - %v", field, field, value)
}

func (q *UpdateQuery) Multiply(field string, value interface{}) string {
    return fmt.Sprintf("%s = %s * %v", field, field, value)
}

func (q *UpdateQuery) Divide(field string, value interface{}) string {
    return fmt.Sprintf("%s = %s / %v", field, field, value)
}
