package main

import "fmt"

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	v := &ListNode{Next: head}
	s0 := v
	s1 := v

	for i := 0; i < n; i++ {
		if s1.Next != nil {
			s1 = s1.Next
		} else {
			return nil
		}
	}

	for s1.Next != nil {
		s0 = s0.Next
		s1 = s1.Next
	}

	s0.Next = s0.Next.Next

	return v.Next

}

func main() {
	a := []int{1, 2, 3, 4, 5}
	l := A2l(a)
	fmt.Println(L2a(l))

	fmt.Println(L2a(removeNthFromEnd(l, 5)))

}
