/*将指定go源码文件中的注释去掉*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	ENTER    = 10 //回车
	DIVIDE   = 47 //斜线
	ASTERISK = 42 //星号
)

func main() {
	filePath := "/Users/yh80816072/gopath/src/github.com/liuhuanyuxfq/execise/test1/main.go"
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	/* 1 表示/开头
	   2 表示//开头
	   3 表示/*开头
	   4 表示/*开头且*结尾
	*/
	flag := 0
	var target []byte

	for _, b := range buf {
		switch flag {
		case 0:
			if b == DIVIDE {
				flag = 1
			} else {
				target = append(target, b)
			}
			break
		case 1:
			if b == DIVIDE {
				flag = 2
			} else if b == ASTERISK {
				flag = 3
			} else {
				flag = 0
				target = append(target, DIVIDE)
				target = append(target, b)
			}
			break
		case 2:
			if b == ENTER {
				flag = 0
				target = append(target, ENTER)
			}
			break
		case 3:
			if b == ASTERISK {
				flag = 4
			}
			break
		case 4:
			if b == DIVIDE {
				flag = 0
				target = append(target, ENTER)
			} else if b == ASTERISK {
				flag = 4
			} else {
				flag = 3
			}
			break
		}
	}
	fmt.Print(string(target))
}
