package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)


var markSummits [][]string
func resetMarkSummits(input [][]string){
	markSummits = make([][]string, len(input))
	for i:=0; i<len(markSummits); i++{
		markSummits[i]=make([]string, len(input[0]))
	}
}
func markTheSummit(loc []int){
	if checkWithinbounds(markSummits, loc){
		markSummits[loc[0]][loc[1]]= "."
	}
}
func checkIfMarked(loc []int) bool{
	if checkWithinbounds(markSummits, loc){
		return markSummits[loc[0]][loc[1]]=="."
	}
	return false
}



var debugArray [][]int
func resetDebugArray(input [][]string){
	debugArray = make([][]int, len(input))
	for i:=0; i<len(debugArray); i++{
		debugArray[i]=make([]int, len(input[0]))
	}
}
func markDebugArray(loc []int){
	if checkWithinbounds(debugArray, loc){
		debugArray[loc[0]][loc[1]]= 1
	}
}
func printDebugArray(){
	for _, line := range(debugArray){
		fmt.Println(line)
	}
}

func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	input, trailHeadLocs := GetInput(inputPath)
	fmt.Println(trailHeadLocs)
	fmt.Println(p1(trailHeadLocs, input))
}


func p1(trailHeadLocs [][]int, input [][]string)int{
	listOfTrailHeadScores := make([]int, 0)

	for _, loc := range(trailHeadLocs){
		resetDebugArray(input)
		resetMarkSummits(input)
		trailHeadScore := checkTrail(loc, input, loc, 0)
		fmt.Println(trailHeadScore)
		listOfTrailHeadScores = append(listOfTrailHeadScores, trailHeadScore)
	}
	return calcTrailHeadSum(listOfTrailHeadScores)
}
func checkTrail(trailHead []int, input [][]string, currLoc []int, expHeight int)int{
	markDebugArray(currLoc)
	// printDebugArray()
	// fmt.Println("--------------------")
	if !checkWithinbounds(input, currLoc){
		return 0
	}
	currHeight, _ := strconv.Atoi(input[currLoc[0]][currLoc[1]])
	// when the input is . the currHeight will be 0 and thatexp height is only possible at the start when we know the location is not . 
	if currHeight!=expHeight{
		return 0
	}
	if currHeight==9 && !checkIfMarked(currLoc){
		// fmt.Println("returned one")
		markTheSummit(currLoc)
		return 1
	}
	return (checkTrail(trailHead, input, []int{currLoc[0]+1, currLoc[1]}, expHeight+1)+ checkTrail(trailHead, input, []int{currLoc[0]-1, currLoc[1]}, expHeight+1)+ checkTrail(trailHead, input, []int{currLoc[0], currLoc[1]+1}, expHeight+1)+ checkTrail(trailHead, input, []int{currLoc[0], currLoc[1]-1}, expHeight+1))
}


func calcTrailHeadSum(input []int)int{
	out:=0
	MapFuncToList(input, func(num int)[]string{
		out+=num
		return []string{""}
	})
	return out
}
