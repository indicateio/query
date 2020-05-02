package query

import (
    "fmt"
    "strings"
)

type SelectQuery struct {
    distinct      bool
    limit         string
    offset        string
    table         string
    defaultSelect string
    columns       []string
    joins         []Join
    using         []string
    orderBy       []string
    groupBy       []string
    having        Predicate
    where         Predicate
    union         []SelectQuery
    placeholder   Placeholder
}

func NewSelect(t string) SelectQuery {
    return SelectQuery{
        table:         t,
        defaultSelect: "1",
    }
}

func (q SelectQuery) String() string {
    builder := strings.Builder{}
    builder.WriteString("SELECT ")

    if q.distinct {
        builder.WriteString("DISTINCT ")
    }
    // TODO maybe throw error if columns empty
    if len(q.columns) > 0 {
        builder.WriteString(fmt.Sprintf("%s ", strings.Join(q.columns, ", ")))
    } else {
        builder.WriteString(fmt.Sprintf("%s ", q.defaultSelect))
    }
    builder.WriteString(fmt.Sprintf("FROM %s ", q.table))

    if len(q.joins) > 0 {
        for _, j := range q.joins {
            builder.WriteString(fmt.Sprintf("%s ", j.String()))
        }
    }
    if len(q.using) > 0 {
        builder.WriteString(fmt.Sprintf("USING (%s) ", strings.Join(q.using, ", ")))
    }

    if w := q.where.String(); w != "" {
        builder.WriteString(fmt.Sprintf("WHERE %s ", w))
    }

    if len(q.orderBy) > 0 {
        builder.WriteString(fmt.Sprintf("ORDER BY %s ", strings.Join(q.orderBy, ", ")))
    }

    if len(q.groupBy) > 0 {
        builder.WriteString(fmt.Sprintf("GROUP BY %s ", strings.Join(q.groupBy, ", ")))

        if h := q.having.String(); h != "" {
            builder.WriteString(fmt.Sprintf("HAVING %s ", h))
        }
    }

    if q.limit != "" {
        builder.WriteString(fmt.Sprintf("LIMIT %s ", q.limit))
    }

    if q.offset != "" {
        builder.WriteString(fmt.Sprintf("OFFSET %s", q.offset))
    }

    if len(q.union) > 0 {
        for _, u := range q.union {
            builder.WriteString(fmt.Sprintf(" UNION %s ", u.String()))
        }
    }

    return strings.Trim(builder.String(), " ")
}

func (q *SelectQuery) Default(s string) {
    q.defaultSelect = s
}

func (q *SelectQuery) OrderBy(a bool, dir Direction, c ...string) {
    if !a {
        q.orderBy = make([]string, 0)
    }
    q.orderBy = append(q.orderBy, fmt.Sprintf("%s %s", strings.Join(c, ", "), dir.String()))
}

func (q *SelectQuery) GroupBy(c ...string) {
    q.groupBy = append(q.groupBy, c...)
}

func (q *SelectQuery) Columns(a bool, c ...string) {
    if a {
        q.columns = append(q.columns, c...)
        return
    }
    q.columns = c
}

func (q *SelectQuery) Using(a bool, u ...string) {
    if a {
        q.using = append(q.using, u...)
        return
    }
    q.using = u
}

func (q *SelectQuery) Distinct(d bool) {
    q.distinct = d
}

func (q *SelectQuery) Limit(v interface{}) {
    q.limit = fmt.Sprintf("%v", v)
}

func (q *SelectQuery) Offset(v interface{}) {
    q.offset = fmt.Sprintf("%v", v)
}

func (q *SelectQuery) Join(a bool, j ...Join) {
    if a {
        q.joins = append(q.joins, j...)
        return
    }
    q.joins = j
}

func (q *SelectQuery) Union(a bool, u ...SelectQuery) {
    // TODO implement UNION DISTINCT and ALL
    if a {
        q.union = append(q.union, u...)
        return
    }
    q.union = u
}

func (q *SelectQuery) Where(exp Predicate) {
    q.where = exp
}

func (q *SelectQuery) Having(exp Predicate) {
    q.having = exp
}
