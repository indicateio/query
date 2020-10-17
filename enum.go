package query

type JoinType int

const (
    // use “iota + 1” to be sure that the enum type is initialized.
    JoinTypeBase JoinType = iota + 1
    JoinTypeLeft
    JoinTypeLeftOuter
    JoinTypeRight
    JoinTypeRightOuter
    JoinTypeInner
    JoinTypeFull
    JoinTypeFullOuter
    JoinTypeNatural
    JoinTypeCross
    endJoinType
)

var types = map[JoinType]string{
    JoinTypeBase:       "JOIN",
    JoinTypeLeft:       "LEFT JOIN",
    JoinTypeLeftOuter:  "LEFT OUTER JOIN",
    JoinTypeRight:      "RIGHT JOIN",
    JoinTypeRightOuter: "RIGHT OUTER JOIN",
    JoinTypeInner:      "INNER JOIN",
    JoinTypeFull:       "FULL JOIN",
    JoinTypeFullOuter:  "FULL OUTER JOIN",
    JoinTypeNatural:    "NATURAL JOIN",
    JoinTypeCross:      "CROSS JOIN",
}

func (e JoinType) IsValid() bool {
    return e >= JoinTypeBase && e < endJoinType
}

func (e JoinType) Ordinal() int {
    return int(e)
}

func (e JoinType) String() string {
    if !e.IsValid() {
        // TODO maybe return error
        return types[endJoinType]
    }

    return types[e]
}

type Operator int

const (
    // use “iota + 1” to be sure that the enum type is initialized.
    OperatorEq Operator = iota + 1
    OperatorNotEq
    OperatorGt
    OperatorGtEq
    OperatorLt
    OperatorLtEq
    endOperator
)

var operators = map[Operator]string{
    OperatorEq:    "=",
    OperatorNotEq: "<>",
    OperatorGt:    ">",
    OperatorGtEq:  ">=",
    OperatorLt:    "<",
    OperatorLtEq:  "<=",
}

func (e Operator) IsValid() bool {
    return e >= OperatorEq && e < endOperator
}

func (e Operator) Ordinal() int {
    return int(e)
}

func (e Operator) String() string {
    if !e.IsValid() {
        // TODO maybe return error
        return operators[endOperator]
    }

    return operators[e]
}

type Combine int

const (
    // use “iota + 1” to be sure that the enum type is initialized.
    CombineAnd Combine = iota + 1
    CombineOr
    endCombine
)

var combines = map[Combine]string{
    CombineAnd: "AND",
    CombineOr:  "OR",
}

func (e Combine) IsValid() bool {
    return e >= CombineAnd && e < endCombine
}

func (e Combine) Ordinal() int {
    return int(e)
}

func (e Combine) String() string {
    if !e.IsValid() {
        // TODO maybe return error
        return combines[endCombine]
    }

    return combines[e]
}

type Placeholder int

const (
    // use “iota + 1” to be sure that the enum type is initialized.
    PlaceholderQMark Placeholder = iota + 1
    PlaceholderDollar
    endPlaceholder
)

var placeholders = map[Placeholder]string{
    PlaceholderQMark:  "?",
    PlaceholderDollar: "$",
}

func (e Placeholder) IsValid() bool {
    return e >= PlaceholderQMark && e < endPlaceholder
}

func (e Placeholder) Ordinal() int {
    return int(e)
}

func (e Placeholder) String() string {
    if !e.IsValid() {
        // TODO maybe return error
        return placeholders[endPlaceholder]
    }

    return placeholders[e]
}

type Direction int

const (
    // use “iota + 1” to be sure that the enum type is initialized.
    DirectionAsc Direction = iota + 1
    DirectionDesc
    endDirection
)

var directions = map[Direction]string{
    DirectionAsc:  "ASC",
    DirectionDesc: "DESC",
}

func (e Direction) IsValid() bool {
    return e >= DirectionAsc && e < endDirection
}

func (e Direction) Ordinal() int {
    return int(e)
}

func (e Direction) String() string {
    if !e.IsValid() {
        // TODO maybe return error
        return directions[endDirection]
    }

    return directions[e]
}
