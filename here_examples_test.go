package here_test

import (
	"fmt"
	"path"

	"github.com/dc0d/here"
)

func ExampleMark() {
	calls := here.Mark()
	calls = calls[:1]
	for i, c := range calls {
		c.File = path.Base(c.File)
		calls[i] = c
	}
	fmt.Println(calls)

	// Output:
	// here_examples_test.go:11 github.com/dc0d/here_test.ExampleMark
}
