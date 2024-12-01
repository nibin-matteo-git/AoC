package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	input1, input2 := GetInput(inputPath)
	fmt.Println(findSimilarity(input1, input2))

}
func findDistance(input1 []int, input2 []int)int{
	sort.Ints(input1)
	sort.Ints(input2)

	outDist := 0
	for i, _ := range(input1){
		outDist+=int(math.Abs(float64(input1[i]-input2[i])))
	}
	return outDist
}
func findSimilarity(input1 []int, input2 []int)int{
	input2Map := addToMap(input2)
	result := 0
	for _, num := range(input1){
		result += num*input2Map[num]
	}
	return result
}
