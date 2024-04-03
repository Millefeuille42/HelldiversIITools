package utils

import "regexp"

var re = regexp.MustCompile("<[^>]*>")

func StripMarkup(text string) string {
	return re.ReplaceAllString(text, "**")
}
