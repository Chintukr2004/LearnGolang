package main

import (
	"fmt"
	"strconv"
)

func main() {

	num, err := strconv.Atoi("10")
	if err!=nil{
		fmt.Println("error", err)
	
	}

	if (num == 0 || num > 0) && (num % 2 == 0 ) { 
		fmt.Println("value is positive and also divisible")
	}else if (num == 0 || num > 0) && (num % 2 != 0 ){
		fmt.Println("value is positive but not div by 2")
		
	} else if num < 0 {
		
		fmt.Println("value is neg")
		}else {
			fmt.Println("value is  positive")
			
	}

}