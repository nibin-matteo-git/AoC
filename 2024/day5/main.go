package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	printOrder, printRules := GetInput(inputPath)
	printOrderList1, printOrderList2 := make([]int,0), make([]int, 0)

	finalOut := 0


	MapFuncToList(printRules, func(input string)int{
		int1, _ := strconv.Atoi(strings.Split(input, "|")[0])
		int2, _ := strconv.Atoi(strings.Split(input, "|")[1])
		printOrderList1 = append(printOrderList1, int1)
		printOrderList2 = append(printOrderList2, int2)
		return 0
	})
	myRules := addToMapTwoLists(printOrderList1, printOrderList2)
	parsedPrintInput := MapFuncToList(printOrder, func(input string) []int{
		out := make([]int, 0)
		for _, num := range(strings.Split(input, ",")){
			intNum, _ := strconv.Atoi(num)
			out = append(out, intNum)
		}
		return out
	})
	for _, printInput := range(parsedPrintInput){
		// _, middleIndex := checkPrintOrder(myRules,printInput)
		// finalOut+=middleIndex
		correctOrder, _ := checkPrintOrder(myRules, printInput)
		if !correctOrder{
			finalOut += calcMiddleIndexVal(fixPrintingOrder(myRules, printInput))
		}
	}
	fmt.Println(finalOut)

}

func calcMiddleIndexVal(input []int)int{
	return input[int(math.Round(float64(len(input))/2))-1] 
}


func fixPrintingOrder(rules map[int][]int, input[]int)[]int{
	earliestPage:=-1
	alrPrintedPgs := make(map[int]int, 0)
	for i, pg := range(input){
		for _, rule := range(rules[pg]){
			earliestPageIndex, wrongOrder := alrPrintedPgs[rule]
			if wrongOrder{
				if earliestPage==-1 || earliestPage>earliestPageIndex{
					earliestPage=earliestPageIndex
				}
			}
		}
		if earliestPage!=-1{
			newInput := append(input[:i],input[i+1:]...)
			newInput = slices.Insert(newInput, earliestPage, pg)
			return fixPrintingOrder(rules, newInput)
		}
		alrPrintedPgs[pg]=i
	}
	return input
}

func checkPrintOrder(rules map[int][]int, input []int)(bool, int){
	alrPrintedPgs := make(map[int]int, 0)
	for _, pg := range(input){
		for _, rule := range(rules[pg]){
			_, wrongOrder := alrPrintedPgs[rule]
			if wrongOrder{
				return false, 0 
			}
		}
		alrPrintedPgs[pg]+=1
	}
	return true, input[int(math.Round(float64(len(input))/2))-1] 
}
