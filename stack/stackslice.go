// slice实现stack
package stack

type Stack2 struct {
	array []interface{}
}

func NewStack2() *Stack2 {
	return &Stack2{}
}

func (s *Stack2) Push(v interface{}) {
	s.array = append(s.array, v)
}

func (s *Stack2) Pop() interface{} {
	if len(s.array) > 0 {
		item := s.array[len(s.array)-1]
		s.array = s.array[0 : len(s.array)-1]
		return item
	}
	return nil
}
func (s *Stack2) Top() interface{} {
	if len(s.array) > 0 {
		item := s.array[len(s.array)-1]
		return item
	}
	return nil
}
func (s *Stack2) Len() int {
	return len(s.array)
}
