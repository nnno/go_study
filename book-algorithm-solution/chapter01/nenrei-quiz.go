package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var min, max int = 20, 36

	slice := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		slice[i] = min + i
	}
	ans := slice

	for j := 0; j < len(ans); j++ {
		index, cost := binarySearchQuestion(ans[j], slice)
		fmt.Printf("年齢: %d 歳\tcost: %d\n", slice[index], cost)
		fmt.Println("----")
	}

	for j := 0; j < len(ans); j++ {
		index, cost := binarySearch(ans[j], slice)
		fmt.Printf("年齢: %d 歳\tcost: %d\n", slice[index], cost)
	}

}

// 初期化 答えの作成
func initAnswer(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	n := int(math.Abs(float64(max - min)))
	return min + rand.Intn(n)
}

func binarySearch(n int, slice []int) (index int, cost int) {
	min, max := 0, len(slice)
	count := 0
	for max-1 >= min {
		count++
		mid := (min + max) / 2
		// 質問を再現しないbinary searchなら答えを見つけた瞬間に終了
		if n == slice[mid] {
			return mid, count
		} else if n > slice[mid] {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return -1, count
}

func binarySearchQuestion(n int, slice []int) (index int, cost int) {
	min, max := 0, len(slice)
	count := 0

	for max-1 >= min {
		count++
		mid := (min + max) / 2
		if slice[mid] > n {
			fmt.Printf("%d 歳未満ですか ... %s\n", slice[mid], "Yes")
			max = mid
			if max-min < 1 {
				if n == slice[mid] {
					return mid, count
				} else if n == slice[mid-1] {
					return mid - 1, count
				}
			}
		} else {
			fmt.Printf("%d 歳未満ですか ... %s\n", slice[mid], "No")
			min = mid + 1
			if max-min < 1 {
				if n == slice[mid] {
					return mid, count
				} else if n == slice[mid+1] {
					return mid + 1, count
				}
			}
		}
	}
	return -1, count
}
