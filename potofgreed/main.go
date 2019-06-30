package main

import (
	"fmt"
	"os"
	"path"

	"github.com/zwzn/potofgreed"
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
	o, err := potofgreed.Load(f)
	check(err)
	goSrc, err := potofgreed.GenerateGoTypes(o)
	check(err)

	check(os.MkdirAll(o.Package, 0777))

	goF, err := os.Create(path.Join(o.Package, "types.go"))
	check(err)
	goF.WriteString(goSrc)
}
