package main

import (
	"fmt"
	"time"
)

func setValue(ch chan int)  {
	ch <- 1
}

func testChnnel()  {
	fmt.Printf(" \n*************管道*****************\n")

	var ch chan int
	ch = make(chan int)
	//ch := (chan int)

	go setValue(ch)
	var value int
	value = <-ch
	fmt.Printf("%d",value)

	close(ch)

	time.Sleep(time.Second)

}