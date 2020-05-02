package query

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestInsertComplex(t *testing.T) {
    // TODO implement with select
}

func TestInsertMultiple(t *testing.T) {
    expectedSql :=
        "INSERT INTO table_name_a (uid, given_name, family_name, display_name, primary_email, created_date_time) VALUES (?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE uid = VALUES(uid), primary_email = VALUES(primary_email)"
    assert.Equal(t, expectedSql, testInsertMultiple())
}

func TestInsertSimple(t *testing.T) {
    expectedSql :=
        "INSERT IGNORE INTO table_name_a (uid, given_name, family_name, display_name, primary_email, created_date_time) VALUES (?, ?, ?, ?, ?, ?)"
    assert.Equal(t, expectedSql, testInsertSimple())
}

func BenchmarkInsertMultiple(b *testing.B) {
    for k := 0.; k <= 5; k++ {
        b.Run("testInsertMultiple", func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                testInsertMultiple()
            }
        })
    }
}

func BenchmarkInsertSimple(b *testing.B) {
    for k := 0.; k <= 5; k++ {
        b.Run("testInsertSimple", func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                testInsertSimple()
            }
        })
    }
}

func testInsertSimple() string {
    q := NewInsert("table_name_a")
    q.Ignore(true)
    insertCols := []string{
        "uid",
        "given_name",
        "family_name",
        "display_name",
        "primary_email",
        "created_date_time",
    }
    q.Columns(true, insertCols...)
    q.Duplicates(
        true,
        Equal("uid", PlaceholderQMark.String()),
        Equal("primary_email", PlaceholderQMark.String()),
    )
    q.ValuesPrepared(1)
    return q.String()
}
func testInsertMultiple() string {
    q := NewInsert("table_name_a")
    q.Ignore(false)
    insertCols := []string{
        "uid",
        "given_name",
        "family_name",
        "display_name",
        "primary_email",
        "created_date_time",
    }
    q.Columns(true, insertCols...)
    q.Duplicates(
        true,
        Expression("uid", OperatorEq, Values("uid")),
        Expression("primary_email", OperatorEq, Values("primary_email")),
    )
    q.ValuesPrepared(2)
    return q.String()
}
