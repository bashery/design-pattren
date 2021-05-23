package find

import (
	"testing"
)

func TestFind(t *testing.T) {
    
    for k, v := range cases {
        res := Find(v.text, v.subtext)
        if res != k {
            t.Errorf("find was incorrect, got: %d, want: %d.", res, k)
        }
    }

}

type Case struct {
    text string 
    subtext string
}
var cases = map[int]Case{
    6: {"hello world","world"},
    7: {"hello world","orld"},
    15: {"hello borld of world","world"},
    0: {"hello world","he"},
    -1: {"hello university","world"},
}
