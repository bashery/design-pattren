package stack

import "testing"

func TestPop(t *testing.T) {
    s := Stack{}
    for i:=cases.from; i <= cases.to; i++ {
        s.Push(i)
    }

    for i:=cases.to; i >= cases.from; i-- {
        popt := s.Pop()
        if popt != i {
            t.Errorf("unexpected %d, expecting %d", popt, i)
        }
    }



}


var cases = struct{
    from int
    to int
}{
    from: 10,
    to: 20,
}
