package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
	gpa  float32
	addr Addr
}

type Addr struct{
	vill string
	city string
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
	s1.addr.city = "hazaribagh"
	s1.addr.vill = "padma"
	fmt.Println(s1)
	
	s2 :=Student{name: "shivam", age: 22, gpa: 6.6, addr: Addr{city: "ranchi", vill: "kuma"}}
	
	fmt.Println(s2)

	
	val := 122
	val2 := "1231"

	var interfaceExample interface{}

	interfaceExample = val
	fmt.Println("Interface value:",interfaceExample)
	interfaceExample = val2
	fmt.Println("Interface value:",interfaceExample)



	var vishladress *Addr
	vishladress = &s1.addr

	shivamAddress := &s2.addr
	
	fmt.Println("vishal Adress: ", vishladress , "shivamAddress" , shivamAddress)

}


