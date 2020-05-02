package query

import (
    "fmt"
    "strings"
)

type InsertQuery struct {
    table       string
    ignore      bool
    columns     []string
    duplicates  []string
    values      [][]string
    placeholder Placeholder
}

func NewInsert(t string) InsertQuery {
    return InsertQuery{
        table:       t,
        placeholder: PlaceholderQMark,
    }
}

func (q InsertQuery) String() string {
    builder := strings.Builder{}
    builder.WriteString("INSERT ")

    if q.ignore {
        builder.WriteString("IGNORE ")
    }

    builder.WriteString(fmt.Sprintf("INTO %s ", q.table))

    builder.WriteString(fmt.Sprintf("(%s) ", strings.Join(q.columns, ", ")))

    valueCount := len(q.values)
    values := make([]string, 0, valueCount)
    for _, v := range q.values {
        values = append(values, fmt.Sprintf("(%s)", strings.Join(v, ", ")))
    }

    builder.WriteString(fmt.Sprintf("VALUES %s ", strings.Join(values, ", ")))

    if !q.ignore && len(q.duplicates) > 0 {
        builder.WriteString(fmt.Sprintf("ON DUPLICATE KEY UPDATE %s", strings.Join(q.duplicates, ", ")))
    }

    return strings.Trim(builder.String(), " ")
}

func (q *InsertQuery) Columns(a bool, c ...string) {
    if a {
        q.columns = append(q.columns, c...)
        return
    }
    q.columns = c
}

func (q *InsertQuery) Duplicates(a bool, d ...string) {
    if a {
        q.duplicates = append(q.duplicates, d...)
        return
    }
    q.duplicates = d
}

func (q *InsertQuery) Ignore(d bool) {
    q.ignore = d
}

func (q *InsertQuery) ValuesPrepared(rows int) {
    // TODO maybe throw error
    if count := len(q.columns); count > 0 {
        q.values = make([][]string, 0, rows)
        for i := 0; i < rows; i++ {
            q.values = append(q.values, strings.Split(strings.Repeat(q.placeholder.String(), count), ""))
        }
    }
}
func (q *InsertQuery) Values(a bool, values ...interface{}) error {
    if len(values) != len(q.columns) {
        return fmt.Errorf("values count does not match column count")
    }
    if !a {
        q.values = make([][]string, 0)
    }

    tmp := make([]string, 0)
    for _, v := range values {
        tmp = append(tmp, fmt.Sprintf("%v", v))
    }
    q.values = append(q.values, tmp)

    return nil
}
