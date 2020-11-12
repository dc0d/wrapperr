package wrapperr

type Annotation struct {
	Loc     Loc    `json:"loc"`
	Message string `json:"message,omitempty"`
}

func (note Annotation) String() string {
	var rest string
	if note.Message != "" {
		rest = " - " + note.Message
	}
	return note.Loc.String() + rest
}
