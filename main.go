package main

import (
	"fmt"
	"log"
)

func main() {
	inputs, err := input()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inputs)
}

// func makeBinom()

func input() ([]string, error) {

	var letter1, letter2, power string
	res := make([]string, 0, 3)

	fmt.Println("input first letter: ")
	_, err := fmt.Scanln(&letter1)
	if err != nil {
		return nil, err
	}

	fmt.Println("input second letter: ")
	_, err = fmt.Scanln(&letter2)
	if err != nil {
		return nil, err
	}

	fmt.Println("input power: ")
	_, err = fmt.Scanln(&power)
	if err != nil {
		return nil, err
	}

	res = append(res, letter1, letter2, power)

	return res, nil
}
