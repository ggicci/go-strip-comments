package stripcomments

import (
	"fmt"
	"regexp"
)

type option func(*engine)

func StripLineCommentsStartsWith(prefix string) option {
	re := regexp.MustCompile(fmt.Sprintf(`(?s)%s.*?\n`, prefix))
	return func(e *engine) {
		e.AddStripper(&lineCommentStripper{re: re})
	}
}
