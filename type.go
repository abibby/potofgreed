package potofgreed

import (
	"regexp"
	"strings"

	"golang.org/x/xerrors"
)

const (
	// ErrMalformedType is returned when the entered type is not of a valid type
	ErrMalformedType = SentinelError("types must be alphanumeric and start with a letter")

	// ErrMismatchedBrackets is returned when there are extra square brackets in
	// a type
	ErrMismatchedBrackets = SentinelError("array brackets must be matched")
)

// Type is the type of a column, it uses GraphQL syntax. e.g. String, Int!,
// [Float!]!
type Type string

type internalType struct {
	BaseType Type
	Array    bool
	Nullable bool
}

// GraphQL returns a representration of the type in GraphQL syntax
func (t *Type) GraphQL() (string, error) {
	_, err := t.internalType()
	if err != nil {
		return "", xerrors.Errorf("failed to parse type: %w", err)
	}
	return string(*t), nil
}

// Golang returns a representration of the type in go syntax
func (t *Type) Golang() (string, error) {
	iType, err := t.internalType()
	if err != nil {
		return "", xerrors.Errorf("failed to parse type: %w", err)
	}

	basicTypes := map[Type]string{
		"String": "string",
		"Float":  "float32",
		"Int":    "int32",
	}

	goSrc := ""

	if iType.Nullable {
		goSrc += "*"
	}

	if iType.Array {
		baseSrc, err := iType.BaseType.Golang()
		if err != nil {
			return "", xerrors.Errorf("failed to generating array sub type: %w", err)
		}
		return goSrc + "[]" + baseSrc, nil
	}

	if goType, ok := basicTypes[iType.BaseType]; ok {
		goSrc += goType
	} else {
		goSrc += string(iType.BaseType)
	}

	return goSrc, nil
}

func (t Type) internalType() (*internalType, error) {
	iType := &internalType{
		BaseType: t,
		Nullable: true,
		Array:    false,
	}
	if strings.HasSuffix(string(iType.BaseType), "!") {
		iType.BaseType = iType.BaseType[0 : len(iType.BaseType)-1]
		iType.Nullable = false
	}
	if strings.HasPrefix(string(iType.BaseType), "[") && strings.HasSuffix(string(iType.BaseType), "]") {
		iType.BaseType = iType.BaseType[1 : len(iType.BaseType)-1]
		iType.Array = true
	}

	if strings.HasPrefix(string(iType.BaseType), "[") != (strings.HasSuffix(string(iType.BaseType), "]") || strings.HasSuffix(string(iType.BaseType), "]!")) {
		return nil, ErrMismatchedBrackets
	}

	if !regexp.MustCompile(`^(\[)*[a-zA-Z][a-zA-Z0-9]*!?(\]!?)*$`).Match([]byte(iType.BaseType)) {
		return nil, ErrMalformedType
	}

	return iType, nil
}
