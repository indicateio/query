<p>
<img alt="relation5" src="https://cdn.relation5.xyz/media/ci/logo/SVG/logo-relation5-horiz-claim-dark-colored.svg" width="256" />
</p>

# Relation5 - Query Builder
___
[![GoDoc Widget]][GoDoc]
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/relation5/query/master/LICENSE)
[![Build Status](https://travis-ci.org/relation5/query.svg?branch=master)](https://travis-ci.org/relation5/query)
[![relation5](https://img.shields.io/badge/relation5-awesome-%23ff4d58?style=flat-square&logo=data:image/png;base64,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)](https://gitlab.com/relation5/)

## Hints

##### This is not an ORM.
but a really simple, yet powerful query builder, that allows you to dynamically build complex queries fast and reliable.

##### Not thread safe
... not yet
##### Currently only for MySQL-queries
... postgress and Cassandra
##### not production-ready
this package is not complete and not fully tested. There may be breaking changes until the final release, so use it at your own risk :)

## Motivation

tbd.

## Install

### Requirements
Requires Go `1.10` and higher

Install with:
`go get -u github.com/relation5/query`

## Inspired by

- [squirrel](https://github.com/lann/squirrel)
- [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder)

## Usage
**As easy as:**

```go
import "github.com/relation5/query"
```

###  Example

Usage in combination with the great [sqlx](https://github.com/jmoiron/sqlx) package

```go
package user

import (
    "fmt"
    "github.com/relation5/query"
    "sort"
    "strconv"
)

func Slice(params SliceParams) {
    // [...]
    selectQuery := query.NewSelect(query.As("accounts", "a"))

    selectQuery.Columns(
        true,
        "a.uid",
        "a.given_name",
        "a.family_name",
        "a.display_name",
    )

    var (
        conditions []string
        arguments  []interface{}
    )

    if params.GroupUid != "" {
        selectQuery.Join(
            true,
            query.NewJoin(
                query.JoinTypeInner,
                query.As("group_members", "cm"),
                query.Expression("a.uid", query.OperatorEq, "cm.account_uid"),
            ),
        )
        conditions = append(conditions, query.Equal("cm.group_uid", query.PlaceholderQMark.String()))
        arguments = append(arguments, params.GroupUid)
    }

    if params.DomainUid != "" {
        selectQuery.Join(
            true,
            query.NewJoin(
                query.JoinTypeInner,
                query.As("domain_accounts", "da"),
                query.Expression("a.uid", query.OperatorEq, "da.account_uid"),
            ),
        )
        conditions = append(conditions, query.Equal("da.domain_uid", query.PlaceholderQMark.String()))
        arguments = append(arguments, params.DomainUid)
    }

    selectQuery.Where(
        query.NewPredicate(
            query.CombineAnd,
            false,
            conditions...
        ),
    )
    selectQuery.Distinct(true)
    selectQuery.OrderBy(true, query.DirectionDesc, "a.uid")

    if params.Limit > 0 {
        selectQuery.Limit(query.PlaceholderQMark.String())
        arguments = append(arguments, params.Limit)
    }

    rows, sqlErr := Database.Queryx(
        selectQuery.String(),
        arguments...,
    )

    defer func() {
        if rows != nil {
            if err := rows.Close(); err != nil {
                Logger.Error(err.Error())
            }
        }
    }()
    // [...]
}

```


### select
```go
package main

import (
    "fmt"
    "github.com/relation5/query"
)

func main() {
    q := query.NewSelect(query.As("table_name_a", "ta"))
    
    q.Columns(
        false,
        "ta.uid",
        "ta.given_name",
        "ta.family_name",
        query.As(query.Concat("ta.given_name", "' ,'", "ta.family_name"), "formatted_name"),
        "ta.primary_email",
    )
    q.Where(
        query.NewPredicate(
            query.CombineAnd,
            false,
            query.GreaterThan("ta.created_date_time", query.PlaceholderQMark.String()),
        ),
    )
    q.Limit(10)
   
    fmt.Println(q)
}
``` 
outputs: 
```mysql
SELECT 
    ta.uid, 
    ta.given_name, 
    ta.family_name, 
    CONCAT(ta.given_name, ' ,', ta.family_name) AS formatted_name, 
    ta.primary_email
FROM table_name_a as ta
WHERE ta.created_date_time > ?
LIMIT 1
```
### insert
```go
package main

import (
    "fmt"
    "github.com/relation5/query"
)

func main() {
    q := query.NewInsert("table_name_a")
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
        query.Equal("uid", query.PlaceholderQMark.String()),
        query.Equal("primary_email", query.PlaceholderQMark.String()),
    )
    q.ValuesPrepared(1)

    fmt.Println(q)
}
``` 
outputs: 
```mysql
INSERT INTO table_name_a (
    uid, 
    given_name,
    family_name, 
    display_name, 
    primary_email, 
    created_date_time
) VALUES (?, ?, ?, ?, ?, ?)
```

### update
```go
package main

import (
    "fmt"
    "github.com/relation5/query"
)

func main() {
    q := query.NewUpdate("table_name_a")
    q.Set(
        true,
        q.Add("login_count", 2),
        q.Assign("display_name", query.PlaceholderQMark.String()),
    )
    where := query.NewPredicate(
        query.CombineAnd,
        false,
        query.LessThan("created_date_time", query.PlaceholderQMark.String()),
    )

    q.Where(where)

    fmt.Println(q)
}
``` 
outputs: 
```mysql
UPDATE 
    table_name_a 
SET 
    login_count = login_count + 2, 
    display_name = ? 
WHERE created_date_time < ?
```

### delete
```go
package main

import (
    "fmt"
    "github.com/relation5/query"
)

func main() {
    q := query.NewDelete("table_name_a")
    where := query.NewPredicate(
        query.CombineAnd,
        false,
        query.LessThan("created_date_time", query.PlaceholderQMark.String()),
    )

    q.Where(where)
   
    fmt.Println(q)
}
``` 
outputs: 
```mysql
DELETE 
    FROM table_name_a
WHERE created_date_time < ?
```


### ...fun stuff :)
```go
package main

import (
    "fmt"
    "github.com/relation5/query"
)

func main() {
    q := query.NewInsert("table_name_a")
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
        query.Equal("uid",  query.Values("uid")),
        query.Equal("primary_email",  query.Values("primary_email")),
    )
    q.ValuesPrepared(2)
   
    fmt.Println(q)
}
``` 
outputs: 
```mysql
INSERT INTO table_name_a (
    uid, 
    given_name, 
    family_name, 
    display_name, 
    primary_email, 
    created_date_time
) 
VALUES 
    (?, ?, ?, ?, ?, ?), 
    (?, ?, ?, ?, ?, ?) 
ON DUPLICATE KEY 
UPDATE 
    uid = VALUES(uid), 
    primary_email = VALUES(primary_email)
```
#### ...or
```go
package main

import (
    "fmt"
    "github.com/relation5/query"
)

func main() {
    q := query.NewSelect(query.As("table_name_a", "ta"))
    
    q.Columns(
        false,
        "ta.uid",
        "ta.given_name",
        "ta.family_name",
        query.As(query.Concat("ta.given_name", "' ,'", "ta.family_name"), "formatted_name"),
        "ta.primary_email",
        "ta.primary_email",

    )
    q.Distinct(true)    

    subSelect := query.NewSelect(query.As("table_name_a", "tc"))
    subSelect.Columns(false, "tc.uid")
    subSelect.Where(
        query.NewPredicate(
            query.CombineAnd,
            false,
            query.GreaterThan("tc.created_date_time", query.PlaceholderQMark.String()),
        ),
    )
    subSelect.Limit(10)

    q.Join(
        true,
        query.NewJoin(
            query.JoinTypeInner,
            query.As("table_name_b", "tb"),
            query.Equal("ta.uid", "tb.table_a_uid"),
            query.NotEqual("tb.parent_uid", "ta.uid"),
        ),
        query.NewJoin(
            query.JoinTypeLeftOuter,
            query.As("table_name_c", "tc"),
            query.Equal("tb.uid", "tc.table_b_uid"),
        ),
        query.NewJoin(
            query.JoinTypeInner,
            query.QueryAs(subSelect, "filter"),
            query.Equal("filter.uid", "ta.uid"),

        ),
    )

    q.Where(
        query.NewPredicate(
            query.CombineAnd,
            false,
            query.Equal("da.account_uid",   query.PlaceholderQMark.String()),
        ),
    )
    q.OrderBy(true, query.DirectionDesc, "ta.given_name", "ta.family_name")
    q.OrderBy(true, query.DirectionAsc, "ta.display_name", "ta.formatted_name")
    q.Limit(query.PlaceholderQMark.String())
    q.Offset(query.PlaceholderQMark.String())
    q.GroupBy("ta.uid")
}
``` 
outputs: 
```mysql
SELECT DISTINCT 
    ta.uid, 
    ta.given_name, 
    ta.family_name, 
    CONCAT(ta.given_name, ' ,', ta.family_name) AS formatted_name, 
    ta.primary_email, 
    ta.primary_email 
FROM table_name_a AS ta 
INNER JOIN table_name_b AS tb 
    ON ta.uid = tb.table_a_uid 
    AND tb.parent_uid <> ta.uid
LEFT JOIN OUTER table_name_c AS tc 
    ON tb.uid = tc.table_b_uid 
INNER JOIN (
    SELECT 
        tc.uid 
    FROM table_name_a AS tc 
    WHERE tc.created_date_time > ? 
    LIMIT 10
    ) AS filter ON filter.uid = ta.uid 
WHERE da.account_uid = ? 
ORDER BY 
    ta.given_name, 
    ta.family_name DESC, 
    ta.display_name, 
    ta.formatted_name ASC 
GROUP BY ta.uid 
LIMIT ? 
OFFSET ?
```

...expect more :)

## Expressions

the `query` package comes equipped with an optional `expression` set, providing a suite of standard
MySQL expressions. Every Expression generates a part of the SQL and can completely be replaced by a standard string 

### built-ins

---------------------------------------------------------------------------------------------------------------------
| func                  | results in                                                                                |
|:----------------------|:-------------------------------------------------------------------------------------------
| `query.As(n, a string)`                                               | [...] AS [...]                            |
| `query.QueryAs(query fmt.Stringer, alias string)`                     | (SELECT FROM [...] ) AS sub               |
| `query.Nested(a string)`                                              | ([...])                                   |
| `query.Expression(a string, o Operator, b string)`                    | [...] [OPERATOR] [...]                    |
| `query.ExplainQuery(q fmt.Stringer)`                                  | EXPLAIN [...]                             |
| `query.Like(field string, value interface{})`                         | field LIKE %a                             |
| `query.Not(expression string)`                                        | NOT [...]                                 |
| `query.IsNull(field string)`                                          | [...] IS NULL                             |
| `query.IsNotNull(field string)`                                       | [...] IS NOT NULL                         |
| `query.Max(field string)`                                             | MAX([...])                                |
| `query.Min(field string)`                                             | MIN([...])                                |
| `query.Count(fields ...string)`                                       | COUNT([...,[...]])                        |
| `query.CountDistinct(fields ...string) `                              | COUNT(DISTINCT [...,[...]])               |
| `query.Coalesce(fields ...string)`                                    | COALESCE([...,[...]])                     |
| `query.Sum(field string)`                                             | SUM([...])                                |
| `query.SumDistinct(field string)`                                     | SUM(DISTINCT [...])                       |
| `query.Avg(field string)`                                             | AVG([...])                                |
| `query.AvgDistinct(field string)`                                     | AVG(DISTINCT [...])                       |
| `query.Concat(fields ...string)`                                      | CONCAT([...,[...]])                       |
| `query.Exists(query fmt.Stringer)`                                    | EXISTS([...])                             |
| `query.NotExists(query fmt.Stringer)`                                 | NOT EXISTS([...])                         |
| `query.InQuery(query fmt.Stringer)`                                   | IN([...])                                 |
| `query.NotInQuery(query fmt.Stringer)`                                | NOT IN([...])                             |
| `query.Values(field string)`                                          | VALUES([...])                             |
| `query.Equal(field string, value string)`                             | [...]  = [...]                            |
| `query.NotEqual(field string, value string)`                          | [...] <> [...]                            |
| `query.GreaterThan(field string, value string)`                       | [...]  > [...]                            |
| `query.GreaterEqualThan(field string, value string)`                  | [...] >= [...]                            |
| `query.LessThan(field string, value string)`                          | [...] < [...]                             |
| `query.LessEqualThan(field string, value string)`                     | [...] =< [...]                            |
| `query.AnyQuery(query fmt.Stringer)`                                  | ALL ([...])                               |
| `query.AllQuery(query fmt.Stringer)`                                  | ANY ([...])                               |
| `query.Between(field string, a interface{}, b interface{})`           | [...] BETWEEN [...] AND [...]             |
| `query.NotBetween(field string, a interface{}, b interface{})`        | [...] NOT BETWEEN [...] AND [...]         |
| `query.In(field string, values ...interface{})`                       | IN([...,[...]])                           |
| `query.NotIn(field string, values ...interface{})`                    | NOT IN([...,[...]])                       |
---------------------------------------------------------------------------------------------------------------------

## Upcoming features
- error handling
- other SQL-like queries
- detailed testing
- docs
- better placeholder handling
- named statements
- ... any idea? tell us!

## LICENSE

relation5 query is released under the
[MIT License](http://www.opensource.org/licenses/MIT).


[GoDoc]: https://pkg.go.dev/github.com/relation5/query
[GoDoc Widget]: https://godoc.org/github.com/relation5/query?status.svg
