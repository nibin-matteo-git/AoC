package main 

import (
	"log"
	"os"
	"strings"
)

func GetInput(inputFile string)([]string, []string){
	contents, err := os.ReadFile(inputFile)
	HandleError(err)
	result := strings.Split(string(contents), "\n\n")
	printOrder := strings.Split(result[1], "\n")
	printRules := strings.Split(result[0], "\n")
	out1 := make([]string, 0)
	out2 := make([]string, 0)
	for _, num := range(printOrder){
		if strings.TrimSpace(num) == ""{
			continue
		}
		out1 = append(out1, strings.TrimSpace(num))
	}
	for _, num := range(printRules){
		if strings.TrimSpace(num) == ""{
			continue
		}
		out2 = append(out2, strings.TrimSpace(num))
	}
	return out1, out2
}
func HandleError(err error){
	if err!=nil{
		log.Fatalf("Encountered error: %s", err.Error())
	}
}

func addToMap(input []int) map[int]int{
	out := make(map[int]int, 0)
	for _, num := range(input){
		out[num] = out[num]+1
	}

	return out
}
func addToMapTwoLists[T comparable, V any](input1 []T, input2 []V) map[T][]V{
	out := make(map[T][]V, 0)
	for i, num := range(input1){
		out[num]=append(out[num], input2[i])
	}

	return out
}

func MapFuncToList[T, V any](input []T, function func(T)V)[]V{
	out:=make([]V, len(input))
	for i, val := range(input){
		out[i]=function(val)
	}
	return out
}
