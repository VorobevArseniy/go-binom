package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputs, err := parsedInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inputs)

	str, err := strNewtonBinom(inputs[0], inputs[1], inputs[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your binom: ", str[:len(str)-2])

}

func newtonBinom(a float64, b float64, n float64) float64 {
	var x, k float64

	for ; k <= n; k++ {
		x += combination(n, k) * math.Pow(a, n-k) * math.Pow(b, k)
	}

	return x
}

func strNewtonBinom(letter1 string, letter2 string, power string) (string, error) {
	var s string

	n, err := strconv.Atoi(power)
	if err != nil {
		return "", err
	}

	for k := 0; k <= n; k++ {
		s += fmt.Sprintf("%s! / (%d! * (%s - %d)!) * %s^(%s - %d) * %s^%s *", power, k, power, k, power, letter1, k, letter2, power)
	}

	return s, nil
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

func strInput() ([]string, error) {

	var num1, num2, power string
	res := make([]string, 0, 3)

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

func parsedInput() ([]string, error) {
	var s string
	res := make([]string, 0, 3)

	fmt.Println("input your binom (ex. (a+b)^n): ")

	_, err := fmt.Scanln(&s)
	if err != nil {
		return nil, err
	}

	s = strings.ReplaceAll(s, " ", "")

	re := regexp.MustCompile(`^\(([a-zA-Z])\+([a-zA-Z])\)\^(\d+)$`)
	matches := re.FindStringSubmatch(s)

	if matches == nil {
		return nil, fmt.Errorf("wrong!")
	}

	res = append(res, matches[1], matches[2], matches[3])

	return res, nil
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
