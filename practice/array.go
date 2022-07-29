package main

import "fmt"

//数组
func arrayExample()  {
	fmt.Printf("*************数组*****************\n")
	var myAarray [10] int
	myAarray[5] = 5
	fmt.Printf("myAarray[5]=%d  \n", myAarray[5])
}