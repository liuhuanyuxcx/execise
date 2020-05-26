package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	bytes := []byte(s)
	max := 0
	i := 0
	switch len(bytes) {
	case 0:
		return 0
	case 1:
		return 1
	default:
		for i < len(bytes) {
			//fmt.Println("i===", i)
			m := make(map[byte]int)
			m[bytes[i]] = 1
			j := i + 1
			for ; j < len(bytes); j++ {
				//fmt.Println("j=", j)
				//fmt.Println(m)
				pSame, ok := m[bytes[j]]
				if ok {
					if max < j-i {
						max = j - i
					}
					i = i + pSame
					break
				} else {
					m[bytes[j]] = j - i + 1
				}
			}
			if j >= len(bytes) {
				if max < j-i {
					max = j - i
				}
				break
			}
		}
		return max
	}
}

func main() {
	str := "abcbcad"
	n := lengthOfLongestSubstring(str)
	fmt.Println(n)
	fmt.Println(18546*.12)
}
