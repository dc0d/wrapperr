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

var (
	rootCause = errors.New("ROOT CAUSE ERROR")
)

type sentinelErr string

func (s sentinelErr) Error() string { return string(s) }

// ?: recursive
