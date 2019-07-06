package _package

import (
	"context"

	"github.com/google/uuid"
)

type _T Model

// _function_start_

// _TArgs are the arguments for the _T command
type _TArgs struct {
	ID uuid.UUID
}

// _T finds a _T in the database
func (*query) _T(ctx context.Context, args _TArgs) *_T {
	return &_T{}
}

// _TFilter is a filter to search the _T database
type _TFilter struct{}

// _TSearchArgs are the arguments for the _TSearch command
type _TSearchArgs struct {
	Limit  int32
	Count  *int32
	Filter *_TFilter
}

// _TResult is the result of a query of the _T database
type _TResult struct {
	Count int32
	Data  []_T
}

// _TSearch searches the _T database and returls all metching reccords
func (*query) _TSearch(ctx context.Context, args _TSearchArgs) _TResult {
	return _TResult{}
}
