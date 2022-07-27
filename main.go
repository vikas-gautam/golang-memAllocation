package main

import "fmt"

func main() { 
	//escapes to heap or doesn't escapes to heat

	x1 := 42        //not escapes to heap
	x2 := x1        // escapes to heap (and escapes itself)
	fmt.Println(x2) //x = 42
	y := make(map[int]string)
	fmt.Println(y) // y = map[]
	z := new(int)
	fmt.Println(z) // z = memory address
}

//causes variable to escape to the heap
