package generate

import (
	"context"

	"github.com/google/uuid"
)

type T Model

type _T_Args struct {
	ID uuid.UUID
}

func (*query) _T_(ctx context.Context, args _T_Args) *T {
	t := &T{}
	return t
}

type _T_Filter interface{}

type _T_SearchArgs struct {
	Limit  int32
	Count  *int32
	Filter *_T_Filter
}

func (*query) _T_Search(ctx context.Context, args _T_SearchArgs) *T {
	t := &T{}
	return t
}
