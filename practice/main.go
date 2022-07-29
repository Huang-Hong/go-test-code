package main

import "fmt"

func main()  {
	fmt.Printf("this is my first go function \n \n")
	arrayExample()
	sliceExample()
	myMap()
	stringArray()
	myStruct()
	myFunction()
}

func init()  {
	fmt.Println("this is init function of this package")
}