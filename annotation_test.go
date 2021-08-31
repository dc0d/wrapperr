package wrapperr_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dc0d/wrapperr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnnotation_to_string(t *testing.T) {
	var (
		note           wrapperr.Annotation
		expectedString string
	)

	note.Loc = sampleLoc()
	note.Message = someMessage

	expectedString = "package/file.go:9 github.com/user/module/package.(*Struct).method - some message"

	actualString := fmt.Sprint(note)

	assert.Equal(t, expectedString, actualString)
}

func TestAnnotation_to_string_without_message(t *testing.T) {
	const (
		expectedString = "package/file.go:9 github.com/user/module/package.(*Struct).method"
	)

	var (
		note wrapperr.Annotation
	)

	note.Loc = sampleLoc()

	actualString := fmt.Sprint(note)

	assert.Equal(t, expectedString, actualString)
}

func TestAnnotation_to_json(t *testing.T) {
	var (
		note         wrapperr.Annotation
		expectedJSON string
	)

	note.Loc = sampleLoc()
	note.Message = someMessage

	expectedJSON = `"package/file.go:9 github.com/user/module/package.(*Struct).method - some message"`

	js, err := json.Marshal(note)
	if err != nil {
		require.NoError(t, err)
	}
	actualJSON := string(js)

	assert.Equal(t, expectedJSON, actualJSON)
}

const (
	someMessage = "some message"
)
