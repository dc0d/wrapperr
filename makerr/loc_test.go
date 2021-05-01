package makerr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoc_to_string(t *testing.T) {
	t.Parallel()

	var (
		aLoc           = sampleLoc()
		expectedString = "file:10 fn"
	)

	actualString := fmt.Sprint(aLoc)

	assert.Equal(t, expectedString, actualString)
}

func sampleLoc() loc {
	var result loc
	result.File = "file"
	result.Line = 10
	result.Func = "fn"

	return result
}
