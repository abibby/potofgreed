package potofgreed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel_Golang(t *testing.T) {
	tests := []struct {
		name     string
		input    Model
		expected string
		err      bool
	}{
		{
			name: "simple",
			input: Model{
				"test": "String",
			},
			expected: "struct {\n\tTest *string `json:\"test\"`\n}",
		},
		{
			name: "simple",
			input: Model{
				"test":         "String!",
				"struct_field": "Book",
			},
			expected: "struct {\n\tStructField *Book `json:\"struct_field\"`\n\tTest string `json:\"test\"`\n}",
		},
		{
			name: "invalid_type",
			input: Model{
				"test": "1",
			},
			err: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			iType, err := test.input.Golang()

			if test.err {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, test.expected, iType)
		})
	}
}

func TestModel_GraphQL(t *testing.T) {
	tests := []struct {
		name     string
		input    Model
		expected string
		err      bool
	}{
		{
			name: "simple",
			input: Model{
				"test": "String",
			},
			expected: "{\n\ttest: String\n}",
		},
		{
			name: "simple",
			input: Model{
				"test":         "String!",
				"struct_field": "Book",
			},
			expected: "{\n\tstruct_field: Book\n\ttest: String!\n}",
		},
		{
			name: "invalid_type",
			input: Model{
				"test": "1",
			},
			err: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			iType, err := test.input.GraphQL()

			if test.err {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, test.expected, iType)
		})
	}
}

func TestModel_Clone(t *testing.T) {
	input := Model{
		"key": "value",
	}

	clone := input.Clone()

	assert.Equal(t, input, clone)
	clone["bar"] = "foo"
	assert.NotEqual(t, input, clone)
}
