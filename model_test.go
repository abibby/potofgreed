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
			expected: "struct {\n\tTest string `json:\"test\"`\n\tStructField *Book `json:\"struct_field\"`\n}",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			iType, err := test.input.Golang()

			if test.err {
				assert.Error(t, err)
				return
			}

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
			expected: "{\n\ttest: String!\n\tstruct_field: Book\n}",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			iType, err := test.input.GraphQL()

			if test.err {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, test.expected, iType)
		})
	}
}
