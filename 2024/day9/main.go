package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)



func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	input := GetInput(inputPath)
	expandedInput := generateExpandedView(input)
	// fmt.Println(expandedInput)
	// fmt.Println(p1(expandedInput))
	fmt.Println(caclCheckSum(p1(expandedInput)))
}

func p1(input []string)[]string{
	emptLocs := make([]int, 0)
	for i, val := range(input){
		if val=="."{
			emptLocs = append(emptLocs, i)
		}
	}
	counter := 0
	for i:=len(input)-1; i>=0 && counter<len(emptLocs) && emptLocs[counter]<i; i--{
		// fmt.Println("the new input array: ", input)
		// fmt.Println(emptLocs[counter])
		if input[i]!="."{
			input[emptLocs[counter]] = input[i]
			input[i]="."	
			counter++
		}
	}
	return input
}

func generateExpandedView(input string)[]string{
	out := make([]string, 0)
	for i, num := range(input){
		intNum, _ := strconv.Atoi(string(num))
		for j:=0; j<intNum; j++{
			if i%2==0{
				out = append(out, strconv.Itoa(i/2))
			}else{
				out = append(out, ".")
			}
		} 	
	}
	return out 
}

func caclCheckSum(input []string) int{
	out := 0
	for i:=0; i<len(input); i++{
		if input[i]!="."{
			checkSum, _ := strconv.Atoi(input[i])
			out+=checkSum*i
		}
	}
	return out

}
