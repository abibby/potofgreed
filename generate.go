package potofgreed

import (
	"fmt"
	"io"
	"sort"

	"github.com/go-yaml/yaml"
	"github.com/iancoleman/strcase"
	"golang.org/x/xerrors"
)

// Options is a configuration object used to generate the go files
type Options struct {
	Version       float32          `yaml:"version"`
	Package       string           `yaml:"package"`
	Models        map[string]Model `yaml:"models"`
	Relationships []Relationship   `yaml:"relationships"`
}

type Relationship struct {
	FromType  string `yaml:"from_type"`
	FromCount string `yaml:"from_count"`
	ToType    string `yaml:"to_type"`
	ToCount   string `yaml:"to_count"`
}

func Load(options io.Reader) (*Options, error) {
	o := &Options{
		Package: "models",
	}
	err := yaml.NewDecoder(options).Decode(o)
	if err != nil {
		return nil, xerrors.Errorf("couldn't decode options: %w", err)
	}
	return o, nil
}

func GenerateGoTypes(options *Options) (string, error) {

	src := fmt.Sprintf("package %s\n\n", options.Package)
	type namedModel struct {
		Name  string
		Model Model
	}
	models := []namedModel{}
	for name, model := range options.Models {
		models = append(models, namedModel{name, model})
	}

	sort.Slice(models, func(i, j int) bool {
		return models[i].Name < models[j].Name
	})

	for _, rawModel := range models {
		goSrc, err := rawModel.Model.Golang()
		if err != nil {
			return "", xerrors.Errorf("failed to generate go definition for Raw%s: %w", rawModel.Name, err)
		}

		src += fmt.Sprintf("type Raw%s %s\n", rawModel.Name, goSrc)

		model := Model{
			"": Type("Raw" + rawModel.Name).NotNull(),
		}

		for _, relation := range options.Relationships {
			if relation.FromType == rawModel.Name {
				model[relation.ToType] = Type(relation.ToType)
			}
			if relation.ToType == rawModel.Name {
				model[relation.FromType] = Type(relation.FromType)
			}
		}

		goSrc, err = model.Golang()
		if err != nil {
			return "", xerrors.Errorf("failed to generate go definition for %s: %w", rawModel.Name, err)
		}

		src += fmt.Sprintf("type %s %s\n", rawModel.Name, goSrc)
	}
	return src, nil
}

func GenerateGraphQL(options *Options) (string, error) {
	src := ""
	type namedModel struct {
		Name       string
		Model      Model
		InputModel Model
	}
	models := []namedModel{}
	for name, model := range options.Models {
		models = append(models, namedModel{name, model.Clone(), model.Nullable()})
	}

	sort.Slice(models, func(i, j int) bool {
		return models[i].Name < models[j].Name
	})

	src += fmt.Sprintf(`
schema {
	query: RootQuery
	mutation: RootMutation
}
type RootQuery {
`)

	for _, model := range models {
		src += fmt.Sprintf("\t%s(id: ID!): %s\n", strcase.ToSnake(model.Name), model.Name)
		src += fmt.Sprintf("\t%s_query(filter: %sFilter limit: Int! skip: Int): %s\n", strcase.ToSnake(model.Name), model.Name, model.Name)
	}

	src += fmt.Sprintf("}\ntype RootMuttation {\n")

	for _, model := range models {
		src += fmt.Sprintf("\tnew_%s(data: %sInput!): %s\n", strcase.ToSnake(model.Name), model.Name, model.Name)
		src += fmt.Sprintf("\tupdate_%s(id: ID! data: %sInput!): %s\n", strcase.ToSnake(model.Name), model.Name, model.Name)
		src += fmt.Sprintf("\tdelete_%s(id: ID!): %s\n", strcase.ToSnake(model.Name), model.Name)
	}
	src += fmt.Sprintf("}\n")

	for _, model := range models {
		for _, relation := range options.Relationships {
			if relation.FromType == model.Name {
				model.Model[relation.ToType] = Type(relation.ToType)
			}
			if relation.ToType == model.Name {
				model.Model[relation.FromType] = Type(relation.FromType)
			}
		}

		gqlSrc, err := model.InputModel.GraphQL()
		if err != nil {
			return "", xerrors.Errorf("failed to generate go definition for %s: %w", model.Name, err)
		}

		src += fmt.Sprintf("input %sInput %s\n", model.Name, gqlSrc)

		gqlSrc, err = model.Model.GraphQL()
		if err != nil {
			return "", xerrors.Errorf("failed to generate go definition for %s: %w", model.Name, err)
		}
		src += fmt.Sprintf("type %s %s\n", model.Name, gqlSrc)
	}
	return src, nil
}
