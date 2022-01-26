package main

// hashに詰める作業と、計算作業を一つのfor文でやると少し高速に
func twoSum(nums []int, target int) []int {
	m := map[int]int{} // map["計算結果値": スライスのkey ...]
	for i, n1 := range nums {
		if v, ok := m[target-n1]; ok == true {
			return []int{v, i}
		}
		m[n1] = i
	}
	return []int{}
}

/**
// 予めtarget-nした値をkeyとしたmapに保持して後から探すパターン
func twoSum(nums []int, target int) []int {
	m := map[int]int{}	// map["計算結果値": スライスのkey ...]
	for i, n1 := range nums {
		m[n1] = i
	}
	for j, n2 := range nums {
		if v, ok := m[target - n2]; ok == true {
			return []int{j, v}
		}
	}
	return []int{}
}
*/

/**
// 単純なループで実現したもの. これでは計算量がべき乗に増えるため改善が必要
func twoSum(nums []int, target int) []int {
	for i, n1 := range nums {
		for j, n2 := range nums {
			if n1 + n2 == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
*/
