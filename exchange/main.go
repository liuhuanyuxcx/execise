package main

import (
	"fmt"
	"github.com/liuhuanyuxfq/execise/exchange/account"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)
	account1 := account.NewAccount("1", "zhangsan")
	account1.SaveMoney(70000)
	var wg sync.WaitGroup

	for i := 0; i < 14; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//fmt.Println(fmt.Sprintf("第%d个人取款%d元。", i, 5000))
			err := account1.DrawMoney(5000,i)
			if err != nil {
				fmt.Println("DrawMoney occur a error. It is ", err, ".")
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Account1's balance is ", account1.QueryBalance(), ".")
}

