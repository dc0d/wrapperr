package wrapperr

import (
	"fmt"
)

type Loc struct {
	Line int
	File string
	Func string
}

func (loc Loc) String() string { return fmt.Sprintf(shortFilePath(loc.File)+":%d "+loc.Func, loc.Line) }

func (loc Loc) MarshalJSON() ([]byte, error) {
	return []byte(`"` + loc.String() + `"`), nil
}
