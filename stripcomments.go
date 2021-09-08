package stripcomments

import "regexp"

var reLineCommentStartsWithDoubleSlash = regexp.MustCompile(`(?s)//.*?\n`)

type Stripper interface {
	Strip([]byte) []byte
}

type engine struct {
	strippers []Stripper
}

func NewStripper(opt ...option) Stripper {
	engine := &engine{
		strippers: make([]Stripper, 0),
	}
	return engine
}

func (e *engine) AddStripper(s Stripper) {
	e.strippers = append(e.strippers, s)
}

func (e *engine) Strip(source []byte) []byte {
	for _, stripper := range e.strippers {
		stripper.Strip(source)
	}
	return source // inplace?
}
