package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type direction struct{
	vert int 
	horiz int 
}

func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	mappedArea, guardLoc := GetInput(inputPath)
	fmt.Println(mappedArea)

	possibleBlocks := p2(mappedArea, guardLoc)
	fmt.Println(possibleBlocks)

	// distinctMapPosn := p1(mappedArea, guardLoc)

	// fmt.Println(getDistinctVisited(distinctMapPosn), guardLoc)
	
}
func checkDirectionMatch(directionMap direction, rowTravel, colTravel int)bool{
	return directionMap.vert==rowTravel && directionMap.horiz==colTravel	
}
func p2(inputMap [][]string, guardLoc []int)int{
	directionMap:= make([][]direction, len(inputMap))
	out:=0
	for i:=range(directionMap){
		directionMap[i]=make([]direction, len(inputMap[0]))
	}
	rowTravel := -1
	colTravel := 0
	directionMap[guardLoc[0]][guardLoc[1]]=direction{rowTravel, colTravel}
	for (true){
		predLoc := []int{guardLoc[0]+rowTravel, guardLoc[1]+colTravel}
		fmt.Println("pred location:", predLoc)
		if predLoc[0]<0 || predLoc[0]>=len(inputMap) || predLoc[1]<0 ||predLoc[1]>=len(inputMap[0]){
			break
		}
		tmpRow, tmpCol := getGuardDirection(rowTravel, colTravel)
		if (directionMap[predLoc[0]][predLoc[1]]!=direction{0,0}) && (checkDirectionMatch(directionMap[predLoc[0]][predLoc[1]], tmpRow, tmpCol)){
			fmt.Println("it matches the direction so we put a block so can turn")
			out++
		}
		if inputMap[predLoc[0]][predLoc[1]]=="#"{
			rowTravel, colTravel =tmpRow, tmpCol
			continue
		}
		// fmt.Println("new location:", predLoc)
		guardLoc=predLoc
		directionMap[guardLoc[0]][guardLoc[1]]=direction{rowTravel, colTravel}

	}
	return out
}
func p1(inputMap [][]string, guardLoc []int)[][]bool{
	out:= make([][]bool, len(inputMap))
	for i:=range(out){
		out[i]=make([]bool, len(inputMap[0]))
	}
	rowTravel := -1
	colTravel := 0
	out[guardLoc[0]][guardLoc[1]]=true
	for (true){
		predLoc := []int{guardLoc[0]+rowTravel, guardLoc[1]+colTravel}
		if predLoc[0]<0 || predLoc[0]>=len(inputMap) || predLoc[1]<0 ||predLoc[1]>=len(inputMap[0]){
			break
		}
		if inputMap[predLoc[0]][predLoc[1]]=="#"{
			rowTravel, colTravel = getGuardDirection(rowTravel, colTravel)
			continue
		}
		// fmt.Println("new location:", predLoc)
		guardLoc=predLoc
		out[guardLoc[0]][guardLoc[1]]=true

	}
	return out
}
func getGuardDirection(row, col int)(int, int){
	switch{
	case row==-1:
		return 0, 1
	case col==1:
		return 1, 0
	case row==1:
		return 0, -1
	case col==-1:
		return -1, 0
	}
	fmt.Println("something went terribly wrong in getting guard new direction")
	return -1, -1
}
func getDistinctVisited(input [][]bool) int{
	out:=0
	for _, num1 := range(input){
		for _, num2 := range(num1){
			if num2 ==true{
				out+=1
			}
		}
	}
	return out
}
