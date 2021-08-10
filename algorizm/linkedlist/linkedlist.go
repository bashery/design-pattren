// linked list data structer
//package linkedlist
package main

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

func (l LinkedList) PrintData() {
    head := l.Head
    
    for l.Lenght != 0 {
        fmt.Printf("%d, ", head.Data)
        head = head.Next
        l.Lenght--
    }
    fmt.Println()
}

func (list *LinkedList) Delete(value int) {
    if list.Lenght == 0 {
        return
    }

    previousToDelete := list.Head
    for previousToDelete.Next.Data != value {
        if previousToDelete.Next.Next == nil {
            return
        }
        previousToDelete = previousToDelete.Next
    }
    previousToDelete.Next = previousToDelete.Next.Next
    list.Lenght--
} 

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
    list.Delete(2220)
    list.PrintData()
}
