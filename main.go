package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	inputs, err := input()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inputs)

	fmt.Println("Your binom: ", newtonBinom(inputs[0], inputs[1], inputs[2]))

}

func newtonBinom(a float64, b float64, n float64) float64 {
	var x, k float64

	for ; k <= n; k++ {
		x += combination(n, k) * math.Pow(a, n-k) * math.Pow(b, k)
	}

	return x
}

func factorial(n float64) float64 {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func combination(n float64, k float64) float64 {
	c := factorial(n) / (factorial(k) * factorial(n-k))
	return c
}

func input() ([]float64, error) {

	var num1, num2, power float64
	res := make([]float64, 0, 3)

	fmt.Println("input first letter: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		return nil, err
	}

	fmt.Println("input second letter: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		return nil, err
	}

	fmt.Println("input power: ")
	_, err = fmt.Scanln(&power)
	if err != nil {
		return nil, err
	}

	res = append(res, num1, num2, power)

	return res, nil
}
