package wrapperr_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dc0d/wrapperr"

	"github.com/stretchr/testify/assert"
)

func TestLoc(t *testing.T) {
	var suite suiteLoc

	t.Run(`to string`, suite.toString)
	t.Run(`to json`, suite.toJSON)
}

type suiteLoc struct{}

func (suiteLoc) toString(t *testing.T) {
	var (
		assert = assert.New(t)

		loc            wrapperr.Loc
		expectedString string
	)

	{
		loc.File = "file"
		loc.Line = 10
		loc.Func = "fn"

		expectedString = "file:10 fn"
	}

	actualString := fmt.Sprint(loc)

	assert.Equal(expectedString, actualString)
}

func (suiteLoc) toJSON(t *testing.T) {
	var (
		assert = assert.New(t)

		loc          wrapperr.Loc
		expectedJSON string
	)

	{
		loc.File = "file"
		loc.Line = 10
		loc.Func = "fn"

		expectedJSON = `{"file":"file:10","func":"fn"}`
	}

	js, _ := json.Marshal(loc)
	actualJSON := string(js)

	assert.Equal(expectedJSON, actualJSON)
}

//
