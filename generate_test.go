package potofgreed

import (
	"bytes"
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
package: models
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
				Package: "models",
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

func TestGenerateGoTypes(t *testing.T) {
	buf := &bytes.Buffer{}
	err := GenerateGoTypes(&Options{
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
	}, buf)

	expected := "type package \n\ntype RawBook struct {\n\tAuthors []string `json:\"authors\"`\n\tChapter *float32 `json:\"chapter\"`\n\tSeries *string `json:\"series\"`\n\tTitle *string `json:\"title\"`\n\tVolume *int32 `json:\"volume\"`\n}\ntype Book struct {\n\t RawBook\n\tUserBook *UserBook `json:\"user_book\"`\n}\ntype RawUser struct {\n\tName *string `json:\"name\"`\n\tPassword *string `json:\"password\"`\n}\ntype User struct {\n\t RawUser\n\tUserBook *UserBook `json:\"user_book\"`\n}\ntype RawUserBook struct {\n\tCurrentPage *int32 `json:\"current_page\"`\n\tRating *float32 `json:\"rating\"`\n}\ntype UserBook struct {\n\t RawUserBook\n\tBook *Book `json:\"book\"`\n\tUser *User `json:\"user\"`\n}\n"
	assert.NoError(t, err)
	assert.Equal(t, expected, buf.String())
}

func TestGenerateGraphQL(t *testing.T) {
	buf := &bytes.Buffer{}
	err := GenerateGraphQL(&Options{
		Version: 1,
		Models: map[string]Model{
			"Book": Model{
				"authors": "[String!]!",
				"chapter": "Float",
				"series":  "String!",
				"title":   "String!",
				"volume":  "Int",
			},
			"User": Model{
				"name":     "String!",
				"password": "String!",
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
	}, buf)

	expected := "\nschema {\n\tquery: RootQuery\n\tmutation: RootMutation\n}\ntype RootQuery {\n\tbook(id: ID!): Book\n\tbook_query(filter: BookFilter limit: Int! skip: Int): Book\n\tuser(id: ID!): User\n\tuser_query(filter: UserFilter limit: Int! skip: Int): User\n\tuser_book(id: ID!): UserBook\n\tuser_book_query(filter: UserBookFilter limit: Int! skip: Int): UserBook\n}\ntype RootMuttation {\n\tnew_book(data: BookInput!): Book\n\tupdate_book(id: ID! data: BookInput!): Book\n\tdelete_book(id: ID!): Book\n\tnew_user(data: UserInput!): User\n\tupdate_user(id: ID! data: UserInput!): User\n\tdelete_user(id: ID!): User\n\tnew_user_book(data: UserBookInput!): UserBook\n\tupdate_user_book(id: ID! data: UserBookInput!): UserBook\n\tdelete_user_book(id: ID!): UserBook\n}\ninput BookInput {\n\tauthors: [String!]\n\tchapter: Float\n\tseries: String\n\ttitle: String\n\tvolume: Int\n}\ntype Book {\n\tauthors: [String!]!\n\tchapter: Float\n\tseries: String!\n\ttitle: String!\n\tuser_book: UserBook\n\tvolume: Int\n}\ninput UserInput {\n\tname: String\n\tpassword: String\n}\ntype User {\n\tname: String!\n\tpassword: String!\n\tuser_book: UserBook\n}\ninput UserBookInput {\n\tcurrent_page: Int\n\trating: Float\n}\ntype UserBook {\n\tbook: Book\n\tcurrent_page: Int\n\trating: Float\n\tuser: User\n}\n"
	assert.NoError(t, err)
	assert.Equal(t, expected, buf.String())
}
