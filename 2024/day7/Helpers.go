package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type myStruct struct {
	target int
	values []int
}

func GetInput(inputFile string) []myStruct {
	contents, err := os.ReadFile(inputFile)
	HandleError(err)
	result := strings.Split(string(contents), "\n")
	out := make([]myStruct, 0)
	for _, line := range result {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		target, _ := strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[0]))
		values := make([]int, 0)
		for _, num := range strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " ") {
			intNum, _ := strconv.Atoi(num)
			values = append(values, intNum)
		}
		out = append(out, myStruct{target, values})
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
