package services

type IQueryService[TQuery any, TResult any] interface {
	Execute(query TQuery) TResult
}
