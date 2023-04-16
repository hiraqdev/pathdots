package core

import "errors"

var (
	ErrRCRowLength = errors.New("row: out of range")
	ErrRCColLength = errors.New("col: out of range")
	ErrTargetPos   = errors.New("invalid target position")
	ErrStartPos    = errors.New("invalid start position")
)

type Stack struct {
	container []any
}

type RC struct {
	Row int
	Col int
}

type RCParser struct {
	input [][]any
}

type Scanner struct {
	input [][]any
	rc    *RCParser
	stack *Stack
}
