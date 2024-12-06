package main 

import (
	"log"
	"os"
	"strings"
)

func GetInput(inputFile string)([][]string, []int){
	contents, err := os.ReadFile(inputFile)
	HandleError(err)
	result := strings.Split(string(contents), "\n")
	out := make([][]string, 0)	
	guardLoc := []int{-1,-1}
	for i, num := range(result){
		if strings.TrimSpace(num) == ""{
			continue
		}
		tmp := strings.TrimSpace(num)
		if guardLoc[1]==-1{
			guardLoc[1] = strings.IndexRune(tmp, '^')
			guardLoc[0] = i
		}	
		out = append(out, strings.Split(tmp, ""))
	}
	return out, guardLoc
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
