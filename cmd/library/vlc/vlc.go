package vlc

import (
	"encoding"
	"strings"
	"unicode"
)

func Encode(str string) string {
	return ""
}

// prepareText prepares text for encoding
// changing all letters to lower case -> Hello = !hello
func prepareText(str string) string {
	var buf strings.Builder
	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

func encodeBin(str string) string {
	var buf strings.Builder
	for _, ch := range str {
		buf.WriteRune(bin(ch))
	}
	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable(ch)

	result, ok := table(ch)
	if !ok {
		panic("unknown char" + string(ch))
	}
	return result
}

func getEncodingTable() encodingTable {
	return getEncodedTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}