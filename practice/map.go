package main

import "fmt"

//map
func myMap()  {
	fmt.Printf(" \n*************map*****************\n")

	var myMap1 = map[string]int {"first": 1, "second": 2}
	myMap1["third"] = 3;
	for key, value := range myMap1 {
		fmt.Printf("key= %s, value = %d \n", key, value)
	} 
}