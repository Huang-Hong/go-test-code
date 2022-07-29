package main

import "fmt"

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