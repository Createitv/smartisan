// @Time        : 2023/3/9 21:39
// @Author      : Createitv
// @FileName    : main.go
// @Software    : GoLand
// @WeChat      : Navalism1
// @Description : 

package main

import (
	"fmt"
	
	"github.com/createitv/smartisan/random"
)

func main() {
	for i := 0; i < 10; i++ {
		
		fmt.Println(random.UUIdV4())
	}
	
}
