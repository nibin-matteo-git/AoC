package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func GetInput(inputFile string)([]string) {
	contents, err := os.ReadFile(inputFile)
	HandleError(err)
	result := strings.Split(string(contents), "\n")
	out := make([]string, 0)
	for _, line := range(result){
		if strings.TrimSpace(line)==""{
			continue
		}
		out = strings.Split(strings.TrimSpace(line), " ")
	}

	return out
}
func HandleError(err error) {
	if err != nil {
		log.Fatalf("Encountered error: %s", err.Error())
	}
}

func addToMap(input []int) map[int]int {
	out := make(map[int]int, 0)
	for _, num := range input {
		out[num] = out[num] + 1
	}

	return out
}
func addToMapTwoLists[T comparable, V any](input1 []T, input2 []V) map[T][]V {
	out := make(map[T][]V, 0)
	for i, num := range input1 {
		out[num] = append(out[num], input2[i])
	}

	return out
}

func MapFuncToList[T, V any](input []T, function func(T) V) []V {
	out := make([]V, len(input))
	for i, val := range input {
		out[i] = function(val)
	}
	return out
}

func checkWithinbounds[T any](input [][]T, currLoc []int) bool{
	return currLoc[0]<len(input) && currLoc[0]>=0 && currLoc[1]<len(input[0]) && currLoc[1]>=0
}

// this wouldn't return error so can use directly on functions 
func convertStringToInt(input string)int{
	num, err := strconv.Atoi(input)
	if err!=nil{
		fmt.Println("error converting string: ", input)
	}
	return num
}


func combineNestedLoops[T any](input [][]T, nestedArrSize int)[]T{
	var out []T
	if nestedArrSize != -1{
		out = make([]T, 0, nestedArrSize)
	}else{
		out = make([]T, 0, len(input))
	}
	for _, val := range(input){
		out = append(out, val...)
	}
	return out
}
