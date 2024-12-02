package main 

import (
	"log"
	"os"
	"strings"
)

func GetInput(inputFile string)([]string){
	contents, err := os.ReadFile(inputFile)
	HandleError(err)
	result := strings.Split(string(contents), "\n")
	out := make([]string, 0)
	for _, num := range(result){
		if strings.TrimSpace(num) == ""{
			continue
		}
		out = append(out, strings.TrimSpace(num))
	}
	return out
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

func MapFuncToList[T, V any](input []T, function func(T)V)[]V{
	out:=make([]V, len(input))
	for i, val := range(input){
		out[i]=function(val)
	}
	return out
}
