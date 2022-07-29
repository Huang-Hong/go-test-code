package main

import "fmt"




func myDefer()  {
	fmt.Printf(" \n*************defer*****************\n")
	defer func ()  {
		fmt.Printf("defer function is called\n")
	}()

	fmt.Printf("test defer \n")
}	