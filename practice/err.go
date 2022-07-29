package main

import (
	"fmt"
	"errors"
)

func errHandler()  {
	fmt.Printf(" \n*************错误处理*****************\n")
	var error1 = errors.New("this is an error")
	fmt.Printf("error: %v \n", error1)
}