package wrapperr_test

import (
	"fmt"
	"testing"

	"github.com/dc0d/wrapperr"

	"github.com/stretchr/testify/assert"
)

func TestAnnotation(t *testing.T) {
	var suite suiteAnnotation

	t.Run(`to string`, suite.toString)
	t.Run(`to string without message`, suite.toStringWthoutMessage)
}

type suiteAnnotation struct{}

func (suiteAnnotation) toString(t *testing.T) {
	var (
		assert = assert.New(t)

		note           wrapperr.Annotation
		expectedString string
	)

	{
		var loc wrapperr.Loc
		loc.File = "file"
		loc.Line = 10
		loc.Func = "fn"

		note.Loc = loc
		note.Message = "some message"

		expectedString = "file:10 fn - some message"
	}

	actualString := fmt.Sprint(note)

	assert.Equal(expectedString, actualString)
}

func (suiteAnnotation) toStringWthoutMessage(t *testing.T) {
	var (
		assert = assert.New(t)

		note           wrapperr.Annotation
		expectedString string
	)

	{
		var loc wrapperr.Loc
		loc.File = "file"
		loc.Line = 10
		loc.Func = "fn"

		note.Loc = loc

		expectedString = "file:10 fn"
	}

	actualString := fmt.Sprint(note)

	assert.Equal(expectedString, actualString)
}
