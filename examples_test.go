package wrapperr_test

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"

	"github.com/dc0d/wrapperr"
)

func show(err error) {
	terr, ok := err.(wrapperr.TracedErr)
	if !ok {
		fmt.Println(err)
		return
	}

	var stack wrapperr.Stack
	for _, v := range terr.Stack {
		if !strings.Contains(v.Loc.Func, "wrapperr") {
			v.Loc.Line = 0
			v.Loc.File = "..."
			v.Loc.Func = "..."
			v.Message = "rest of stack"
			stack = append(stack, v)
			break
		}
		v.Loc.File = ".../" + path.Base(v.Loc.File)
		stack = append(stack, v)
	}
	terr.Stack = stack

	fmt.Println(terr)
}

func showJSON(err error) {
	terr, ok := err.(wrapperr.TracedErr)
	if !ok {
		fmt.Println(err)
		return
	}

	var stack wrapperr.Stack
	for _, v := range terr.Stack {
		if !strings.Contains(v.Loc.Func, "wrapperr") {
			v.Loc.Line = 0
			v.Loc.File = "..."
			v.Loc.Func = "..."
			v.Message = "rest of stack"
			stack = append(stack, v)
			break
		}
		v.Loc.File = ".../" + path.Base(v.Loc.File)
		stack = append(stack, v)
	}
	terr.Stack = stack

	js, err := json.MarshalIndent(terr, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(js))
}
