package lifofifo
import (
)
type Fifo []*string

func (q *Fifo) Push(n *string) {
    *q = append(*q, n)
}
func (q *Fifo) PushOut(n *string, length int) {
    *q = append(*q, n)
	if len(*q) > length {
		n = (*q)[0]
		*q = (*q)[1:length]
	}
}

func (q *Fifo) Pop() (n *string) {
    n = (*q)[0]
    *q = (*q)[1:]
    return
}

func (q *Fifo) Len() int {
    return len(*q)
}

type Lifo []*string

func (q *Lifo) Push(n *string) {
    *q = append(*q, n)
}

func (q *Lifo) Pop() (n *string) {
    x := q.Len() - 1
    n = (*q)[x]
    *q = (*q)[:x]
    return
}
func (q *Lifo) Len() int {
    return len(*q)
}
