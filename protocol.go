package repogen

type Generator interface {
	EntityType(entity interface{}) (EntityType, error)
}

type EntityType interface {
	MethodType(name string) MethodType
}

type MethodType interface {
	Params() Returner
}

type Returner interface {
	Build() QueryTypeBuilder
}

type QueryTypeBuilder interface {
	Select() SelectQuery
	Insert() InsertQuery
	Update() UpdateQuery
	Delete() DeleteQuery
}

type SelectQuery interface {
	AddWhereClause() WhereClause
	AddLimitClause() LimitClause
	AddOffsetClause() OffsetClause
	AddOrderByClause() OrderByClause
}

type InsertQuery interface {
	Columns(columns ...string) InsertColumnsClause
	Values(values ...interface{}) InsertValuesClause
}

type InsertColumnsClause interface {
	Values(values ...interface{}) InsertValuesClause
}

type InsertValuesClause interface {
	BuildQuery()
}

type DeleteQuery interface {
	AddWhereClause() WhereClause
	AddLimitClause() LimitClause
	AddOffsetClause() OffsetClause
}

type UpdateQuery interface {
	AddWhereClause() WhereClause
	AddLimitClause() LimitClause
	AddOffsetClause() OffsetClause
	BuildQuery()
}

type WhereClause interface {
	AddLimitClause() LimitClause
	AddOffsetClause() OffsetClause
	AddOrderByClause() OrderByClause
	BuildQuery()
}

type LimitClause interface {
	AddOffsetClause() OffsetClause
	AddOrderByClause() OrderByClause
	BuildQuery()
}

type OffsetClause interface {
	AddOrderByClause() OrderByClause
	BuildQuery()
}

type OrderByClause interface {
	Asc(column string) OrderByClause
	Desc(column string) OrderByClause
	BuildQuery()
}
