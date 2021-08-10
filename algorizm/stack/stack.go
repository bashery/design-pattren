package stack

type Stack struct {
	Items []int
}

func (s *Stack) Push(i int) {
	s.Items = append(s.Items, i)
}

func (s *Stack) Pop() int {
    l := len(s.Items)
    if l < 1 {
        return -1 // this resturn is not correct becoase -1 mybe item in stack
    }

    toRemove :=  s.Items[l-1]
    s.Items = s.Items[:l-1]
    return toRemove 

}
/*
func main() {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
}
*/
