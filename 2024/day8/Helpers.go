package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)


func GetInput(inputFile string) (map[string][][]int, []string) {
	contents, err := os.ReadFile(inputFile)
	HandleError(err)
	result := strings.Split(string(contents), "\n")
	rawResult := make([]string, 0)
	re := regexp.MustCompile(`[a-zA-Z0-9]`)

	var out = make(map[string]([][]int))
	for i, line := range(result){
		if strings.TrimSpace(line)==""{
			continue
		}
		rawResult = append(rawResult, strings.TrimSpace(line))
		indices := re.FindAllStringIndex(line, -1)
		for _, index := range(indices){
			_, present := out[string(line[index[0]])]
			if !present{
				out[string(line[index[0]])]=make([][]int, 0)
			}
			out[string(line[index[0]])]= append(out[string(line[index[0]])], []int{i, index[0]})

		}
	}	
	return out, rawResult 
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
