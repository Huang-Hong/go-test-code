package main

import "fmt"


func arrayExample()  {
	var myAarray [10] int
	myAarray[5] = 5
	fmt.Printf("myAarray[5]=%d  \n", myAarray[5])
}

func sliceExample()  {
	var mySlice = [] int {1,2,3,4,5}
	// mySlice = 
	fmt.Printf("mySlice[2]=%d  \n", mySlice[2])
}

func stringArray()  {
	var myString = [5]string{"I","am", "stupid", "and", "weak"}
	for key,value := range myString {
		fmt.Printf(" stringArray %d %s \n", key, value)
	}
	for i := 0; i < len(myString); i++ {
		
	}
}


