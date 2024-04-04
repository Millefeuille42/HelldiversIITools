package utils

import "regexp"

var re = regexp.MustCompile("<[^>]*>")

func StripMarkup(text, replacement string) string {
	return re.ReplaceAllString(text, replacement)
}
