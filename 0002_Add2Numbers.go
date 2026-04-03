package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	r := &ListNode{}
	p := 0

	c1 := l1
	c2 := l2
	c0 := r

	for c1 != nil || c2 != nil || p > 0 {

		v := p
		if c1 != nil {
			v += c1.Val
			c1 = c1.Next
		}
		if c2 != nil {
			v += c2.Val
			c2 = c2.Next
		}

		p = v / 10
		v = v % 10

		c0.Next = &ListNode{Val: v}
		c0 = c0.Next
	}

	return r.Next
}

func main() {
	a1 := []int{2, 4, 3}
	a2 := []int{5, 6, 4}

	fmt.Println(L2a(addTwoNumbers(A2l(a1), A2l(a2))))

}
