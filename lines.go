package stripcomments

import "regexp"

type lineCommentStripper struct {
	re *regexp.Regexp
}

func (s *lineCommentStripper) Strip(content []byte) []byte {
	// TODO
	return []byte(nil)
}
