package query

import (
    "fmt"
    "strings"
)

type Predicate struct {
    expressions []string
    combine     Combine
    nested      bool
}

func NewPredicate(c Combine, nested bool, ex ...string) Predicate {
    return Predicate{
        expressions: ex,
        nested:      nested,
        combine:     c,
    }
}

func (p Predicate) String() string {
    if p.nested {
        return fmt.Sprintf("(%s)", strings.Join(p.expressions, fmt.Sprintf(" %s ", p.combine.String())))
    }
    return strings.Join(p.expressions, fmt.Sprintf(" %s ", p.combine.String()))
}
