package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
	gpa  float32
}

type Employee struct{
	name string
	age int
	salary float32
}

type Phone struct{
	BasicInfo BasicInfo
	IMEI int
	PRICE int
	Configuration Config
}

type Laptop struct{
	BasicInfo BasicInfo
	SerialNo string
	Configuration Config
}
type BasicInfo struct{
	Brand string
	Model string
}
type Config struct {
	RAM int
	ROM int
	Processor string
	OS string
}

func main() {
	fmt.Println("learning datatype")
	var s1 Student

	s1.name = "John"
	s1.age = 21
	s1.gpa = 3.5
	fmt.Println(s1)

	eam := Employee{"Arv", 22, 50000.50}

	fmt.Println(eam)

	val := 122
	val2 := "1231"

	var interfaceExample interface{}

	interfaceExample = val
	fmt.Println("Interface value:",interfaceExample)
	interfaceExample = val2
	fmt.Println("Interface value:",interfaceExample)

}


