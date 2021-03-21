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
		expectedString = "file:10 fn"
	)

	actualString := fmt.Sprint(loc)

	assert.Equal(t, expectedString, actualString)
}

func TestLoc_to_json(t *testing.T) {
	var (
		loc          = sampleLoc()
		expectedJSON = `"file:10 fn"`
	)

	js, err := json.Marshal(loc)
	if err != nil {
		require.NoError(t, err)
	}
	actualJSON := string(js)

	assert.Equal(t, expectedJSON, actualJSON)
}
