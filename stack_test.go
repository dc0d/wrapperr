package wrapperr_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"testing"

	"github.com/dc0d/wrapperr"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	t.Run(`to string`, func(t *testing.T) {
		const expectedString = "file-1:1 fn1 - message-1\n>> file-2:2 fn2 - message-2"
		var stack wrapperr.Stack = []wrapperr.Annotation{
			{
				Loc: wrapperr.Loc{
					File: "file-1",
					Line: 1,
					Func: "fn1",
				},
				Message: "message-1",
			},
			{
				Loc: wrapperr.Loc{
					File: "file-2",
					Line: 2,
					Func: "fn2",
				},
				Message: "message-2",
			},
		}

		actualString := fmt.Sprint(stack)

		assert.Equal(t, expectedString, actualString)
	})
}

func Test_WithStack_messages(t *testing.T) {
	t.Run(`with message`, func(t *testing.T) {
		var (
			err             error
			expectedStrings []string
		)

		err = fn3()

		expectedStrings = []string{
			"stack_test.go",
			"github.com/dc0d/wrapperr_test",
			"CAUSEERR",
			"message 3",
		}

		actualString := fmt.Sprint(err)

		for _, txt := range expectedStrings {
			assert.Contains(t, actualString, txt)
		}
	})

	t.Run(`with multiple messages`, func(t *testing.T) {
		var (
			err             error
			expectedStrings []string
		)

		err = fn10()

		expectedStrings = []string{
			message1,
			message2,
			message3,
			message1 + " " + message2 + " " + message3,
		}

		actualString := fmt.Sprint(err)

		for _, txt := range expectedStrings {
			assert.Contains(t, actualString, txt)
		}
	})
}

func Test_WithStack_annotations(t *testing.T) {
	t.Run(`with annotation`, func(t *testing.T) {
		var (
			err             error
			expectedStrings []string
		)

		err = fn6()

		expectedStrings = []string{
			"stack_fixtures_test.go:16 github.com/dc0d/wrapperr_test.fn3 - message 3",
			"stack_fixtures_test.go:20 github.com/dc0d/wrapperr_test.fn4",
			"stack_fixtures_test.go:24 github.com/dc0d/wrapperr_test.fn5 - message 5",
			"stack_fixtures_test.go:28 github.com/dc0d/wrapperr_test.fn6",
			"stack_test.go",
			"CAUSEERR",
		}

		actualString := fmt.Sprint(err)

		for _, txt := range expectedStrings {
			assert.Contains(t, actualString, txt)
		}
	})

	t.Run(`with empty annotation`, func(t *testing.T) {
		var (
			err             error
			expectedStrings []string
		)

		err = fn9()

		expectedStrings = []string{
			"stack_fixtures_test.go:16 github.com/dc0d/wrapperr_test.fn3 - message 3",
			"stack_fixtures_test.go:32 github.com/dc0d/wrapperr_test.fn7",
			"stack_fixtures_test.go:36 github.com/dc0d/wrapperr_test.fn8",
			"stack_fixtures_test.go:40 github.com/dc0d/wrapperr_test.fn9",
			"stack_test.go",
			"CAUSEERR",
		}

		actualString := fmt.Sprint(err)

		for _, txt := range expectedStrings {
			assert.Contains(t, actualString, txt)
		}
	})
}

func Test_WithStack(t *testing.T) {
	t.Run(`use short file path`, func(t *testing.T) {
		err := fn2()

		actualError := err.(wrapperr.TracedErr)
		for _, a := range actualError.Stack {
			assert.Equal(t, shortFilePath(a.Loc.File), a.Loc.File)
		}
	})

	t.Run(`to string`, func(t *testing.T) {
		var (
			err             error
			expectedStrings []string
		)

		err = fn2()

		expectedStrings = []string{
			"stack_fixtures_test.go:8 github.com/dc0d/wrapperr_test.fn1",
			"stack_fixtures_test.go:12 github.com/dc0d/wrapperr_test.fn2",
			"stack_test.go",
			"CAUSEERR",
		}

		actualString := fmt.Sprint(err)

		for _, txt := range expectedStrings {
			assert.Contains(t, actualString, txt)
		}
	})

	t.Run(`unwrap cause`, func(t *testing.T) {
		var (
			err           error
			expectedError = errRootCause
		)

		err = fn2()

		actualError := errors.Unwrap(err)

		assert.Equal(t, expectedError, actualError)
	})

	t.Run(`to json`, func(t *testing.T) {
		var (
			err             error
			expectedStrings []string
		)

		err = fn6()

		expectedStrings = []string{
			"stack_fixtures_test.go:16",
			"github.com/dc0d/wrapperr_test.fn3",
			"message 3",
			"stack_fixtures_test.go:20",
			"github.com/dc0d/wrapperr_test.fn4",
			"stack_fixtures_test.go:24",
			"github.com/dc0d/wrapperr_test.fn5",
			"message 5",
			"stack_fixtures_test.go:28",
			"github.com/dc0d/wrapperr_test.fn6",
			"stack_test.go",
			"CAUSEERR",
		}

		js, jsErr := json.Marshal(err)
		if jsErr != nil {
			panic(jsErr)
		}
		actualString := string(js)

		for _, txt := range expectedStrings {
			assert.Contains(t, actualString, txt)
		}
	})

	t.Run(`to json with specialized error struct cause`, func(t *testing.T) {
		sampleError := sampleErr{Data1: "data 1", Data2: "data 2"}
		err := wrapperr.WithStack(sampleError)

		expectedStrings := []string{
			"data_1",
			"data_2",
		}

		js, jsErr := json.Marshal(err)
		if jsErr != nil {
			panic(jsErr)
		}
		actualString := string(js)

		for _, txt := range expectedStrings {
			assert.Contains(t, actualString, txt)
		}
	})
}

func Test_WithStackf(t *testing.T) {
	t.Run(`fotmats the message properly`, func(t *testing.T) {
		testCases := []struct {
			err           func() error
			expectedTexts []string
		}{
			{err: fn11, expectedTexts: []string{"fn11"}},
			{err: fn11, expectedTexts: []string{"data 1"}},
			{err: fn11, expectedTexts: []string{"sampleErr"}},
			{err: fn11, expectedTexts: []string{"the cause is a special error of type"}},
		}

		for _, tc := range testCases {
			err := tc.err()

			errorText := err.Error()

			for _, txt := range tc.expectedTexts {
				assert.Contains(t, errorText, txt)
			}
		}
	})
}

func Test_fix_middle_annotation(t *testing.T) {
	t.Run(`annotated indirectly`, func(t *testing.T) {
		errTxt := fix1().Error()

		assert.Contains(t, errTxt, "annotated root cause")
		assert.Contains(t, errTxt, "sample annotation")
	})

	t.Run(`annotated indirectly multiple times`, func(t *testing.T) {
		errTxt := fix4().Error()

		assert.Contains(t, errTxt, "annotated root cause")
		assert.Contains(t, errTxt, "sample annotation")
		assert.Contains(t, errTxt, "second sample annotation")
		assert.Contains(t, errTxt, "fix2")
		assert.Contains(t, errTxt, "fix3")
		assert.Contains(t, errTxt, "fix4")
	})
}

func shortFilePath(fp string) string {
	return path.Join(path.Base(path.Dir(fp)), path.Base(fp))
}

type sampleErr struct {
	Data1 string `json:"data_1,omitempty"`
	Data2 string `json:"data_2,omitempty"`
}

func (err sampleErr) Error() string { return err.Data1 + " " + err.Data2 }

var (
	errRootCause    = errors.New("CAUSEERR")
	emptyAnnotation = ""
)
