package linkedlist

import (
	"testing"
)


func TestPrepend(t *testing.T) {
    list := LinkedList{}
    for i := cases.from; i <= cases.to; i++ {
        list.Prepend(&Node{Data: i})    
    }
    
    for i := cases.from; i <= cases.to; i++ {


    toPrint := list.Head
    // if list.Lenght != 0 {
    t.Errorf("%d, %d\n", toPrint.Data, list.Lenght)
    toPrint = toPrint.Next
    list.Lenght --
    //}
}
}

var cases = struct {
    from int
    to int
}{
    from:10,
    to:20,
}
