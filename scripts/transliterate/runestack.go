package main

type RuneStack struct {
	stack []rune
}

func (s *RuneStack) Push(r rune) {
	s.stack = append(s.stack, r)
}

func (s *RuneStack) Pop() rune {
	r := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return r
}
