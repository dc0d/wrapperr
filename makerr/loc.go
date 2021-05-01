package makerr

import (
	"fmt"
)

type loc struct {
	Line int    `json:"line,omitempty"`
	File string `json:"file,omitempty"`
	Func string `json:"func,omitempty"`
}

func (obj loc) String() string { return fmt.Sprintf(obj.File+":%d "+obj.Func, obj.Line) }
