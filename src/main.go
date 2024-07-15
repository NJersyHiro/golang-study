package main

import "fmt"

func main() {
	var input1 int
	fmt.Scan(&input1)
	fmt.Printf("kaijo is %d", kaijo(input1))

}

func kaijo(b int) int {
	var sum int
	sum = b
	for b > 1 {
		sum = sum * (sum - b)
		b -= 1
	}
	return sum
}
