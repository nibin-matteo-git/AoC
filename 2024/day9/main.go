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
	fmt.Println(caclCheckSum(p2(input, expandedInput)))
	// fmt.Println(caclCheckSum(p1(expandedInput)))
}

func p2(rawInput string, inputt []string)[]string{
	input := make([]string, len(inputt))
	copy(input, inputt)
	emptLocs := p2GetEmptLocs(rawInput)
	for i:=len(input)-1; i>=0; i--{
		if input[i]!="."{
			// fmt.Println(input)
			fileIndex, _ := strconv.Atoi(input[i])
			lenOfFile, _ := strconv.Atoi(string(rawInput[fileIndex*2]))
			for j:=0; j<len(emptLocs); j++{
				if emptLocs[j][1]>=int(lenOfFile) && emptLocs[j][0]<i{
					for k:=0; k<int(lenOfFile); k++{
						input[emptLocs[j][0]]=input[i]
						input[i]="."
						emptLocs[j][1]-=1
						emptLocs[j][0]+=1
						i-=1
					}
					i++
					break
				}
			}
		}
	}
	return input
}


// output of this func is [[<start index of empty space>, length of free space]]
func p2GetEmptLocs(rawInput string)[][]int{
	emptLocs := make([][]int, 0)
	expandedInputIndex := 0
	for i, val := range(rawInput){
		indInc, _ := strconv.Atoi(string(val))
		if i%2==0{
			expandedInputIndex+=indInc
		}else{
			emptLocs = append(emptLocs, []int{expandedInputIndex, indInc})
			expandedInputIndex+=indInc
		}
	}
	return emptLocs
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
