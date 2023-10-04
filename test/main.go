package main

import "fmt"

func main() {
	var r int
	r = adder(1, 3)
	fmt.Println(r)

}

func adder(a int, b int) int {
	return a + b
}
