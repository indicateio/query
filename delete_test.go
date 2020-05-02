package query

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestDeleteComplex(t *testing.T) {
    d := NewDelete(As("table_name_a", "ta"))
    d.Ignore(true)

    p1 := NewPredicate(
        CombineAnd,
        true,
        Expression("ta.uid", OperatorEq, PlaceholderQMark.String()),
        Expression("ta.name", OperatorEq, PlaceholderQMark.String()),
    )
    where := NewPredicate(
        CombineOr,
        false,
        p1.String(),
        Expression("ta.created_date_time", OperatorLt, PlaceholderQMark.String()),
    )

    d.Where(where)

    d.Limit(1)
    d.OrderBy(true, DirectionDesc, "ta.given_name", "ta.family_name")
    d.OrderBy(true, DirectionAsc, "ta.display_name", "ta.formatted_name")

    expectedSql :=
        "DELETE IGNORE FROM table_name_a AS ta WHERE (ta.uid = ? AND ta.name = ?) OR ta.created_date_time < ? ORDER BY ta.given_name, ta.family_name DESC, ta.display_name, ta.formatted_name ASC LIMIT 1"
    assert.Equal(t, expectedSql, d.String())
}
