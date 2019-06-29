package potofgreed

import (
	"fmt"
	"io"
	"sort"

	"github.com/go-yaml/yaml"
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

func GenerateGoTypes(options *Options, out io.Writer) error {

	_, err := fmt.Fprintf(out, "type package %s\n\n", options.Package)
	if err != nil {
		return xerrors.Errorf("failed to write go package: %w", err)
	}
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

	for _, model := range models {
		goSrc, err := model.Model.Golang()
		if err != nil {
			return xerrors.Errorf("failed to generate go definition for %s: %w", model.Name, err)
		}

		_, err = fmt.Fprintf(out, "type %s %s\n", model.Name, goSrc)
		if err != nil {
			return xerrors.Errorf("failed to write go source %s: %w", model.Name, err)
		}
	}
	return nil
}
