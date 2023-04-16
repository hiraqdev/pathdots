package core

import (
	"fmt"
	"strconv"
	"strings"
)

func NewRCParser(input [][]any) *RCParser {
	return &RCParser{
		input: input,
	}
}

func (rc *RCParser) Pos(row, col int) (any, error) {
	if row > len(rc.input)-1 {
		return nil, ErrRCRowLength
	}

	columns := rc.input[row]
	if col > len(columns)-1 {
		return nil, ErrRCColLength
	}

	return columns[col], nil
}

func (rc *RCParser) ColumnsAt(row int) []any {
	return rc.input[row]
}

func RCFromString(val string) (*RC, error) {
	val = strings.Replace(val, "(", "", -1)
	val = strings.Replace(val, ")", "", -1)
	splitted := strings.Split(val, ",")
	rowInt, err := strconv.Atoi(splitted[0])
	if err != nil {
		return nil, err
	}

	colInt, err := strconv.Atoi(splitted[1])
	if err != nil {
		return nil, err
	}

	return &RC{
		Row: rowInt,
		Col: colInt,
	}, nil
}

func (pos *RC) String() string {
	return fmt.Sprintf("(%d,%d)", pos.Row, pos.Col)
}

func (pos *RC) MoveRight(maxCol int) *RC {
	var nextCol int
	if pos.Col+1 > maxCol {
		nextCol = pos.Col
	} else {
		nextCol = pos.Col + 1
	}

	return &RC{
		Row: pos.Row,
		Col: nextCol,
	}
}

func (pos *RC) MoveDown(maxRow int) *RC {
	var nextRow int
	if pos.Row+1 > maxRow {
		nextRow = pos.Row
	} else {
		nextRow = pos.Row + 1
	}

	return &RC{
		Row: nextRow,
		Col: pos.Col,
	}
}
