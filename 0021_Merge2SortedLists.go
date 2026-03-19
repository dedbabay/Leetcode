package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l2
	}

	l3 := &ListNode{}
	cur3 := l3
	cur2 := l2
	cur1 := l1

	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			cur3.Next = cur1
			cur1 = cur1.Next
		} else {
			cur3.Next = cur2
			cur2 = cur2.Next
		}
		cur3 = cur3.Next
	}

	if cur1 != nil {
		cur3.Next = cur1
	}
	if cur2 != nil {
		cur3.Next = cur2
	}

	return l3.Next

}

func makeList(a []int) *ListNode {
	r := &ListNode{Val: a[0]}

	cur := r
	for i := 1; i < len(a); i++ {
		cur.Next = &ListNode{Val: a[i]}
		cur = cur.Next
	}

	return r
}

func l2a(l *ListNode) []int {
	a := []int{}

	cur := l
	for {
		a = append(a, cur.Val)
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}

	return a
}

func main() {
	l1 := makeList([]int{1, 2, 4})
	l2 := makeList([]int{1, 3, 4})
	fmt.Println(l2a(l1))
	fmt.Println(l2a(l2))
	fmt.Println(l2a(mergeTwoLists(l1, l2)))
	fmt.Println(l2a(l1))
	fmt.Println(l2a(l2))
}
