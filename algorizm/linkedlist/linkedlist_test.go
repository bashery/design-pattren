//package linkedlist

package main

import (
	"testing"
)


func TestPrepend(t *testing.T) {
    list := LinkedList{}
    for i := cases.from; i >= cases.to; i-- {
        list.Prepend(&Node{Data: i})    
    }
}

var cases = struct {
    from int
    to int
}{
    from:10,
    to:20,
}
