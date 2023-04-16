package core_test

import (
	"amartha/pathdots/core"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ScannerTestSuite struct {
	suite.Suite
}

func (s *ScannerTestSuite) TestTargetPosAndValidate() {
	var input [][]interface{}

	columns := make([]interface{}, 0)
	columns1 := make([]interface{}, 0)

	columns = append(columns, 0, 0)
	columns1 = append(columns1, -1, 0)

	input = append(input, columns, columns1)
	scanner := core.NewScanner(input)
	targetPos := scanner.TargetPos()
	err := scanner.ValidateTargetPos(targetPos)
	s.Assert().NoError(err)
	s.Assert().Equal(0, input[targetPos.Row][targetPos.Col])
	s.Assert().Equal("(1,1)", targetPos.String())
}

func (s *ScannerTestSuite) TestTargetPosAndValidateError() {
	var input [][]interface{}

	columns := make([]interface{}, 0)
	columns1 := make([]interface{}, 0)

	columns = append(columns, 0, 0)
	columns1 = append(columns1, -1, -1)

	input = append(input, columns, columns1)
	scanner := core.NewScanner(input)
	targetPos := scanner.TargetPos()
	err := scanner.ValidateTargetPos(targetPos)
	s.Assert().Error(err)
}

func (s *ScannerTestSuite) TestStartPosAndValidate() {
	var input [][]interface{}

	columns := make([]interface{}, 0)
	columns1 := make([]interface{}, 0)

	columns = append(columns, 0, 0)
	columns1 = append(columns1, -1, 0)

	input = append(input, columns, columns1)
	scanner := core.NewScanner(input)
	startPos := scanner.StartPos()
	err := scanner.ValidateStartPos(startPos)
	s.Assert().NoError(err)
	s.Assert().Equal(0, input[startPos.Row][startPos.Col])
	s.Assert().Equal("(0,0)", startPos.String())
}

func (s *ScannerTestSuite) TestStartPosAndValidateError() {
	var input [][]interface{}

	columns := make([]interface{}, 0)
	columns1 := make([]interface{}, 0)

	columns = append(columns, -1, 0)
	columns1 = append(columns1, -1, 0)

	input = append(input, columns, columns1)
	scanner := core.NewScanner(input)
	startPos := scanner.StartPos()
	err := scanner.ValidateStartPos(startPos)
	s.Assert().Error(err)
}

func (s *ScannerTestSuite) TestScan() {
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

	for _, input := range inputs {
		scanner := core.NewScanner(input)
		startPos := scanner.StartPos()
		targetPos := scanner.TargetPos()

		hasPath, err := scanner.Scan(startPos, targetPos)
		s.Assert().NoError(err)
		s.Assert().True(hasPath)
	}
}

func (s *ScannerTestSuite) TestScanNoPath() {
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
	scanner := core.NewScanner(input2)
	startPos := scanner.StartPos()
	targetPos := scanner.TargetPos()

	hasPath, err := scanner.Scan(startPos, targetPos)
	s.Assert().NoError(err)
	s.Assert().False(hasPath)
}

func TestScannerTestSuite(t *testing.T) {
	suite.Run(t, new(ScannerTestSuite))
}
