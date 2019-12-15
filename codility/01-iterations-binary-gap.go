package solution

import (
	"fmt"
	"strings"
)

func strToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}

func binaryGap(n int) int {
	var (
		result = 0
		s      = strToBin(string(n))
		list   = strings.Split(s, "1")
	)
	fmt.Println(s)
	list = list[0 : len(list)-1] // 最後の要素を削除 (1で囲われているとはいえないため)

	for _, v := range list {
		if len(v) > result {
			result = len(v)
		}
	}
	return result
}

/** 正規表現の先読み,後読み(look ahead, look bihind)が使えれば...
func binaryGap(n int) int {
	var result = 0

	r := regexp2.MustCompile(`(?<=1)0+(?=1)`, 0)
	if m, _ := r.FindStringMatch(strToBin(string(n))); m != nil {
		if len(m.String()) > result {
			result = len(m.String())
		}
		for m != nil {
			m, _ = r.FindNextMatch(m)
			if m != nil && len(m.String()) > result {
				result = len(m.String())
			}
		}
	}
	return result
}
*/
