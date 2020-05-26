//求MAX以内的素数，并打印之
package main

import (
	"fmt"
	"math"
)

const (
	MAX    = 1100
	NUMBER = 20 //每行打印多少个数字
)

func main() {
	n := MAX
	target := []int{1, 2}
Loop:
	for i := 3; i <= n; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue Loop
			}
		}
		target = append(target, i)
	}
	for i, t := range target {
		fmt.Print(format(t, n))
		if (i+1)%NUMBER == 0 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
}

func format(in int, n int) string {
	f := int(math.Ceil(math.Log10(float64(n))))
	patten := fmt.Sprintf("%s%d%s", "%", f, "d")
	out := fmt.Sprintf(patten, in)
	return out
}
