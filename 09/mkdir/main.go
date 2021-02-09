package main

import (
	"os"
)

/**
$ tree .
.
├── main.go
└── test_dir
    └── foo
        └── bar
            └── baz
*/
func main() {
	os.Mkdir("test_dir", 0755)
	os.MkdirAll("test_dir/foo/bar/baz", 0755)
}
