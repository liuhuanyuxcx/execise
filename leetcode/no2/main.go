package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	var cur *ListNode
	var res *ListNode
	x := 0
	for {
		switch {
		case p1 != nil && p2 != nil:
			val, tag := sum(p1.Val, p2.Val, x)
			cur = &ListNode{val, res}
			res = cur
			x = tag
			p1 = p1.Next
			p2 = p2.Next
		case p1 == nil && p2 != nil:
			val, tag := sum(0, p2.Val, x)
			cur = &ListNode{val, res}
			res = cur
			x = tag
			p2 = p2.Next
		case p1 != nil && p2 == nil:
			val, tag := sum(p1.Val, 0, x)
			cur = &ListNode{val, res}
			res = cur
			x = tag
			p1 = p1.Next
		case p1 == nil && p2 == nil && x == 1:
			val, tag := sum(0, 0, x)
			cur = &ListNode{val, res}
			res = cur
			x = tag
		default:
			return res.reverse()
		}
	}
	//return res
}

func NewList(values ...int) *ListNode {
	node := &ListNode{}
	var list *ListNode
	for _, value := range values {
		node = &ListNode{value, list}
		list = node
	}
	return list
}

func (list *ListNode) String() string {
	var res string
	p := list
	for p != nil {
		res = res + fmt.Sprintf("%d", p.Val)
		p = p.Next
	}
	return res
}

func (list *ListNode) reverse() *ListNode {
	cur := list
	var res *ListNode
	for cur != nil {
		cur, res, cur.Next = cur.Next, cur, res
	}
	return res
}

func sum(v1, v2, tag int) (int, int) {
	if v1+v2+tag >= 10 {
		return v1 + v2 + tag - 10, 1
	} else {
		return v1 + v2 + tag, 0
	}
}

func main() {
	l1 := NewList(5)
	fmt.Println(l1)
	l2 := NewList(5)
	fmt.Println(l2)
	res := addTwoNumbers(l1, l2)
	fmt.Println(res)
}
