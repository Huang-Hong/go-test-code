package main

import "fmt"

//切片
func sliceExample()  {
	fmt.Printf(" \n*************切片*****************\n")

	//初始化
	var mySlice = [] int {1,2,3,4,5}
	fmt.Printf("mySlice[2]=%d  \n", mySlice[2])
	//追加元素
    mySlice = append(mySlice, 6)
	fmt.Printf("mySlice[5]=%d  \n", mySlice[5])

}