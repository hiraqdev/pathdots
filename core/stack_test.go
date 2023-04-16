package core_test

import (
	"amartha/pathdots/core"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StackTestSuite struct {
	suite.Suite
}

func (s *StackTestSuite) TestPushPeek() {
	stack := core.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	value := stack.Peek()
	s.Assert().Equal(3, value)
}

func (s *StackTestSuite) TestPop() {
	stack := core.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	value := stack.Pop()
	s.Assert().Equal(3, value)
	s.Assert().Equal(2, stack.Length())
}

func (s *StackTestSuite) TestTopLength() {
	stack := core.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	s.Assert().Equal(3, stack.Length())
	s.Assert().Equal(2, stack.Top())
}

func (s *StackTestSuite) TestIsEmpty() {
	stack := core.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	s.Assert().False(stack.IsEmpty())

	stack.Pop()
	stack.Pop()
	stack.Pop()
	s.Assert().True(stack.IsEmpty())
}

func TestStackTestSuite(t *testing.T) {
	suite.Run(t, new(StackTestSuite))
}
