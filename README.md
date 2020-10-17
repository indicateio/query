# <img alt="relation5" src="https://cdn.relation5.de/media/logo/logo-relation5-horiz-colored.svg" width="256" />

[![GoDoc Widget]][GoDoc]
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/relation5/query/master/LICENSE)
[![Build Status](https://travis-ci.org/relation5/query.svg?branch=master)](https://travis-ci.org/relation5/query)

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
