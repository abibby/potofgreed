package potofgreed

import (
	"fmt"
	"sort"

	"github.com/iancoleman/strcase"
	"golang.org/x/xerrors"
)

type Model map[string]Type

// GraphQL returns a representration of the model in GraphQL syntax
func (m Model) GraphQL() (string, error) {
	goSrc := "{\n"
	for field, typ := range m {
		typeSrc, err := typ.GraphQL()
		if err != nil {
			return "", xerrors.Errorf("failed to generate type for %s, :w", field, err)
		}
		goSrc += fmt.Sprintf("\t%s: %s\n", field, typeSrc)
	}
	goSrc += "}"
	return goSrc, nil
}

// Golang returns a representration of the model in go syntax
func (m Model) Golang() (string, error) {

	type namedType struct {
		Field string
		Type  Type
	}
	types := []namedType{}
	for field, typ := range m {
		types = append(types, namedType{field, typ})
	}

	// you need to sort the
	sort.Slice(types, func(i, j int) bool {
		return types[i].Field > types[j].Field
	})

	goSrc := "struct {\n"
	for _, typ := range types {
		typeSrc, err := typ.Type.Golang()
		if err != nil {
			return "", xerrors.Errorf("failed to generate type for %s, :w", typ.Field, err)
		}
		goSrc += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", strcase.ToCamel(typ.Field), typeSrc, typ.Field)
	}
	goSrc += "}"
	return goSrc, nil
}
