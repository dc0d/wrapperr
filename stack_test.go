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

func TestStack(t *testing.T) {
	var suite suiteStack

	t.Run(`to string`, suite.toString)
}

type suiteStack struct{}

func (suiteStack) toString(t *testing.T) {
	var (
		assert = assert.New(t)

		stack          wrapperr.Stack
		expectedString string
	)

	{
		stack = []wrapperr.Annotation{
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

		expectedString = "file-1:1 fn1 - message-1\n>> file-2:2 fn2 - message-2"
	}

	actualString := fmt.Sprint(stack)

	assert.Equal(expectedString, actualString)
}

func TestWithStack(t *testing.T) {
	var suite suiteWithStack

	t.Run(`use short file path`, suite.useShortFilePath)
	t.Run(`to string`, suite.toString)
	t.Run(`unwrap cause`, suite.unwrapCause)
	t.Run(`with message`, suite.withMessage)
	t.Run(`with annotation`, suite.withAnnotation)
	t.Run(`with empty annotation`, suite.withEmptyAnnotation)
	t.Run(`to json`, suite.toJSON)
}

type suiteWithStack struct{}

func (suiteWithStack) useShortFilePath(t *testing.T) {
	var (
		assert = assert.New(t)

		err error
	)

	{
		err = fn2()
	}

	actualError := err.(wrapperr.TracedErr)
	for _, a := range actualError.Stack {
		assert.Equal(shortFilePath(a.Loc.File), a.Loc.File)
	}
}

func (suiteWithStack) toString(t *testing.T) {
	var (
		assert = assert.New(t)

		err             error
		expectedStrings []string
	)

	{
		err = fn2()

		expectedStrings = []string{
			"stack_fixtures_test.go:8 github.com/dc0d/wrapperr_test.fn1",
			"stack_fixtures_test.go:12 github.com/dc0d/wrapperr_test.fn2",
			"stack_test.go",
			"github.com/dc0d/wrapperr_test.suiteWithStack.toString",
			"CAUSEERR",
		}
	}

	actualString := fmt.Sprint(err)

	for _, txt := range expectedStrings {
		assert.Contains(actualString, txt)
	}
}

func (suiteWithStack) unwrapCause(t *testing.T) {
	var (
		assert = assert.New(t)

		err           error
		expectedError = rootCause
	)

	{
		err = fn2()
	}

	actualError := errors.Unwrap(err)

	assert.Equal(expectedError, actualError)
}

func (suiteWithStack) withMessage(t *testing.T) {
	var (
		assert = assert.New(t)

		err             error
		expectedStrings []string
	)

	{
		err = fn3()

		expectedStrings = []string{
			"stack_test.go",
			"github.com/dc0d/wrapperr_test.suiteWithStack.withMessage",
			"CAUSEERR",
			"message 3",
		}
	}

	actualString := fmt.Sprint(err)

	for _, txt := range expectedStrings {
		assert.Contains(actualString, txt)
	}
}

func (suiteWithStack) withAnnotation(t *testing.T) {
	var (
		assert = assert.New(t)

		err             error
		expectedStrings []string
	)

	{
		err = fn6()

		expectedStrings = []string{
			"stack_fixtures_test.go:16 github.com/dc0d/wrapperr_test.fn3 - message 3",
			"stack_fixtures_test.go:20 github.com/dc0d/wrapperr_test.fn4",
			"stack_fixtures_test.go:24 github.com/dc0d/wrapperr_test.fn5 - message 5",
			"stack_fixtures_test.go:28 github.com/dc0d/wrapperr_test.fn6",
			"stack_test.go",
			"github.com/dc0d/wrapperr_test.suiteWithStack.withAnnotation",
			"CAUSEERR",
		}
	}

	actualString := fmt.Sprint(err)

	for _, txt := range expectedStrings {
		assert.Contains(actualString, txt)
	}
}

func (suiteWithStack) withEmptyAnnotation(t *testing.T) {
	var (
		assert = assert.New(t)

		err             error
		expectedStrings []string
	)

	{
		err = fn9()

		expectedStrings = []string{
			"stack_fixtures_test.go:16 github.com/dc0d/wrapperr_test.fn3 - message 3",
			"stack_fixtures_test.go:32 github.com/dc0d/wrapperr_test.fn7",
			"stack_fixtures_test.go:36 github.com/dc0d/wrapperr_test.fn8",
			"stack_fixtures_test.go:40 github.com/dc0d/wrapperr_test.fn9",
			"stack_test.go",
			"github.com/dc0d/wrapperr_test.suiteWithStack.withEmptyAnnotation",
			"CAUSEERR",
		}
	}

	actualString := fmt.Sprint(err)

	for _, txt := range expectedStrings {
		assert.Contains(actualString, txt)
	}
}

func (suiteWithStack) toJSON(t *testing.T) {
	var (
		assert = assert.New(t)

		err             error
		expectedStrings []string
	)

	{
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
			"github.com/dc0d/wrapperr_test.suiteWithStack.toJSON",
			"CAUSEERR",
		}
	}

	js, jsErr := json.Marshal(err)
	if jsErr != nil {
		panic(jsErr)
	}
	actualString := string(js)

	for _, txt := range expectedStrings {
		assert.Contains(actualString, txt)
	}
}

func shortFilePath(fp string) string {
	return path.Join(path.Base(path.Dir(fp)), path.Base(fp))
}

var (
	rootCause       = errors.New("CAUSEERR")
	emptyAnnotation = ""
)
