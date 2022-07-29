package main

import "fmt"

//结构体
func myStruct()  {
	fmt.Printf(" \n*************结构体*****************\n")

	type student struct {
		name string
		score int
	}
	var student1 = student{name: "mike", score: 5}
	fmt.Printf(" name = %s, score = %d \n", student1.name, student1.score)

	//todo: 结构体标签
}


