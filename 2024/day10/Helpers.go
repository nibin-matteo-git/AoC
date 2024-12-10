package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)


func GetInput(inputFile string)([][]string, [][]int) {
	contents, err := os.ReadFile(inputFile)
	HandleError(err)
	result := strings.Split(string(contents), "\n")
	trailheadLocs := make([][]int, 0)
	re := regexp.MustCompile("0")
	for i, line := range(result){
		if strings.TrimSpace(line)==""{
			continue
		}
		indices := re.FindAllStringIndex(line, -1)
		for _, loc := range(indices){
			trailheadLocs = append(trailheadLocs, []int{i, loc[0]})
		}
	}


	out := make([][]string, 0)	
	for _, line := range(result){
		if strings.TrimSpace(line) == ""{
			continue
		}
		tmp := strings.TrimSpace(line)
		out = append(out, strings.Split(tmp, ""))
	}
	return out, trailheadLocs 
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
