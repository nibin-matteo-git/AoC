package main 

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput(inputFile string)([]int, []int){
	// file, err := os.OpenFile(inputFile)
	// HandleError(err)
	// defer file.Close()
	contents, err := os.ReadFile(inputFile)
	HandleError(err)
	result := strings.Split(string(contents), "\n")
	out1 := make([]int,0)
	out2 := make([]int,0)
	for _, num := range(result){
		if strings.TrimSpace(num) == ""{
			continue
		}
		num1, num2 := strings.Split(num, "  ")[0], strings.Split(num, "  ")[1]
		num1 = strings.TrimSpace(num1)
		num2 = strings.TrimSpace(num2)
		intNum1, err := strconv.Atoi(num1)
		HandleError(err)
		intNum2, err := strconv.Atoi(num2)
		HandleError(err)
		out1 = append(out1, intNum1)
		out2 = append(out2, intNum2)
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
