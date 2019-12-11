package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}
func main() {
	num1, num2 := 5.6, 9.5
	fmt.Println(add(num1, num2))
	fmt.Println(mutiple("c", "d"))

}
func mutiple(a, b string) (string, string) {
	return a, b
}
// this is a new message