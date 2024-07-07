package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(addNumber(3, 4))
	var arg [3]int = [3]int{1, 2, 3}
	fmt.Println(arg)
	slice := []int{1, 22, 3}
	fmt.Println(slice)
	var mapname map[int]int = map[int]int{
		3: 2,
		4: 2,
		2: 1,
	}
	fmt.Println(mapname)

	yama := Human{false, 2}
	fmt.Println(yama)

	fmt.Println(yama.run(4, 5))
}

func addNumber(a int, b int) int {
	return a + b
}

type Human struct {
	Head bool
	Arm  int
}

func (h Human) run(power int, height int) int {
	return power * height * h.Arm
}
