package queue

import "testing"


func TestQueue(t *testing.T) {
    q := Queue{}
    c := cases
    for  i := c.from; i <= c.to; i++ {
        q.Enqueue(i)
    }


    for  i := c.from; i <= c.to; i++ {
        removed := q.Dequeue()
        if removed != i {
            t.Errorf("expecting %d, not %d", i, removed )
            // or just print what hapend
        }
    }
}

var cases = struct {
    from int
    to int
}{
    from:10,
    to: 20,
}
