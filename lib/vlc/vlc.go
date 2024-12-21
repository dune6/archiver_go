package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk
type BinaryChunk string
type encodingTable map[rune]string
type HexChunk string
type HexChunks []HexChunk

const chunkSize = 8

func Encode(str string) string {
	str = prepareText(str)

	bStr := encodeBin(str)

	chunks := splitByChunks(bStr, chunkSize)

	encodedChunks := chunks.ToHex()
	fmt.Println(encodedChunks)

	return encodedChunks.ToString()
}

func (binChunks HexChunks) ToString() string {
	var strBuilder strings.Builder
	for _, chunk := range binChunks {
		strBuilder.WriteString(string(chunk) + " ")
	}
	return strBuilder.String()
}

func (binChunks BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(binChunks))
	for _, chunk := range binChunks {
		hexChunk := chunk.ToHexChunk()
		res = append(res, hexChunk)
	}
	return res
}

func (binChunk BinaryChunk) ToHexChunk() HexChunk {
	num, err := strconv.ParseUint(string(binChunk), 2, chunkSize)
	if err != nil {
		panic(err)
	}
	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}
	return HexChunk(res)
}

func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunksCount := strLen / chunkSize

	if strLen%chunkSize != 0 {
		chunksCount++
	}

	chunks := make(BinaryChunks, 0, chunksCount)

	var builder strings.Builder
	for i, ch := range bStr {
		builder.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			chunks = append(chunks, BinaryChunk(builder.String()))
			builder.Reset()
		}
	}

	if builder.Len() != 0 {
		lastChunk := builder.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		chunks = append(chunks, BinaryChunk(lastChunk))
	}

	return chunks
}

// encodeBin encode string to binary codes string without separates
func encodeBin(str string) string {
	var builder strings.Builder
	for _, ch := range str {
		builder.WriteString(bin(ch))
	}
	return builder.String()
}

func bin(r rune) string {
	table := getEncodingTable()
	res, ok := table[r]
	if !ok {
		panic("Unknown encoding rune" + res)
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

// prepareText prepares text to be fit for encode
// changes upper case letters to lower by the rule: A -> !a
func prepareText(str string) string {
	var builder strings.Builder
	for _, ch := range str {
		if unicode.IsUpper(ch) {
			builder.WriteRune('!')
			builder.WriteRune(unicode.ToLower(ch))
		} else {
			builder.WriteRune(ch)
		}
	}

	return builder.String()
}
