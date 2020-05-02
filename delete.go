package query

import (
    "fmt"
    "strconv"
    "strings"
)

type DeleteQuery struct {
    tables      []string
    ignore      bool
    limit       int
    orderBy     []string
    where       Predicate
    placeholder Placeholder
}

func NewDelete(t ...string) DeleteQuery {
    return DeleteQuery{
        tables:      t,
        placeholder: PlaceholderQMark,
    }
}

func (q DeleteQuery) String() string {
    builder := strings.Builder{}
    builder.WriteString("DELETE ")
    if q.ignore {
        builder.WriteString("IGNORE ")
    }
    builder.WriteString(fmt.Sprintf("FROM %s ", strings.Join(q.tables, ", ")))

    if w := q.where.String(); w != "" {
        builder.WriteString(fmt.Sprintf("WHERE %s ", w))
    }

    if len(q.orderBy) > 0 {
        builder.WriteString(fmt.Sprintf("ORDER BY %s ", strings.Join(q.orderBy, ", ")))
    }
    if q.limit > 0 {
        builder.WriteString(fmt.Sprintf("LIMIT %s ", strconv.Itoa(q.limit)))
    }

    return strings.Trim(builder.String(), " ")
}

func (q *DeleteQuery) Ignore(d bool) {
    q.ignore = d
}

func (q *DeleteQuery) OrderBy(a bool, dir Direction, c ...string) {
    if !a {
        q.orderBy = make([]string, 0)
    }
    q.orderBy = append(q.orderBy, fmt.Sprintf("%s %s", strings.Join(c, ", "), dir.String()))
}

func (q *DeleteQuery) Limit(l int) {
    q.limit = l
}

func (q *DeleteQuery) Where(exp Predicate) {
    q.where = exp
}
