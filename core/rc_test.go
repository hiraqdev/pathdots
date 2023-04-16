package core_test

import (
	"amartha/pathdots/core"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RCTestSuite struct {
	suite.Suite
}

func (rct *RCTestSuite) TestPos() {
	column1 := []interface{}{0, 0, -1}
	column2 := []interface{}{-1, 0, 0}
	inputs := [][]interface{}{column1, column2}
	rc := core.NewRCParser(inputs)

	value1, _ := rc.Pos(0, 0)
	rct.Assert().Equal(0, value1)

	value2, _ := rc.Pos(0, 1)
	rct.Assert().Equal(0, value2)

	value3, _ := rc.Pos(0, 2)
	rct.Assert().Equal(-1, value3)

	value4, _ := rc.Pos(1, 0)
	rct.Assert().Equal(-1, value4)

	value5, _ := rc.Pos(1, 1)
	rct.Assert().Equal(0, value5)

	value6, _ := rc.Pos(1, 2)
	rct.Assert().Equal(0, value6)
}

func (rct *RCTestSuite) TestPosErrors() {
	column1 := []interface{}{0, 0, -1}
	inputs := [][]interface{}{column1}

	rc := core.NewRCParser(inputs)
	_, err := rc.Pos(5, 0)
	rct.Assert().Error(err)
	rct.Assert().ErrorIs(err, core.ErrRCRowLength)

	_, err = rc.Pos(0, 3)
	rct.Assert().Error(err)
	rct.Assert().ErrorIs(err, core.ErrRCColLength)
}

func (rct *RCTestSuite) TestPosFromString() {
	val := "(0,1)"
	rc, err := core.RCFromString(val)
	rct.Assert().NoError(err)
	rct.Assert().Equal(0, rc.Row)
	rct.Assert().Equal(1, rc.Col)
}

func (rct *RCTestSuite) TestPosFromStringErrorRow() {
	val := "(invalid,1)"
	rc, err := core.RCFromString(val)
	rct.Assert().Nil(rc)
	rct.Assert().Error(err)
}

func (rct *RCTestSuite) TestPosFromStringErrorColumn() {
	val := "(0,column)"
	rc, err := core.RCFromString(val)
	rct.Assert().Nil(rc)
	rct.Assert().Error(err)
}

func TestRCTestSuite(t *testing.T) {
	suite.Run(t, new(RCTestSuite))
}
