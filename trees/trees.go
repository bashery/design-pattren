package trees

import "fmt"

type Node struct {
    Key int
    Left *Node
    Right *Node
}


func (n *Node)Insert(k int) {
    if n.Key < k {
        if n.Right == nil {
            n.Right = &Node{Key:k}
        } else {
            n.Right.Insert(k)
        }
    } else if n.Key > k {
        if n.Left == nil {
            n.Left = &Node{Key:k}
        } else {
            n.Left.Insert(k)
        }

    }
}

// Search takes a key and return true if key exest in the trees

var count int
func (n *Node) Search(k int) bool {
    if n == nil {
        return false
    }
    if n.Key < k {
        // move to right
        return n.Right.Search(k)
    } else if n.Key > k {
        // move to left
        return n.Left.Search(k)
        
    }
    count ++
    return true
}

func main() {
    tree := &Node{}

    tree.Insert(100)
    tree.Insert(200)
    tree.Insert(370)
    tree.Insert(330)
    tree.Insert(200)
    tree.Insert(400)
    tree.Insert(50)
    tree.Insert(67)
    tree.Insert(32)
    tree.Insert(42)
    tree.Insert(52)
    tree.Insert(29)
    tree.Insert(25)

    fmt.Println("expected true : ", tree.Search(100), count)
    fmt.Println("expected false : ", tree.Search(237777), count)
    fmt.Println("expected false : ", tree.Search(77), count)
    fmt.Println("expected true : ", tree.Search(25), count)
    fmt.Println("expected true : ", tree.Search(200), count)
}
