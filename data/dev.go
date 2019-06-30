// +build dev

package data

import "net/http"

// to be used with https://github.com/shurcooL/vfsgen

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("generate")
