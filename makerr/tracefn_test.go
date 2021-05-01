package makerr_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dc0d/wrapperr/makerr"
	"github.com/stretchr/testify/assert"
)

func Test_TraceFn_uses_same_location(t *testing.T) {
	t.Parallel()

	var (
		onEnterLoc  string
		onExitLoc   string
		elapsedTime time.Duration
	)

	onEnter := func(loc string) { onEnterLoc = loc }
	onExit := func(elapsed time.Duration, loc string) { elapsedTime, onExitLoc = elapsed, loc }

	makerr.TraceFn(onEnter)(onExit)

	assert.Equal(t, onEnterLoc, onExitLoc)
	assert.True(t, elapsedTime > 0)
}

func Test_TraceFn_using_defer(t *testing.T) {
	t.Parallel()

	var (
		onEnterLoc  string
		onExitLoc   string
		elapsedTime time.Duration
	)

	onEnter := func(loc string) { onEnterLoc = loc }
	onExit := func(elapsed time.Duration, loc string) { elapsedTime, onExitLoc = elapsed, loc }

	checkScope(onEnter, onExit)

	assert.Equal(t, "makerr/fixtures_test.go:23 github.com/dc0d/wrapperr/makerr_test.checkScope", onEnterLoc)
	assert.Equal(t, onEnterLoc, onExitLoc)
	assert.True(t, elapsedTime > time.Millisecond*50)
}

func ExampleTraceFn() {
	func() {
		defer makerr.TraceFn(func(loc string) {
			// logging it, or any other required antion on entry
			fmt.Println(loc)
		})(func(elapsed time.Duration, loc string) {
			// logging it, or any other required antion on exit
			fmt.Println(loc, elapsed > 0)
		})

		time.Sleep(time.Millisecond * 10)
	}()

	// Output:
	// makerr/tracefn_test.go:51 github.com/dc0d/wrapperr/makerr_test.ExampleTraceFn.func1
	// makerr/tracefn_test.go:51 github.com/dc0d/wrapperr/makerr_test.ExampleTraceFn.func1 true
}
