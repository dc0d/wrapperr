package wrapperr_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoc_to_string(t *testing.T) {
	var (
		loc            = sampleLoc()
		expectedString = "package/file.go:9 github.com/user/module/package.(*Struct).method"
	)

	actualString := fmt.Sprint(loc)

	assert.Equal(t, expectedString, actualString)
}

func TestLoc_to_json(t *testing.T) {
	var (
		loc          = sampleLoc()
		expectedJSON = `"package/file.go:9 github.com/user/module/package.(*Struct).method"`
	)

	js, err := json.Marshal(loc)
	if err != nil {
		require.NoError(t, err)
	}
	actualJSON := string(js)

	assert.Equal(t, expectedJSON, actualJSON)
}
