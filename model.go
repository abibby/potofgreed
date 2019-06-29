package potofgreed

import (
	"fmt"

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
	goSrc := "struct {\n"
	for field, typ := range m {
		typeSrc, err := typ.Golang()
		if err != nil {
			return "", xerrors.Errorf("failed to generate type for %s, :w", field, err)
		}
		goSrc += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", strcase.ToCamel(field), typeSrc, field)
	}
	goSrc += "}"
	return goSrc, nil
}
