package query

import (
    "fmt"
    "strings"
)

type Join struct {
    join        JoinType
    table       string
    expressions []string
}

func NewJoin(j JoinType, t string, p ...string) Join {
    return Join{
        join:        j,
        table:       t,
        expressions: p,
    }
}

func (p Join) String() string {
    return fmt.Sprintf(
        "%s %s ON %s",
        p.join.String(),
        p.table,
        strings.Join(p.expressions, fmt.Sprintf(" %s ", CombineAnd.String())),
    )
}
