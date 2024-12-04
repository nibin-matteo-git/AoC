package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	input := GetInput(inputPath)
	result := 0
	// why doesn't * work instead of d
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	doR := regexp.MustCompile(`do\(\)`)
	dontR := regexp.MustCompile(`don't\(\)`)


	matches := r.FindAllString(input, -1)
	matchIndices := r.FindAllStringIndex(input, -1)
	// fmt.Println(matchIndices)
	doIndices := doR.FindAllStringIndex(input, -1)
	dontIndices := dontR.FindAllStringIndex(input, -1)
	fmt.Println(doIndices)
	fmt.Println(dontIndices)

	for i, vals := range matches {
		// fmt.Println(shouldInclude(matchIndices[i][0], doIndices, dontIndices))
		if shouldInclude(matchIndices[i][0], doIndices, dontIndices) {

			result += getSum(vals)
		}
		// fmt.Println(matchIndices)
		// fmt.Println(doIndices)
		// fmt.Println(dontIndices)
		// fmt.Println(i)

	}


	fmt.Println(result)

}
func shouldInclude(matchIndex int, doIndices [][]int, dontIndices [][]int) bool {
	if len(dontIndices) == 0 {
		return true
	}
	if matchIndex < dontIndices[0][0] {
		return true
	}

	// for _, dontIndex := range(dontIndices){
	// 	if matchIndex>dontIndex{
	// 		for _,doIndex := range(doIndices){
	// 			if matchIndex < doIndex{
	// 				result= false
	// 			}if doIndex>dontIndex &&
	// 		}
	// 	}
	// }
	closestDont := 0
	for _, dontIndex := range dontIndices {
		if dontIndex[0] < matchIndex {
			closestDont = dontIndex[0]
		}
	}
	closestDo := 0
	for _, doIndex := range doIndices {
		if doIndex[0] < matchIndex {
			closestDo = doIndex[0]
		}
	}
	// fmt.Println("closest inices", closestDo, closestDont)
	return closestDont <= closestDo
}
func getSum(val string) int {
	expr := val[4 : len(val)-1]
	// fmt.Println(strings.Split(expr, ","))
	num1, _ := strconv.Atoi(strings.Split(expr, ",")[0])
	num2, _ := strconv.Atoi(strings.Split(expr, ",")[1])
	// fmt.Println(num1,num2)
	return num1 * num2
}
