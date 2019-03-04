package utils

import "strconv"

const (
	encodePath encoding = 1 + iota
	encodeHost
	encodeUserPassword
	encodeQueryComponent
	encodeFragment
)

type encoding int
type EscapError string

func (e EscapError) Error() string {
	return "invalid URL escape " + strconv.Quote(string(e))
}

func ishex(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}
	return false
}
