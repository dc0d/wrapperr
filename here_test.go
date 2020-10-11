package here_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dc0d/here"

	"github.com/stretchr/testify/assert"
)

func TestLoc(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		var (
			assert         = assert.New(t)
			loc            here.Loc
			expectedString string
			actualString   string
		)

		{
			loc.File = "file"
			loc.Line = 10
			loc.Func = "fn"

			expectedString = "file:10 fn"
		}

		actualString = fmt.Sprint(loc)

		assert.Equal(expectedString, actualString)
	})

	t.Run("json", func(t *testing.T) {
		var (
			assert       = assert.New(t)
			loc          here.Loc
			expectedJSON string
			actualJSON   string
		)

		{
			loc.File = "file"
			loc.Line = 10
			loc.Func = "fn"

			expectedJSON = `{"file":"file:10","func":"fn"}`
		}

		js, _ := json.Marshal(loc)
		actualJSON = string(js)

		assert.Equal(expectedJSON, actualJSON)
	})

	t.Run("shorten file", func(t *testing.T) {
		var (
			assert         = assert.New(t)
			loc            here.Loc
			expectedString string
			actualString   string
		)

		{
			loc.File = "/path/to/file.go"
			loc.Line = 10
			loc.Func = "github.com/user/mod/v6/pkg.fn"

			expectedString = "file.go:10 github.com/user/mod/v6/pkg.fn"
		}

		loc.ShortenFile()
		actualString = fmt.Sprint(loc)

		assert.Equal(expectedString, actualString)
	})

	t.Run("shorten func", func(t *testing.T) {
		var (
			assert         = assert.New(t)
			loc            here.Loc
			expectedString string
			actualString   string
		)

		{
			loc.File = "/path/to/file.go"
			loc.Line = 10
			loc.Func = "github.com/user/mod/v6/pkg.fn"

			expectedString = "/path/to/file.go:10 pkg.fn"
		}

		loc.ShortenFunc()
		actualString = fmt.Sprint(loc)

		assert.Equal(expectedString, actualString)
	})
}

func TestCalls(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		var (
			assert         = assert.New(t)
			calls          here.Calls
			expectedString string
			actualString   string
		)

		{
			for i := 1; i <= 10; i++ {
				loc := here.Loc{
					File: fmt.Sprintf("file-%02d", i),
					Func: fmt.Sprintf("fn%02d", i),
					Line: i,
				}
				calls = append(calls, loc)
			}

			expectedString = "file-01:1 fn01 >\nfile-02:2 fn02 >\nfile-03:3 fn03 >\nfile-04:4 fn04 >\nfile-05:5 fn05 >\nfile-06:6 fn06 >\nfile-07:7 fn07 >\nfile-08:8 fn08 >\nfile-09:9 fn09 >\nfile-10:10 fn10"
		}

		actualString = fmt.Sprint(calls)

		assert.Equal(expectedString, actualString)
	})
}

func TestMark(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		var (
			assert          = assert.New(t)
			calls           here.Calls
			expectedStrings []string
			actualString    string
		)

		{
			calls = fn1()

			expectedStrings = []string{
				"here/here_call_fixture_test.go",
				"github.com/dc0d/here_test.fn1 >\n",
				"here/here_test.go",
				"github.com/dc0d/here_test.TestMark.func1 >\n",
				"testing.tRunner",
			}
		}

		actualString = fmt.Sprint(calls)

		for _, s := range expectedStrings {
			assert.Contains(actualString, s)
		}
	})

	t.Run("with short files", func(t *testing.T) {
		var (
			assert          = assert.New(t)
			calls           here.Calls
			expectedStrings []string
			actualString    string
		)

		{
			calls = fn2()

			expectedStrings = []string{
				"here_call_fixture_test.go",
				"github.com/dc0d/here_test.fn2 >\n",
				"here_test.go",
				"github.com/dc0d/here_test.TestMark.func2 >\n",
				"testing.tRunner",
			}
		}

		actualString = fmt.Sprint(calls)

		for _, s := range expectedStrings {
			assert.Contains(actualString, s)
		}
		assert.NotContains(actualString, "here/here_call_fixture_test.go")
		assert.NotContains(actualString, "here/here_test.go")
	})

	t.Run("with short funcs", func(t *testing.T) {
		var (
			assert          = assert.New(t)
			calls           here.Calls
			expectedStrings []string
			actualString    string
		)

		{
			calls = fn3()

			expectedStrings = []string{
				"here_call_fixture_test.go",
				"here_test.fn3 >\n",
				"here_test.go",
				"here_test.TestMark.func3 >\n",
				"testing.tRunner",
			}
		}

		actualString = fmt.Sprint(calls)

		for _, s := range expectedStrings {
			assert.Contains(actualString, s)
		}
		assert.NotContains(actualString, "github.com/dc0d/here_test.fn3")
		assert.NotContains(actualString, "github.com/dc0d/here_test.TestMark.func3")
	})

	t.Run("with skip", func(t *testing.T) {
		var (
			assert          = assert.New(t)
			calls           here.Calls
			expectedStrings []string
			actualString    string
		)

		{
			calls = fn5()

			expectedStrings = []string{
				"here_call_fixture_test.go",
				"here_test.fn5 >\n",
				"here_test.go",
				"testing.tRunner",
			}
		}

		actualString = fmt.Sprint(calls)

		for _, s := range expectedStrings {
			assert.Contains(actualString, s)
		}
		assert.NotContains(actualString, "fn4")
	})
}
