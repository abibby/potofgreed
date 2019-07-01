package main

import (
	"fmt"
	"os"
	"path"

	"github.com/zwzn/potofgreed"
	"golang.org/x/tools/imports"
)

func check(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	f, err := os.Open("./potofgreed.yml")
	check(err)
	defer func() { check(f.Close()) }()

	o, err := potofgreed.Load(f)
	check(err)

	check(os.MkdirAll(o.Package, 0777))

	goSrc, err := potofgreed.GenerateGoTypes(o)
	check(err)
	writeGoFile(path.Join(o.Package, "types.go"), goSrc)

	goSrc, err = potofgreed.GenerateGoGraphQL(o)
	check(err)
	writeGoFile(path.Join(o.Package, "schema.go"), goSrc)
}

func writeGoFile(fileName string, src []byte) {
	src, err := imports.Process(fileName, src, nil)
	check(err)

	goF, err := os.Create(fileName)
	check(err)
	_, err = goF.Write(src)
	check(err)
}
