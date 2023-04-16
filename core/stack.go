package core

func NewStack() *Stack {
	inputs := make([]any, 0)
	return &Stack{
		container: inputs,
	}
}

func (s *Stack) Top() int {
	return s.Length() - 1
}

func (s *Stack) Length() int {
	return len(s.container)
}

func (s *Stack) IsEmpty() bool {
	return s.Length() < 1
}

func (s *Stack) Push(value any) {
	s.container = append(s.container, value)
}

func (s *Stack) Peek() any {
	if s.IsEmpty() {
		return nil
	}

	index := len(s.container) - 1
	return s.container[index]
}

func (s *Stack) Pop() any {
	top := s.Top()
	temp := s.container[top]
	s.container = s.container[:top]
	return temp
}
