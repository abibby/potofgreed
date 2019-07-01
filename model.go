package potofgreed

import (
	"fmt"
	"sort"

	"github.com/iancoleman/strcase"
	"golang.org/x/xerrors"
)

type Model map[string]Type

type namedType struct {
	Field string
	Type  Type
}

func (m Model) Nullable() Model {
	clone := m.Clone()
	for field, typ := range clone {
		clone[field] = typ.Nullable()
	}
	return clone
}

func (m Model) slice() []namedType {

	types := []namedType{}
	for field, typ := range m {
		types = append(types, namedType{strcase.ToSnake(field), typ})
	}

	// you need to sort the
	sort.Slice(types, func(i, j int) bool {
		return types[i].Field < types[j].Field
	})
	return types
}

// GraphQL returns a representration of the model in GraphQL syntax
func (m Model) GraphQL() (string, error) {

	gqlSrc := "{\n"
	for _, typ := range m.slice() {
		typeSrc, err := typ.Type.GraphQL()
		if err != nil {
			return "", xerrors.Errorf("failed to generate type for %s, :w", typ.Field, err)
		}
		gqlSrc += fmt.Sprintf("\t%s: %s\n", typ.Field, typeSrc)
	}
	gqlSrc += "}"

	return gqlSrc, nil
}

// Golang returns a representration of the model in go syntax
func (m Model) Golang() (string, error) {

	goSrc := "struct {\n"
	for _, typ := range m.slice() {
		typeSrc, err := typ.Type.Golang()
		if err != nil {
			return "", xerrors.Errorf("failed to generate type for %s, :w", typ.Field, err)
		}
		goSrc += fmt.Sprintf("\t%s %s", strcase.ToCamel(typ.Field), typeSrc)
		if typ.Field != "" {
			goSrc += fmt.Sprintf(" `json:\"%s\"`", strcase.ToSnake(typ.Field))
		}
		goSrc += "\n"
	}
	goSrc += "}"
	return goSrc, nil
}

func (m Model) Clone() Model {
	newModel := Model{}
	for field, typ := range m {
		newModel[field] = typ
	}
	return newModel
}
