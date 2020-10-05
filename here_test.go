package here_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMark_first(t *testing.T) {
	var (
		assert        = assert.New(t)
		expectedError string
	)

	{
		expectedError = "[here_call_fixture_test.go:8 here_test.firstFn]\nROOT CAUSE ERROR"
	}

	actualError := firstFn()

	assert.Contains(actualError.Error(), expectedError)
}

func TestMark_another(t *testing.T) {
	var (
		assert        = assert.New(t)
		expectedError string
	)

	{
		expectedError = "[here_call_fixture_test.go:12 here_test.anotherFn]\nROOT CAUSE ERROR"
	}

	actualError := anotherFn()

	assert.Contains(actualError.Error(), expectedError)
}

func TestMark_third(t *testing.T) {
	var (
		assert        = assert.New(t)
		expectedError string
	)

	{
		expectedError = "[here_call_fixture_test.go:20 here_test.thirdFn]\n[here_call_fixture_test.go:16 here_test.secondFn]\n[here_call_fixture_test.go:8 here_test.firstFn]\nROOT CAUSE ERROR"
	}

	actualError := thirdFn()

	assert.Contains(actualError.Error(), expectedError)
}

func TestMark_anonymous_func(t *testing.T) {
	var (
		assert        = assert.New(t)
		expectedError string
	)

	{
		expectedError = "[here_call_fixture_test.go:25 here_test.callAnonymousFunc]\n[here_call_fixture_test.go:24 here_test.callAnonymousFunc.func1]\nROOT CAUSE ERROR"
	}

	actualError := callAnonymousFunc()

	assert.Contains(actualError.Error(), expectedError)
}

func TestMark_first_unwrap(t *testing.T) {
	var (
		assert        = assert.New(t)
		expectedError error
	)

	{
		expectedError = rootCause
	}

	actualError := errors.Unwrap(firstFn())

	assert.Equal(expectedError, actualError)
}

func TestMark_third_unwrap(t *testing.T) {
	var (
		assert        = assert.New(t)
		expectedError error
	)

	{
		expectedError = rootCause
	}

	actualError := errors.Unwrap(thirdFn())

	assert.Equal(expectedError, actualError)
}

func TestMark_nil_unwrap(t *testing.T) {
	originalRootCause := rootCause
	defer func() { rootCause = originalRootCause }()
	rootCause = nil

	var (
		assert        = assert.New(t)
		expectedError error
	)

	{
		expectedError = rootCause
	}

	actualError := errors.Unwrap(thirdFn())

	assert.Equal(expectedError, actualError)
}

func TestMark_json(t *testing.T) {
	var (
		assert       = assert.New(t)
		expectedJSON string
	)

	{
		expectedJSON = "{\"Calls\":[\"[here_call_fixture_test.go:20 here_test.thirdFn]\",\"[here_call_fixture_test.go:16 here_test.secondFn]\",\"[here_call_fixture_test.go:8 here_test.firstFn]\"],\"Cause\":\"ROOT CAUSE ERROR\"}"
	}

	err := thirdFn()

	js, err := json.Marshal(err)
	if err != nil {
		t.Fatal(err)
	}
	jsonStr := string(js)

	assert.Equal(expectedJSON, jsonStr)
}

func TestMark_json_nil(t *testing.T) {
	originalRootCause := rootCause
	defer func() { rootCause = originalRootCause }()
	rootCause = nil

	var (
		assert       = assert.New(t)
		expectedJSON string
	)

	{
		expectedJSON = "{\"Calls\":[\"[here_call_fixture_test.go:20 here_test.thirdFn]\",\"[here_call_fixture_test.go:16 here_test.secondFn]\",\"[here_call_fixture_test.go:8 here_test.firstFn]\"],\"Cause\":null}"
	}

	err := thirdFn()

	js, err := json.Marshal(err)
	if err != nil {
		t.Fatal(err)
	}
	jsonStr := string(js)

	assert.Equal(expectedJSON, jsonStr)
}

func TestMark_json_sentinel(t *testing.T) {
	originalRootCause := rootCause
	defer func() { rootCause = originalRootCause }()
	rootCause = sentinelErr("ROOT CAUSE ERROR")

	var (
		assert       = assert.New(t)
		expectedJSON string
	)

	{
		expectedJSON = "{\"Calls\":[\"[here_call_fixture_test.go:20 here_test.thirdFn]\",\"[here_call_fixture_test.go:16 here_test.secondFn]\",\"[here_call_fixture_test.go:8 here_test.firstFn]\"],\"Cause\":\"ROOT CAUSE ERROR\"}"
	}

	err := thirdFn()

	js, err := json.Marshal(err)
	if err != nil {
		t.Fatal(err)
	}
	jsonStr := string(js)

	assert.Equal(expectedJSON, jsonStr)
}

func TestHere_get_call_location(t *testing.T) {
	var (
		assert = assert.New(t)

		expectedLine            = 29
		expectedFilePathSegment = "here/here_call_fixture_test.go"
		expectedFunc            = "github.com/dc0d/here_test.whereIsThisPlace"
	)

	actualLocation := whereIsThisPlace()

	assert.Equal(expectedLine, actualLocation.Line)
	assert.Contains(actualLocation.File, expectedFilePathSegment)
	assert.Equal(actualLocation.Func, expectedFunc)
}

func TestHere_get_call_location_in_upper_caller(t *testing.T) {
	var (
		assert = assert.New(t)

		expectedLine            = 37
		expectedFilePathSegment = "here/here_call_fixture_test.go"
		expectedFunc            = "github.com/dc0d/here_test.theCaller"
	)

	actualLocation := theCaller()

	assert.Equal(expectedLine, actualLocation.Line)
	assert.Contains(actualLocation.File, expectedFilePathSegment)
	assert.Equal(actualLocation.Func, expectedFunc)
}

func TestHere_get_call_location_less_than_one_skip_is_ignored(t *testing.T) {
	var (
		assert = assert.New(t)

		expectedLine            = 41
		expectedFilePathSegment = "here/here_call_fixture_test.go"
		expectedFunc            = "github.com/dc0d/here_test.lessThanOneSkipIsIgnored"
	)

	actualLocation := lessThanOneSkipIsIgnored()

	assert.Equal(expectedLine, actualLocation.Line)
	assert.Contains(actualLocation.File, expectedFilePathSegment)
	assert.Equal(actualLocation.Func, expectedFunc)
}

func TestHere_get_call_location_short_names(t *testing.T) {
	var (
		assert = assert.New(t)

		expectedLine            = 45
		expectedFilePathSegment = "here_call_fixture_test.go"
		expectedFunc            = "here_test.inShortWhereIsThisPlace"
	)

	actualLocation := inShortWhereIsThisPlace()

	assert.Equal(expectedLine, actualLocation.Line)
	assert.Equal(expectedFilePathSegment, actualLocation.File)
	assert.Equal(actualLocation.Func, expectedFunc)
}

var (
	rootCause = errors.New("ROOT CAUSE ERROR")
)

type sentinelErr string

func (s sentinelErr) Error() string { return string(s) }
