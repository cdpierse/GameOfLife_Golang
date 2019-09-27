package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	make2Dslice()
}


func make2Dslice() {
	darry := make([][]int,10)
		
	for i :=0; i<10;i++{
		darry[i] = make([]int,4)
	}
	darry[1][1] =3
	
	for i:=0; i<10; i++{
		fmt.Println(darry[i])
	
	}
	
	
}	
