/*
  题目：n个小孩子都有自己都编号1，2...n,让这些小孩按顺序围成一个圈，
       然后从1号开始报数，报到5的小孩出圈，记录出圈小孩的编号，然后
       由出圈小孩下一个位置的小孩开始重新报数，报到5的小孩仍出圈，并记录
       编号，以此类推，直到圈内无小孩为止。最后请依次输出出圈小孩的编号。
*/
package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type Child struct {
	Id   int
	Next *Child
}

func NewCircle(n int) (*Child, error) {
	if n < 1 {
		return nil, errors.New("小孩总数需大于0")
	}
	tail := &Child{Id: n}
	head := tail
	for i := n - 1; i > 0; i-- {
		child := Child{i, head}
		head = &child
	}
	tail.Next = head
	return head, nil
}

func (circle *Child) String() string {
	s := fmt.Sprintf("%d ", circle.Id)
	for next := circle.Next; next != circle; {
		s = s + fmt.Sprintf("%d ", next.Id)
		next = next.Next
	}
	return s
}

func Report(circle *Child) []int {
	res := make([]int, 0)
	for circle.Next != circle {
		cur := circle
		for i := 2; i < 5; i++ {
			cur = cur.Next
		}
		res = append(res, cur.Next.Id)
		cur.Next = cur.Next.Next
		circle = cur.Next
	}
	res = append(res, circle.Id)
	return res
}

func main() {
	circle, _ := NewCircle(13)
	fmt.Println(circle)
	res := Report(circle)
	fmt.Println("res:", res)
}
