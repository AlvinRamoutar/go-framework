package http

import (
	"path"
	"strings"
)

type Method byte

const (
	GET Method = iota
	HEAD
	POST
	PUT
	PATCH
	DELETE
	CONNECT
	OPTIONS
	TRACE
	STATUS
)

func (m Method) Byte() byte {
	switch m {
	case GET:
		return 0
	case HEAD:
		return 1
	case POST:
		return 2
	case PUT:
		return 3
	case PATCH:
		return 4
	case DELETE:
		return 5
	case CONNECT:
		return 6
	case OPTIONS:
		return 7
	case TRACE:
		return 8
	case STATUS:
		return 9
	}

	return 0
}

type PathType byte

const (
	EXACT PathType = iota
	REGEX
)

func (pt PathType) Byte() byte {
	switch pt {
	case EXACT:
		return 0
	case REGEX:
		return 1
	}

	return 0
}

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

type HttpError struct {
	code string
}

func (e *HttpError) Error() string {
	return e.code
}

type ResponseResult struct {
	ResponseCode int
	Error        error
}
