package data

import (
	"io/ioutil"
	"net/http"
)

func MustOpen(path string) http.File {
	f, err := Assets.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

func ReadFile(path string) ([]byte, error) {
	f, err := Assets.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func MustReadFile(path string) []byte {
	b, err := ReadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}
