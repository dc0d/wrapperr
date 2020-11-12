package wrapperr

import (
	"encoding/json"
	"fmt"
)

type Loc struct {
	Line int    `json:"line,omitempty"`
	File string `json:"file,omitempty"`
	Func string `json:"func,omitempty"`
}

func (loc Loc) String() string { return fmt.Sprintf(loc.File+":%d "+loc.Func, loc.Line) }

func (loc Loc) MarshalJSON() ([]byte, error) {
	var payload struct {
		File string `json:"file,omitempty"`
		Func string `json:"func,omitempty"`
	}
	payload.File = fmt.Sprintf(loc.File+":%d", loc.Line)
	payload.Func = loc.Func
	return json.Marshal(payload)
}
