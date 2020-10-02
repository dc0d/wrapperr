package here_test

import (
	"encoding/json"
	"fmt"

	"github.com/dc0d/here"
)

func ExampleMark() {
	fmt.Println(here.Mark(rootCause))

	// Output:
	// [here_examples_test.go:11 here_test.ExampleMark]
	// ROOT CAUSE ERROR
}

func ExampleMark_second() {
	fmt.Println(here.Mark(firstFn()))

	// Output:
	// [here_examples_test.go:19 here_test.ExampleMark_second]
	// [here_call_fixture_test.go:8 here_test.firstFn]
	// ROOT CAUSE ERROR
}

func ExampleMark_json() {
	err := here.Mark(secondFn())

	js, _ := json.MarshalIndent(err, "", "  ")
	fmt.Println(string(js))

	// Output:
	// {
	//   "Calls": [
	//     "[here_examples_test.go:28 here_test.ExampleMark_json]",
	//     "[here_call_fixture_test.go:16 here_test.secondFn]",
	//     "[here_call_fixture_test.go:8 here_test.firstFn]"
	//   ],
	//   "Cause": "ROOT CAUSE ERROR"
	// }
}
