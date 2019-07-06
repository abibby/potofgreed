package models

import (
	"context"

	"github.com/google/uuid"
)

type _T Model

// BookArgs are the arguments for the Book command
type BookArgs struct {
	ID uuid.UUID
}

// Book finds a Book in the database
func (*query) Book(ctx context.Context, args BookArgs) *Book {
	return &Book{}
}

// BookFilter is a filter to search the Book database
type BookFilter struct{}

// BookSearchArgs are the arguments for the BookSearch command
type BookSearchArgs struct {
	Limit  int32
	Count  *int32
	Filter *BookFilter
}

// BookResult is the result of a query of the Book database
type BookResult struct {
	Count int32
	Data  []Book
}

// BookSearch searches the Book database and returls all metching reccords
func (*query) BookSearch(ctx context.Context, args BookSearchArgs) BookResult {
	return BookResult{}
}

// UserArgs are the arguments for the User command
type UserArgs struct {
	ID uuid.UUID
}

// User finds a User in the database
func (*query) User(ctx context.Context, args UserArgs) *User {
	return &User{}
}

// UserFilter is a filter to search the User database
type UserFilter struct{}

// UserSearchArgs are the arguments for the UserSearch command
type UserSearchArgs struct {
	Limit  int32
	Count  *int32
	Filter *UserFilter
}

// UserResult is the result of a query of the User database
type UserResult struct {
	Count int32
	Data  []User
}

// UserSearch searches the User database and returls all metching reccords
func (*query) UserSearch(ctx context.Context, args UserSearchArgs) UserResult {
	return UserResult{}
}

// UserBookArgs are the arguments for the UserBook command
type UserBookArgs struct {
	ID uuid.UUID
}

// UserBook finds a UserBook in the database
func (*query) UserBook(ctx context.Context, args UserBookArgs) *UserBook {
	return &UserBook{}
}

// UserBookFilter is a filter to search the UserBook database
type UserBookFilter struct{}

// UserBookSearchArgs are the arguments for the UserBookSearch command
type UserBookSearchArgs struct {
	Limit  int32
	Count  *int32
	Filter *UserBookFilter
}

// UserBookResult is the result of a query of the UserBook database
type UserBookResult struct {
	Count int32
	Data  []UserBook
}

// UserBookSearch searches the UserBook database and returls all metching reccords
func (*query) UserBookSearch(ctx context.Context, args UserBookSearchArgs) UserBookResult {
	return UserBookResult{}
}
