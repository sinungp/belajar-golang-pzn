package main

import (
	"fmt"
)

func main() {
	type noktp string

	var ktpsinung noktp = "11111"
	fmt.Println(ktpsinung)
	fmt.Println(noktp("2222"))
}
