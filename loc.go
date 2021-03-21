package wrapperr

import (
	"fmt"
)

type Loc struct {
	Line int    `json:"line,omitempty"`
	File string `json:"file,omitempty"`
	Func string `json:"func,omitempty"`
}

func (loc Loc) String() string { return fmt.Sprintf(loc.File+":%d "+loc.Func, loc.Line) }

func (loc Loc) MarshalJSON() ([]byte, error) {
	return []byte(`"` + loc.String() + `"`), nil
}
