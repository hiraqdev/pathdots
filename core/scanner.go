package core

import (
	"fmt"
)

func NewScanner(input [][]any) *Scanner {
	return &Scanner{
		input: input,
		stack: NewStack(),
		rc:    NewRCParser(input),
	}
}

func (s *Scanner) TargetPos() *RC {
	rowIndex := len(s.input) - 1
	columns := s.input[rowIndex]
	colIndex := len(columns) - 1

	return &RC{
		Row: rowIndex,
		Col: colIndex,
	}
}

func (s *Scanner) StartPos() *RC {
	return &RC{
		Row: 0,
		Col: 0,
	}
}

func (s *Scanner) ValidateTargetPos(pos *RC) error {
	value := s.input[pos.Row][pos.Col]

	if value != 0 {
		return ErrTargetPos
	}

	return nil
}

func (s *Scanner) ValidateStartPos(pos *RC) error {
	value := s.input[pos.Row][pos.Col]
	if value != 0 {
		return ErrStartPos
	}

	return nil
}

func (s *Scanner) Scan(startPos, targetPos *RC) (bool, error) {
	currentPos := startPos

	stopper := false
	hasPath := false
	posHistory := make(map[string][]string)

	// push currentPos that our first starting post
	// to the stack to initialize process
	s.stack.Push(currentPos.String())

	for !stopper && !s.stack.IsEmpty() {
		columns := s.rc.ColumnsAt(currentPos.Row)
		columnMax := len(columns) - 1

		rightPos := currentPos.MoveRight(columnMax)
		bottomPos := currentPos.MoveDown(len(s.input) - 1)

		rightPosIndexed := false
		bottomPosIndexed := false

		// need to check if current right and bottom pos has been indexed or not
		// if both position has already indexed , then we just need to pop current position
		// and continue the iteration to the next current position
		if indexed, exists := posHistory[currentPos.String()]; exists {
			rightPosStr := rightPos.String()
			bottomPosStr := bottomPos.String()

			for _, val := range indexed {
				if val == rightPosStr {
					rightPosIndexed = true
				}

				if val == bottomPosStr {
					bottomPosIndexed = true
				}
			}
		}

		// if both righ and bottom pos has already indexed then just forget it
		// pop it from stack and continue the iteration
		// doesn't need to continue to next process
		if rightPosIndexed && bottomPosIndexed {
			s.stack.Pop()
			if !s.stack.IsEmpty() {
				topPosFromStack := s.stack.Peek()
				topPos, err := RCFromString(fmt.Sprintf("%s", topPosFromStack))
				if err != nil {
					return false, err
				}

				currentPos = topPos
				continue
			} else {
				// if current stack is empty then just break iteration
				break
			}
		}

		// if our righPos or bottomPos is equal with our targetPos it means we've just found
		// a path to the target itself, stop the iteration and mark hasPath as true
		if rightPos.String() == targetPos.String() || bottomPos.String() == targetPos.String() {
			stopper = true
			hasPath = true
			break
		}

		valueBottom := -1
		valueRight := -1

		if bottomPos.String() != currentPos.String() {
			// we need to prioritize our bottomPos
			// if our bottomPos has a value "0" it means we choose it and put it to the stack
			valueBottom, err := s.rc.Pos(bottomPos.Row, bottomPos.Col)
			if err != nil {
				return false, err
			}

			if valueBottom == 0 && !bottomPosIndexed {
				prevHistories := posHistory[currentPos.String()]
				prevHistories = append(prevHistories, bottomPos.String())
				posHistory[currentPos.String()] = prevHistories

				// need to change currentPos and continue to next iteration
				currentPos = bottomPos
				s.stack.Push(bottomPos.String())
				continue
			}
		}

		if rightPos.String() != currentPos.String() {
			// if our bottom value doesnt match with 0, then we need to looking for our rightPos
			// if our rightPos has a value "0" we put it to the stack
			valueRight, err := s.rc.Pos(rightPos.Row, rightPos.Col)
			if err != nil {
				return false, err
			}

			if valueRight == 0 && !rightPosIndexed {
				prevHistories := posHistory[currentPos.String()]
				prevHistories = append(prevHistories, rightPos.String())
				posHistory[currentPos.String()] = prevHistories

				// need to change currentPos and continue to next iteration
				currentPos = rightPos
				s.stack.Push(rightPos.String())
				continue
			}
		}

		// if both value bottom and value right has a value with -1 it means we need to forget
		// current pos and pop it from stack
		if valueBottom == -1 && valueRight == -1 {
			s.stack.Pop()
			topPosFromStack := s.stack.Peek()

			// if current top position is nil, just break the iteration
			if topPosFromStack == nil {
				break
			}

			topPos, err := RCFromString(fmt.Sprintf("%s", topPosFromStack))
			if err != nil {
				return false, err
			}

			currentPos = topPos
			continue
		}
	}

	return hasPath, nil
}
