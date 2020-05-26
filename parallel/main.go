/*
  题目：给定一个数组（数组包含很多元素），请用5个协程并发读取并打印数据元素。
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	list := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, //11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	}
	ch := make(chan int, 0)
	flag := make(chan int, 0)
	fmt.Println("数组内容：", list)
	fmt.Println("输出结果：")
	go func() {
		for i := 0; i < len(list); i++ {
			ch <- list[i]
		}
		flag <- -1
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
		Loop:
			for {
				select {
				case e, ok := <-ch:
					if ok {
						fmt.Printf("第%d个协程，数组元素是：%d\n", i, e)
					} else {
						break Loop
					}
				case <-flag:
					close(ch)
					break Loop
				}
			}
		}(i)
	}
	wg.Wait()
}
