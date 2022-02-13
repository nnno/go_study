package solution

import (
	"reflect"
	"testing"
)

func Test_PostorderTraversal(t *testing.T) {
	type testData struct {
		name string
		root *TreeNode
		want []int
	}
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
		want: []int{3, 2, 1},
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := postorderTraversal(tt.root)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("postorderTraversal want(%v) : actual(%v)", tt.want, got)
			}
		})
	}
}
