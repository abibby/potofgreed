package models

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type Model struct {
	ID uuid.UUID
}

type query struct{}

// StringCompair is used to search for string in GraphQL
type StringCompair struct {
	EQ         string
	NE         string
	LT         string
	LE         string
	GT         string
	GE         string
	Contains   string
	StartsWith string
	EndsWith   string
}

// IntCompair is used to search for ints in GraphQL
type IntCompair struct {
	EQ int32
	NE int32
	LT int32
	LE int32
	GT int32
	GE int32
}

// FloatCompair is used to search for floats in GraphQL
type FloatCompair struct {
	EQ float32
	NE float32
	LT float32
	LE float32
	GT float32
	GE float32
}

// IDCompair is used to search for ids in GraphQL
type IDCompair struct {
	EQ uuid.UUID
	NE uuid.UUID
}

func Handler() http.Handler {
	schema, err := graphql.ParseSchema(string(schema), &query{})
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	return &relay.Handler{Schema: schema}
}
