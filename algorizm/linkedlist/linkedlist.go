// linked list data structer
package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
    Lenght int
}

func (l *LinkedList) Prepend(n *Node) {
    secend := l.Head
     l.Head = n
     l.Head.Next= secend
     l.Lenght ++
}

func (l *LinkedList) PrintData() {
    toPrint := l.Head
    for l.Lenght != 0 {
        fmt.Printf("%d, %d\n", toPrint.Data, l.Lenght)
        toPrint = toPrint.Next
        l.Lenght --
    }
}

/*
func main() {
    list := LinkedList{}

    list.Prepend(&Node{Data:30})
    list.Prepend(&Node{Data:31})
    list.Prepend(&Node{Data:32})
    list.Prepend(&Node{Data:33})
    list.Prepend(&Node{Data:34})
    list.Prepend(&Node{Data:35})
    list.Prepend(&Node{Data:36})
    list.Prepend(&Node{Data:37})

    list.PrintData()
}
*/
