package trees

import (
	"testing"
)

func TestSearch(t *testing.T) {
    
    var node = &Node{}
    for _,v := range keys {
        node.Insert(v)
    } 


    for k , v := range cases {
        res := node.Search(k)
        if res != v {
            t.Errorf("expexted %t, find %t with %d case", v, res, k )
        }
    }
}


var keys = [...]int{2,34,5, 6, 7, 55, 43, 21, 45, 789, 12, 13, 495, 467, 34}

var cases = map[int]bool{
    2: true,
    5: true,
    65: false,
    333: false,
}
