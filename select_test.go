package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectComplex(t *testing.T) {
	expectedSql :=
		"SELECT ta.uid, ta.given_name, ta.family_name, CONCAT(ta.given_name, ' ,', ta.family_name) AS formatted_name, ta.primary_email, ta.primary_email FROM table_name_a AS ta INNER JOIN table_name_b AS tb ON ta.uid = tb.table_a_uid AND tb.parent_uid <> ta.uid LEFT OUTER JOIN table_name_c AS tc ON tb.uid = tc.table_b_uid INNER JOIN (SELECT tc.uid FROM table_name_a AS tc WHERE tc.created_date_time > ? LIMIT 10) AS filter ON filter.uid = ta.uid WHERE da.account_uid = ? UNION SELECT tc.uid FROM table_name_a AS tc WHERE tc.created_date_time > ? ORDER BY ta.given_name, ta.family_name DESC, ta.display_name, ta.formatted_name ASC GROUP BY ta.uid LIMIT ? OFFSET ?"
	assert.Equal(t, expectedSql, testSelectComplex())
}

func BenchmarkSelectComplex(b *testing.B) {
	for k := 0.; k <= 5; k++ {
		b.Run("testSelectComplex", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				testSelectComplex()
			}
		})
	}
}

func testSelectComplex() string {
	q := NewSelect(As("table_name_a", "ta"))

	q.Columns(
		false,
		"ta.uid",
		"ta.given_name",
		"ta.family_name",
		As(Concat("ta.given_name", "' ,'", "ta.family_name"), "formatted_name"),
		"ta.primary_email",
		"ta.primary_email",
	)

	subSelect := NewSelect(As("table_name_a", "tc"))
	subSelect.Columns(false, "tc.uid")
	subSelect.Where(
		NewPredicate(
			CombineAnd,
			false,
			GreaterThan("tc.created_date_time", PlaceholderQMark.String()),
		),
	)
	subSelect.Limit(10)

	unionSelect := NewSelect(As("table_name_a", "tc"))
	unionSelect.Columns(false, "tc.uid")
	unionSelect.Where(
		NewPredicate(
			CombineAnd,
			false,
			GreaterThan("tc.created_date_time", PlaceholderQMark.String()),
		),
	)
	q.Union(true, unionSelect)

	q.Join(
		true,
		NewJoin(
			JoinTypeInner,
			As("table_name_b", "tb"),
			Equal("ta.uid", "tb.table_a_uid"),
			NotEqual("tb.parent_uid", "ta.uid"),
		),
		NewJoin(
			JoinTypeLeftOuter,
			As("table_name_c", "tc"),
			Equal("tb.uid", "tc.table_b_uid"),
		),
		NewJoin(
			JoinTypeInner,
			QueryAs(subSelect, "filter"),
			Equal("filter.uid", "ta.uid"),
		),
	)

	q.Where(
		NewPredicate(
			CombineAnd,
			false,
			Expression("da.account_uid", OperatorEq, PlaceholderQMark.String()),
		),
	)
	q.OrderBy(true, DirectionDesc, "ta.given_name", "ta.family_name")
	q.OrderBy(true, DirectionAsc, "ta.display_name", "ta.formatted_name")
	q.Limit(PlaceholderQMark.String())
	q.Offset(PlaceholderQMark.String())
	q.GroupBy("ta.uid")

	return q.String()
}
