package main

import (
	"fmt"
)

func main() {
	var name [3]string

	name[0] = "sinung"
	name[1] = "pra"
	name[2] = "yetno"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	//membuat array secara langsung

	var value = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(value)
	fmt.Println(value[0])
	fmt.Println(len(value))
	value[0] = 100
	fmt.Println(value)
}
