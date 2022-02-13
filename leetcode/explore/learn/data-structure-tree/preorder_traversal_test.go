package solution

import (
	"reflect"
	"testing"
)

type testData struct {
	name string
	root *TreeNode
	want []int
}

func makeTestData() []*testData {
	var tests []*testData

	//-- test 1
	i1 := func() *TreeNode {
		root := &TreeNode{1, nil, nil}
		root.Right = &TreeNode{2, nil, nil}
		root.Right.Left = &TreeNode{3, nil, nil}
		return root
	}
	var t1 = &testData{
		name: "[1,null,2,3]",
		root: i1(),
		want: []int{1, 2, 3},
	}
	tests = append(tests, t1)

	//-- test 2
	var r2 *TreeNode = nil
	var w2 []int
	var t2 = &testData{
		name: "[]",
		root: r2,
		want: w2,
	}
	tests = append(tests, t2)

	//-- test 3
	var r3 *TreeNode = &TreeNode{1, nil, nil}
	var t3 = &testData{
		name: "[1]",
		root: r3,
		want: []int{1},
	}
	tests = append(tests, t3)

	return tests
}

func Test_PreorderTraversal(t *testing.T) {
	tests := makeTestData()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := preorderTraversal(tt.root)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("preorderTraversal want(%v) : actual(%v)", tt.want, got)
			}
		})
	}
}

func Test_PreorderTraversalStack(t *testing.T) {
	tests := makeTestData()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := preorderTraversalStack(tt.root)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("preorderTraversalStack want(%v) : actual(%v)", tt.want, got)
			}
		})
	}
}
