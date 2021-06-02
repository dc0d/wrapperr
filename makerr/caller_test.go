package makerr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetCaller_exact(t *testing.T) {
	t.Parallel()

	result := returnsExactLocation()

	assert.Contains(t, result, "github.com/dc0d/wrapperr/makerr_test.returnsExactLocation")
	assert.Contains(t, result, "makerr/fixtures_test.go")
	assert.Contains(t, result, ":8")
}

func Test_GetCaller_caller(t *testing.T) {
	t.Parallel()

	result := returnsCallerLocation()

	assert.Contains(t, result, "github.com/dc0d/wrapperr/makerr_test.returnsCallerLocation")
	assert.Contains(t, result, "makerr/fixtures_test.go")
	assert.Contains(t, result, ":17")
}
