package main

import (
	"fmt"
)	
func main() {
        a := []int{}
		fmt.Println(a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("capacity: %v\n", cap(a))
		a = append(a, 1)
		fmt.Println(a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("capacity: %v\n", cap (a))
		fmt.Println(a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("capacity: %v\n", cap(a))
}