package main

import "fmt"

//数组
func arrayExample()  {
	fmt.Printf("*************数组*****************\n")
	var myAarray [10] int
	myAarray[5] = 5
	fmt.Printf("myAarray[5]=%d  \n", myAarray[5])
}

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

//循环
func stringArray()  {
	fmt.Printf(" \n*************循环*****************\n")

	var myString = [5]string{"I","am", "stupid", "and", "weak"}
	// for...range
	for key,value := range myString {
		fmt.Printf(" stringArray %d %s \n", key, value)
	}
	//for ; ;
	for i := 0; i < len(myString); i++ {
		if myString[i] == "stupid" {
			myString[i] = "strong"
		}
		if myString[i] == "weak" {
			myString[i] = "clear"
		}
		fmt.Printf("%s ", myString[i] )
	}
	fmt.Printf("\n")

}


