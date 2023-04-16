# Amartha Test 

Chosen domain problem: `find path between 2 dots`

## Problem

Given a 2D array(m x n). The task is to check if there is any path from top left to bottom right. In the matrix, -1 is considered as blockage (can’t go through this cell) and 0 is considered path cell (can go through it)

> Top left cell always contains 0

Examples 

```
Input : arr[][] = { { 0, 0, 0, -1, 0},
                    {-1, 0, 0, -1, -1},
                    { 0, 0, 0, -1, 0},
                    {-1, 0, 0, 0, 0},
                    { 0, 0, -1, 0, 0}}
Output : Yes
```

## Analyze

Reading the input data like `row` and `column` , example:

```shell
	   0		  1			 2		   3		   4
0  (r0, c0:0)  (r0,c1:0)  (r0,c2:0)  (r0,c3:-1)  (r0,c4:0)
1  (r1, c0:-1) (r1,c1:0)  (r1,c2:0)  (r1,c3:-1)  (r1,c4:-1)
2  (r2, c0:0)  (r2,c1:0)  (r2,c2:0)  (r2,c3:-1)  (r2,c4:0)
3  (r3, c0:-1) (r3,c1:0)  (r3,c2:0)  (r3,c3:0)   (r3,c4:0)
4  (r4, c0:0)  (r4,c1:0)  (r4,c2:-1) (r4,c3:0)   (r4,c4:0)
```

I've just modeling our data into *row* and *column* based to give us clear context how to read the data and move our `Scanner`

Using [[Data Structure - Depth First Traversal]], try to check if given input has a path from top left to bottom right

Imagine we have to create a `Scanner` used to read and parse given input

- A `Scanner` is a program to parse input and scan from top left position , and moving to scan until it's found a path to bottom right (at least one path)
- The reason I'm using `DFS` traversal algorithm because our `Scanner` need a `stack` to remember each of visited node, and if our `Scanner` met some blockage (-1), we just need to pop our stack value

## Algorithm

- Count how many rows from input's length
- Initialize stack's variable
- Initialize stopper variable, this variable is a flag to terminate our `Scanner` iteration
- Initialize `currentPos`: (r0,c0) that must has `0` value
	- Put `(r0, c0)` to the stack
- Initialize target position: `(r4, c4)`
- Initialize `hasPath` to `false`, this variable indicate that given input has a path or not
- Initialize `postHistory` which is a dictionary/map to remember chosen position from `currentPos`, this variable used to not select next node position which was chosen before
- Start to iterate while `stopper` not true
	- `righPos` = current column + 1 -> (r0, c1) -> `0`
		- if our `rightPos` is equal with `(r4,c4)` then stop the iteration and set `hasPath` to `true`
	- `bottomPos` = current row + 1 ->  (r1, c0) -> `-1`
		- because `-1` is blockage we just ignore it
		- if our `bottomPos` is equal with `(r4,c4)` then stop the iteration and set `hasPath` to `true`
	- if `rightPos` and `bottomPos` both has `-1` value then we just forget current node, and pop it from our stack
		- if stack goes to empty, then we need to stop the iteration by set our `stopper` to true
	- next `currentPos` is : `(r0, c1)` and push it to the stack and continue iteration
		- If we have `rightPos` and `bottomPos` with `0` value we need to prioritize `bottomPos` why? because our `targetPos` is at the very bottom ( `r4` )

## Implementation

Taken from `core/scanner.go` 

```go
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
```

# Testing

```
make test
```

Test cases

```go
	var input1 [][]interface{}
	var input2 [][]interface{}
	var input3 [][]interface{}

	columnsInput1 := make([]interface{}, 0)
	columns1Input1 := make([]interface{}, 0)

	columnsInput1 = append(columnsInput1, 0, 0)
	columns1Input1 = append(columns1Input1, -1, 0)

	input1 = append(input1, columnsInput1, columns1Input1)

	// input 2 : more broader data set but still has a path

	columnsInput2 := make([]interface{}, 0)
	columns1Input2 := make([]interface{}, 0)
	columns2Input2 := make([]interface{}, 0)
	columns3Input2 := make([]interface{}, 0)
	columns4Input2 := make([]interface{}, 0)

	columnsInput2 = append(columnsInput2, 0, 0, 0, -1, 0)
	columns1Input2 = append(columns1Input2, -1, 0, 0, -1, -1)
	columns2Input2 = append(columns2Input2, 0, 0, 0, -1, 0)
	columns3Input2 = append(columns3Input2, -1, 0, 0, 0, 0)
	columns4Input2 = append(columns4Input2, 0, 0, -1, 0, 0)

	input2 = append(input2, columnsInput2, columns1Input2, columns2Input2, columns3Input2, columns4Input2)

	// input 3 : it still has a path but different path

	columnsInput3 := make([]interface{}, 0)
	columns1Input3 := make([]interface{}, 0)
	columns2Input3 := make([]interface{}, 0)
	columns3Input3 := make([]interface{}, 0)
	columns4Input3 := make([]interface{}, 0)

	columnsInput3 = append(columnsInput3, 0, 0, 0, -1, 0)
	columns1Input3 = append(columns1Input3, -1, 0, 0, -1, -1)
	columns2Input3 = append(columns2Input3, 0, 0, 0, -1, 0)
	columns3Input3 = append(columns3Input3, -1, 0, 0, 0, 0)
	columns4Input3 = append(columns4Input3, 0, 0, -1, -1, 0)

	input3 = append(input3, columnsInput3, columns1Input3, columns2Input3, columns3Input3, columns4Input3)

	inputs := make([][][]interface{}, 0)
	inputs = append(inputs, input1, input2, input3)
```

Negative test cases, it means an input that doesn't have any paths

```go
	var input2 [][]interface{}
	columnsInput2 := make([]interface{}, 0)
	columns1Input2 := make([]interface{}, 0)
	columns2Input2 := make([]interface{}, 0)
	columns3Input2 := make([]interface{}, 0)
	columns4Input2 := make([]interface{}, 0)

	columnsInput2 = append(columnsInput2, 0, 0, 0, -1, 0)
	columns1Input2 = append(columns1Input2, -1, 0, 0, -1, -1)
	columns2Input2 = append(columns2Input2, 0, 0, 0, -1, 0)
	columns3Input2 = append(columns3Input2, -1, 0, 0, -1, 0)
	columns4Input2 = append(columns4Input2, 0, 0, -1, -1, 0)

	input2 = append(input2, columnsInput2, columns1Input2, columns2Input2, columns3Input2, columns4Input2)
```