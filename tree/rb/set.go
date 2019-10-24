package rb

type Set struct {
	Root *RBTree
}

func Less(a, b interface{}) bool {
	return a.(int) < b.(int)
}
func More(a, b interface{}) bool {
	return a.(int) > b.(int)
}
func Equal(a, b interface{}) bool {
	return a.(int) == b.(int)
}
func NewSet() *Set {
	return &Set{
		Root: NewRBTree(Less, More, Equal),
	}
}

func (s *Set) Add(i int) {
	s.Root.Add(i, i)
}
func (s *Set) Delete(i int) {
	s.Root.Remove(i)
}
func (s *Set) Has(i int) bool {
	return s.Root.Contains(i)
}
