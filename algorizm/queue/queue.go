package queue

type Queue struct {
	Items []int
}

func (q *Queue) Enqueue(i int) {
	q.Items = append(q.Items, i)
}

func (q *Queue) Dequeue() int {
	if len(q.Items) < 1 {
		return -1 // mybe -1 is item in queue
	}
	toRemove := q.Items[0]
	q.Items = q.Items[1:]
	return toRemove
}

/*
func main() {
    q := Queue{}

    q.Enqueue(0)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(-1)
    q.Enqueue(3)
    q.Enqueue(4)

    fmt.Println(q)
    fmt.Println(q.Dequeue())
    fmt.Println(q)
    fmt.Println(q.Dequeue())
    fmt.Println(q)
    fmt.Println(q.Dequeue())
    fmt.Println(q)
    fmt.Println(q.Dequeue())
    fmt.Println(q)
    fmt.Println(q.Dequeue())
    fmt.Println(q)
    fmt.Println(q.Dequeue())
}
*/
