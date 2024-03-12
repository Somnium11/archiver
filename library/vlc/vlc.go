package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const chunkSize = 8

type BinaryChunks []BinaryChunk

type BinaryChunk string

type HexChunk string

type HexChunks []HexChunk

type encodingTable map[rune]string

// Encode is a Go function that takes a string and returns a string.
// It takes a string parameter and does not return anything.
func Encode(str string) string {
	str = prepareText(str)

	chunks := splitByChunks(encodeBin(str), chunkSize)

	fmt.Println(chunks)

	return chunks.ToHex().ToString()
}

func (hcs HexChunks) ToString() string {
	const separatop = " "
	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hcs[0]))

	for _, hc := range hcs[1:] {
		buf.WriteString(separatop)
		buf.WriteString(string(hc))
	}
	return buf.String()
}

// ToHex converts the BinaryChunks to HexChunks
func (bcs BinaryChunks) ToHex() HexChunks {
	result := make(HexChunks, 0, len(bcs))

	for _, chunk := range bcs {
		hexChunk := chunk.ToHex()
		result = append(result, hexChunk)
	}
	return result
}

func (bc BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bc), 2, chunkSize)
	if err != nil {
		panic("cant parse binary chunk:" + err.Error())
	}
	res := strings.ToUpper(fmt.Sprintf("%x", num))
	if len(res) == 1 {
		res = "0" + res
	}
	return HexChunk(res)
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

// splitByChunks splits binary string by chunks with given size,
// '100101101011100010100011' -> '10010110  10111000 10100011'
func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunksCount := utf8.RuneCountInString(bStr) / chunkSize

	if strLen/chunkSize != 0 {
		chunksCount += 1
	}
	result := make(BinaryChunks, 0, chunksCount)
	var buf strings.Builder
	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			result = append(result, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}
	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))
		result = append(result, BinaryChunk(lastChunk))
	}
	return result
}

// encodeBin encodes string to binary without spaces
func encodeBin(str string) string {
	var buf strings.Builder
	for _, ch := range str {
		buf.WriteString(bin(ch))
	}
	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()
	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}
	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
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
