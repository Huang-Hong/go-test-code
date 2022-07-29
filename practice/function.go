package main

import "fmt"

//多返回值函数
func multiReturnValueFunc(a int, b int) (c int, d int){
	var A = a + 1
	var B = b + 1
	return A, B 
}

func callBackFunc(a int)  {
	fmt.Printf("this is call back function, a=%d\n", a)
}

func testCallBackFunc(b func(a int))  {
	fmt.Printf("call the callBack function\n")
	b(12)
}

func myFunction()  {
	fmt.Printf(" \n*************函数*****************\n")

	a, b := multiReturnValueFunc(1, 2)
	fmt.Printf("multi return value function:%d %d \n", a , b)

	//test callack function
	testCallBackFunc(callBackFunc)
}


//todo: 闭包