package _package

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/comicbox/comicbox/comicboxd/data"
	"github.com/comicbox/comicbox/comicboxd/errors"
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
	dir := "comicboxd/app/schema/gql"
	s := ""
	files, err := data.AssetDir(dir)
	errors.Check(err)
	for _, file := range files {
		s += string(data.MustAsset(filepath.Join(dir, file))) + "\n"
	}
	schema, err := graphql.ParseSchema(s, &RootQuery{})
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	return addUser(&relay.Handler{Schema: schema})
}
