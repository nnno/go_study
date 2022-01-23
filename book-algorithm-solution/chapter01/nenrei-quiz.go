package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var min, max int = 20, 36
	ans := initAnswer(min, max)
	exec(min, max, ans)
}

func initAnswer(min int, max int) int {
	// minとmaxを小さい順に返す
	f := func(a int, b int) (int, int) {
		if a >= b {
			return b, a
		} else {
			return a, b
		}
	}
	min, max = f(min, max)

	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func exec(min, max, ans int) {
	var end bool = false
	for end == false {
		n := (max + min) / 2
		// 正解なら終了
		if n == ans {
			fmt.Printf("%d 歳ですね。\n", n)
			fmt.Println("正解です")
			end = true
			break
		}

		fmt.Printf("%d 歳未満ですか\n", n)
		// High or Lowで次の値を求める
		if n > ans {
			fmt.Println("Yes")
			max = n
		} else {
			fmt.Println("No")
			min = n
		}
		time.Sleep(time.Second)
	}
}
