package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	input := GetInput(inputPath)
	myMap := make(map[int]int)
	for _, num := range(input){
		myMap[num] = myMap[num]+1
	}
	for index1, num1 := range input {
		expVal1 := 2020 - num1
		for index2, num2 := range input{
			if index2 != index1 {
				expVal2 := expVal1 - num2
				_, ok := myMap[expVal2]
				if ok{
					fmt.Printf("We found a match: %d and %d and %d \nanswer: %d", num1, num2, expVal2, (num1 * num2 * expVal2))

				}
			}
		}
	}
}
