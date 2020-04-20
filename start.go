package main

import "fmt"

type trade struct {
	symbol string
	price  float64
}

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}
func arraysum(arr []int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum
}
func main() {
	// var x float64
	// var y float64
	// x = 1
	// y = 2
	// fmt.Printf("x=%v\n", x)
	// fmt.Printf("y=%v\n", y)
	// var mean float64
	// mean = (x + y) / 2

	// fmt.Printf("hey there! mean is %v\n", mean)

	// z := 10
	// if z == 10 {
	// 	fmt.Println("z is:", z)
	// }
	// poem := "hello this is saksham garg"
	// words := strings.Fields(poem)
	// fmt.Println(words[1])
	// a := 1
	// b := 2
	// val := add(a, b)
	// fmt.Println(val)
	// fmt.Println(divmod(a, b))
	// arr := []int{1, 2, 3, 4}
	// sum1 := arraysum(arr)
	// fmt.Println("sum is:", sum1)
	// t1 := trade{"MSFT", 40}
	fmt.Println("hello")
	data := make(map[string]int)

}
