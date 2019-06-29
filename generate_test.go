package potofgreed

import (
	"os"
	"strings"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *Options
		err      bool
	}{
		{
			name: "basic",
			input: `
version: 1
models:
  User:
    name: String
    password: String
  Book:
    title: String
    series: String
    authors: "[String!]!"
    chapter: Float
    volume: Int
  UserBook:
    current_page: Int
    rating: Float
relationships:
  - from_type: User
    from_count: one
    to_type: UserBook
    to_count: many
  - from_type: UserBook
    from_count: one
    to_type: Book
    to_count: one
`,
			expected: &Options{
				Version: 1,
				Models: map[string]Model{
					"Book": Model{
						"authors": "[String!]!",
						"chapter": "Float",
						"series":  "String",
						"title":   "String",
						"volume":  "Int",
					},
					"User": Model{
						"name":     "String",
						"password": "String",
					},
					"UserBook": Model{
						"current_page": "Int",
						"rating":       "Float",
					},
				},
				Relationships: []Relationship{
					Relationship{
						FromType:  "User",
						FromCount: "one",
						ToType:    "UserBook",
						ToCount:   "many",
					},
					Relationship{
						FromType:  "UserBook",
						FromCount: "one",
						ToType:    "Book",
						ToCount:   "one",
					},
				},
			},
			err: false,
		},
		{
			name:     "error",
			input:    "invalid options",
			expected: nil,
			err:      true,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			o, err := Load(strings.NewReader(test.input))

			if test.err {
				assert.Error(t, err)
				return
			}
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(t, test.expected, o)
		})
	}
}

func TestGenerate(t *testing.T) {
	Generate(&Options{
		Version: 1,
		Models: map[string]Model{
			"Book": Model{
				"authors": "[String!]!",
				"chapter": "Float",
				"series":  "String",
				"title":   "String",
				"volume":  "Int",
			},
			"User": Model{
				"name":     "String",
				"password": "String",
			},
			"UserBook": Model{
				"current_page": "Int",
				"rating":       "Float",
			},
		},
		Relationships: []Relationship{
			Relationship{
				FromType:  "User",
				FromCount: "one",
				ToType:    "UserBook",
				ToCount:   "many",
			},
			Relationship{
				FromType:  "UserBook",
				FromCount: "one",
				ToType:    "Book",
				ToCount:   "one",
			},
		},
	}, os.Stderr)
	os.Exit(1)
}
