package query

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestUpdate(t *testing.T) {
    q := NewUpdate("table_name_a")
    q.Ignore(true)

    q.SetPrepared(
        true,
        "formatted_name",
    )
    q.Set(
        true,
        q.Add("login_count", 2),
        q.Assign("display_name", PlaceholderQMark.String()),
    )
    p1 := NewPredicate(
        CombineAnd,
        true,
        Expression("uid", OperatorEq, PlaceholderQMark.String()),
        Expression("name",OperatorEq, PlaceholderQMark.String()),
    )
    where := NewPredicate(
        CombineOr,
        false,
        p1.String(),
        Expression("created_date_time", OperatorLt, PlaceholderQMark.String()),
    )

    q.Where(where)

    expectedSql :=
        "UPDATE IGNORE table_name_a SET formatted_name = ?, login_count = login_count + 2, display_name = ? WHERE (uid = ? AND name = ?) OR created_date_time < ?"
    assert.Equal(t, expectedSql, q.String())
}
