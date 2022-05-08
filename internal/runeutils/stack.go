package runeutils

type Stack struct {
	stack []rune
}

func (s *Stack) Push(r rune) {
	s.stack = append(s.stack, r)
}

func (s *Stack) Pop() rune {
	r := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return r
}
