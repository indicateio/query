package query

import (
    "fmt"
    "strings"
)

func As(n, a string) string {
    return fmt.Sprintf("%s AS %s", n, a)
}
func QueryAs(query fmt.Stringer, alias string) string {
    return fmt.Sprintf("(%s) AS %s", query.String(), alias)
}

func Nested(a string) string {
    return fmt.Sprintf("(%s)", a)
}

func Expression(a string, o Operator, b string) string {
    return fmt.Sprintf("%s %s %s", a, o.String(), b)
}

func ExplainQuery(q fmt.Stringer) string {
    return fmt.Sprintf("EXPLAIN %s", q.String())
}

func Like(field string, value interface{}) string {
    return fmt.Sprintf("%s LIKE %v", field, value)
}

//func  NotLike(field string, value interface{}) string {
//    return fmt.Sprintf("%s NOT LIKE %s", field, c.Args.Add(value))
//}

func Not(expression string) string {
    return fmt.Sprintf("NOT %s", expression)
}

func IsNull(field string) string {
    return fmt.Sprintf("%s IS NULL", field)
}

func IsNotNull(field string) string {
    return fmt.Sprintf("%s IS NOT NULL", field)
}

func Max(field string) string {
    return fmt.Sprintf("MAX(%s)", field)
}

func Min(field string) string {
    return fmt.Sprintf("MIN(%s)", field)
}

func Count(fields ...string) string {
    return fmt.Sprintf("COUNT(%s)", strings.Join(fields, ", "))
}

func CountDistinct(fields ...string) string {
    return fmt.Sprintf("COUNT(DISTINCT %s)", strings.Join(fields, ", "))
}

func Coalesce(fields ...string) string {
    return fmt.Sprintf("COALESCE(%s)", strings.Join(fields, ", "))
}

func Sum(field string) string {
    return fmt.Sprintf("SUM(%s)", field)
}

func SumDistinct(field string) string {
    return fmt.Sprintf("SUM(DISTINCT %s)", field)
}

func Avg(field string) string {
    return fmt.Sprintf("AVG(%s)", field)
}

func AvgDistinct(field string) string {
    return fmt.Sprintf("AVG(DISTINCT %s)", field)
}

func Concat(fields ...string) string {
    // TODO e.g. space needs to be escaped
    return fmt.Sprintf("CONCAT(%s)", strings.Join(fields, ", "))
}

func Exists(query fmt.Stringer) string {
    return fmt.Sprintf("EXISTS(%s)", query.String())
}

func NotExists(query fmt.Stringer) string {
    return fmt.Sprintf("NOT EXISTS(%s)", query.String())
}

func InQuery(query fmt.Stringer) string {
    return fmt.Sprintf("IN(%s)", query.String())
}

func NotInQuery(query fmt.Stringer) string {
    return fmt.Sprintf("NOT IN(%s)", query.String())
}
func Values(field string) string {
    return fmt.Sprintf("VALUES(%s)", field)
}

func Equal(field string, value string) string {
    return Expression(field, OperatorEq, value)
}

func NotEqual(field string, value string) string {
    return Expression(field, OperatorNotEq, value)
}

func GreaterThan(field string, value string) string {
    return Expression(field, OperatorGt, value)
}

func GreaterEqualThan(field string, value string) string {
    return Expression(field, OperatorGtEq, value)
}

func LessThan(field string, value string) string {
    return Expression(field, OperatorLt, value)
}

func LessEqualThan(field string, value string) string {
    return Expression(field, OperatorLtEq, value)
}

func AnyQuery(query fmt.Stringer) string {
    return fmt.Sprintf("ANY (%s)", query.String())
}

func AllQuery(query fmt.Stringer) string {
    return fmt.Sprintf("ALL (%s)", query.String())
}

func Between(field string, a interface{}, b interface{}) string {
    return fmt.Sprintf("%s BETWEEN %s AND %s", field, a, b)
}

func NotBetween(field string, a interface{}, b interface{}) string {
    return fmt.Sprintf("%s NOT BETWEEN %s AND %s", field, a, b)
}

func In(field string, values ...interface{}) string {
    tmp := make([]string, 0, len(values))

    for _, v := range values {
        tmp = append(tmp, fmt.Sprintf("%v", v))
    }

    return fmt.Sprintf("%s IN (%s)", field, strings.Join(tmp, ", "))
}

func NotIn(field string, values ...interface{}) string {
    tmp := make([]string, 0, len(values))

    for _, v := range values {
        tmp = append(tmp, fmt.Sprintf("%v", v))
    }

    return fmt.Sprintf("%s NOT IN (%s)", field, strings.Join(tmp, ", "))
}
