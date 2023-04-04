package main

type Stack struct {
	elements []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return value
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
