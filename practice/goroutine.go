package main

import (
	"fmt"
	"time"
)
func goRoutine1()  {
	for  i := 0; i < 5; i++ {
		fmt.Printf("%d \n", i)
	}
}

func testGoRoutine()  {
	fmt.Printf(" \n*************协程*****************\n")
	go goRoutine1();
	fmt.Printf(" end \n")
	time.Sleep(time.Second)
}