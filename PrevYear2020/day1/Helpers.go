package main 

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput(inputFile string)[]int{
	// file, err := os.OpenFile(inputFile)
	// HandleError(err)
	// defer file.Close()
	contents, err := os.ReadFile(inputFile)
	HandleError(err)
	result := strings.Split(string(contents), "\n")
	out := make([]int,0)
	for _, num := range(result){
		num := strings.TrimSpace(num)
		if num == ""{
			continue
		}
		intNum, err := strconv.Atoi(num)
		HandleError(err)
		out = append(out, intNum)
	}
	return out
}
func HandleError(err error){
	if err!=nil{
		log.Fatalf("Encountered error: %s", err.Error())
	}
}
