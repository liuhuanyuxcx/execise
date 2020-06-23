package main

import (
	"fmt"
	"github.com/liuhuanyuxcx/execise/exchange/account"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)
	account1 := account.NewAccount("1", "zhangsan")
	account1.SaveMoney(70000)
	var wg sync.WaitGroup

	for i := 1; i < 14; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := account1.DrawMoney(5000,i)
			if err != nil {
				fmt.Println("DrawMoney occur a error. It is ", err, ".")
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Account1's balance is %0.2f元\n", account1.QueryBalance())
}

