package main

import (
	"fmt"
	"time"
)

var phlegms = make(chan int) //这里存所有吐得痰

func spit(phlegm int) { //吐痰函数
	phlegms <- phlegm
}

func cleanItUp(index int) { //大妈函数
	for {
		select {
		case phlegm := <-phlegms:
			fmt.Printf("I am %v ,and I cleaned %v \n", index, phlegm)
		}
	}
}

func main() {
	for i := 0; i < 100; i++ {
		go spit(i) //连吐1000口痰
	}
	for i := 0; i < 10; i++ {
		go cleanItUp(i) //生成10个大妈
	}

	<-time.After(5 * time.Second)
}
