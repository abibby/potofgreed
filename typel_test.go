package potofgreed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType_internalType(t *testing.T) {
	tests := []struct {
		name     string
		input    Type
		expected *internalType
		err      bool
	}{
		{
			name:  "simple",
			input: Type("String"),
			expected: &internalType{
				BaseType: "String",
				Array:    false,
				Nullable: true,
			},
		},
		{
			name:  "array",
			input: Type("[Int]"),
			expected: &internalType{
				BaseType: "Int",
				Array:    true,
				Nullable: true,
			},
		},
		{
			name:  "not_nullable",
			input: Type("Int!"),
			expected: &internalType{
				BaseType: "Int",
				Array:    false,
				Nullable: false,
			},
		},
		{
			name:  "not_nullable_array",
			input: Type("[Int]!"),
			expected: &internalType{
				BaseType: "Int",
				Array:    true,
				Nullable: false,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			iType, err := test.input.internalType()

			if test.err {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, test.expected, iType)
		})
	}
}

func TestType_Golang(t *testing.T) {
	tests := []struct {
		name     string
		input    Type
		expected string
		err      bool
	}{
		{
			name:     "simple",
			input:    Type("String"),
			expected: "*string",
		},
		{
			name:     "array",
			input:    Type("[Int]"),
			expected: "*[]*int32",
		},
		{
			name:     "not_nullable",
			input:    Type("Int!"),
			expected: "int32",
		},
		{
			name:     "not_nullable_array",
			input:    Type("[Int]!"),
			expected: "[]*int32",
		},
		{
			name:     "not_nullable_not_nullable_array",
			input:    Type("[Int!]!"),
			expected: "[]int32",
		},
		{
			name:     "non_basic_type",
			input:    Type("StructType!"),
			expected: "StructType",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			goSrc, err := test.input.Golang()

			if test.err {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, test.expected, goSrc)
		})
	}
}
func TestType_GraphQL(t *testing.T) {
	tests := []struct {
		name     string
		input    Type
		expected string
		err      bool
	}{
		{
			name:     "simple",
			input:    Type("String"),
			expected: "String",
		},
		{
			name:     "array",
			input:    Type("[Int]"),
			expected: "[Int]",
		},
		{
			name:     "not_nullable",
			input:    Type("Int!"),
			expected: "Int!",
		},
		{
			name:     "not_nullable_array",
			input:    Type("[Int]!"),
			expected: "[Int]!",
		},
		{
			name:     "not_nullable_not_nullable_array",
			input:    Type("[Int!]!"),
			expected: "[Int!]!",
		},
		{
			name:     "non_basic_type",
			input:    Type("StructType!"),
			expected: "StructType!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			goSrc, err := test.input.GraphQL()

			if test.err {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, test.expected, goSrc)
		})
	}
}
