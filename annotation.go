package wrapperr

type Annotation struct {
	Loc     Loc
	Message string
}

func (note Annotation) String() string {
	var rest string
	if note.Message != "" {
		rest = " - " + note.Message
	}
	return note.Loc.String() + rest
}

func (note Annotation) MarshalJSON() ([]byte, error) {
	return []byte(`"` + note.String() + `"`), nil
}
