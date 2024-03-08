package vlc

import (
	"strings"
	"unicode"
)

func VlcEncode(str string) string {
	return str
}

// prepareText prepares text for encoding
// changing all letters to lower case -> Hello = !hello
func prepareText(str string) string {
	var buf strings.Builder
	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}