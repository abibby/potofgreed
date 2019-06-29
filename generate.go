package potofgreed

import (
	"fmt"
	"io"

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

func Generate(options *Options, out io.Writer) error {

	_, err := fmt.Fprintf(out, "type package %s\n\n", options.Package)
	if err != nil {
		return xerrors.Errorf("failed to write go package: %w", err)
	}

	for name, model := range options.Models {
		goSrc, err := model.Golang()
		if err != nil {
			return xerrors.Errorf("failed to generate go definition for %s: %w", name, err)
		}

		_, err = fmt.Fprintf(out, "type %s %s\n", name, goSrc)
		if err != nil {
			return xerrors.Errorf("failed to write go source %s: %w", name, err)
		}
	}
	return nil
}
