package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	// 1st test case.
	execAddTwoNumbers(t, sliceToNode([]int{2, 4, 3}), sliceToNode([]int{5, 6, 4}), sliceToNode([]int{7, 0, 8}))
	// longer than other.
	execAddTwoNumbers(t, sliceToNode([]int{1, 8}), sliceToNode([]int{0}), sliceToNode([]int{1, 8}))
	// carry
	execAddTwoNumbers(t, sliceToNode([]int{9, 9}), sliceToNode([]int{1}), sliceToNode([]int{0, 0, 1}))
	// one list is empty.
	execAddTwoNumbers(t, sliceToNode([]int{}), sliceToNode([]int{0, 1}), sliceToNode([]int{0, 1}))
}

func execAddTwoNumbers(t *testing.T, l1 *ListNode, l2 *ListNode, correct *ListNode) {
	result := addTwoNumbers(l1, l2)
	r := nodeToSlice(result)
	c := nodeToSlice(correct)

	if reflect.DeepEqual(r, c) == false {
		t.Fatalf("Wrong Answer. result: %v, correct: %v\n", r, c)
	}
}

func sliceToNode(input []int) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	for _, i := range input {
		current.Next = &ListNode{Val: i, Next: nil}
		current = current.Next
	}

	return dummyHead.Next
}

func nodeToSlice(node *ListNode) []int {
	var output []int
	for node != nil {
		output = append(output, node.Val)
		node = node.Next
	}
	return output
}
