package main

import "fmt"

//结构体类型+方法 就相当于 c++ 当中的类和类方法
type student struct {
	name string
	score int
}

//方法就是作用在接收者上的函数，接收者是一种类型，最常用的是结构体类型,当然也可以是其他类型
func (stu *student)getName()(name string)  {
	return stu.name
}

//1.接收者一般是传指针，假如传值的话，会有实例的拷贝，增加开销
//2.在这个函数这里，假如接收者这里传的是一个接收者的值，则是改变不了name的值的
func (stu *student)changeName(name string) {
	stu.name = name
}

func myMethod()  {
	fmt.Printf(" \n*************方法*****************\n")
	var student1 = student {name: "mike", score: 5}
	name := student1.getName()
	fmt.Printf("student1 name is %s \n", name)
	student1.changeName("paul")
	fmt.Printf("student1 change name is %s \n", student1.name)
}