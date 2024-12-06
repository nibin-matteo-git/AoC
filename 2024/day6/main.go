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

	// possibleBlocks := p2(mappedArea, guardLoc)
	// fmt.Println(possibleBlocks)

	// distinctMapPosn := p1(mappedArea, guardLoc)

	// fmt.Println(getDistinctVisited(distinctMapPosn), guardLoc)
	
	rowTravel := -1
	colTravel := 0
	startDirection := direction{rowTravel, colTravel}
	directionMap:= make([][]direction, len(mappedArea))
	for i:=range(directionMap){
		directionMap[i]=make([]direction, len(mappedArea[0]))
	}
	fmt.Println(p2BruteForce(mappedArea, guardLoc, directionMap, startDirection, true))
}
func checkDirectionMatch(directionMap direction, otherDir direction)bool{
	return directionMap.vert==otherDir.vert && directionMap.horiz==otherDir.horiz	
}
 
func p2BruteForce(inputMap [][]string, guardLoc []int, directionMap[][]direction, guardDir direction, addBlock bool)int{
	out:=0
	directionMap[guardLoc[0]][guardLoc[1]]= guardDir
	for (true){
		predLoc := []int{guardLoc[0]+guardDir.vert, guardLoc[1]+guardDir.horiz}
		fmt.Println("pred location:", predLoc)
		if predLoc[0]<0 || predLoc[0]>=len(inputMap) || predLoc[1]<0 ||predLoc[1]>=len(inputMap[0]){
			break
		}
		tmpGuardDir := getGuardDirection(guardDir)
		if (directionMap[predLoc[0]][predLoc[1]]!=direction{0,0}) && (checkDirectionMatch(directionMap[predLoc[0]][predLoc[1]], tmpGuardDir)){
			// fmt.Println("it matches the direction so we put a block so can turn")
			out++
		}
		if inputMap[predLoc[0]][predLoc[1]]=="#"{
			guardDir = tmpGuardDir
			continue
		}
		if addBlock{
			tmpVar := inputMap[predLoc[0]][predLoc[1]]
			inputMap[predLoc[0]][predLoc[1]]="#"
			fmt.Println("GOING INSIDE RECURSION")
			out+=p2BruteForce(inputMap, guardLoc, directionMap, guardDir, false)
			inputMap[predLoc[0]][predLoc[1]]=tmpVar 
			fmt.Println("OUTSIDE RECURSION")
		}

		// fmt.Println("new location:", predLoc)
		directionMap[guardLoc[0]][guardLoc[1]]=guardDir
		guardLoc=predLoc

	}
	return out
}
// func p2(inputMap [][]string, guardLoc []int)int{
// 	directionMap:= make([][]direction, len(inputMap))
// 	out:=0
// 	for i:=range(directionMap){
// 		directionMap[i]=make([]direction, len(inputMap[0]))
// 	}
// 	rowTravel := -1
// 	colTravel := 0
// 	directionMap[guardLoc[0]][guardLoc[1]]=direction{rowTravel, colTravel}
// 	for (true){
// 		predLoc := []int{guardLoc[0]+rowTravel, guardLoc[1]+colTravel}
// 		fmt.Println("pred location:", predLoc)
// 		if predLoc[0]<0 || predLoc[0]>=len(inputMap) || predLoc[1]<0 ||predLoc[1]>=len(inputMap[0]){
// 			break
// 		}
// 		tmpRow, tmpCol := getGuardDirection(rowTravel, colTravel)
// 		if (directionMap[predLoc[0]][predLoc[1]]!=direction{0,0}) && (checkDirectionMatch(directionMap[predLoc[0]][predLoc[1]], tmpRow, tmpCol)){
// 			fmt.Println("it matches the direction so we put a block so can turn")
// 			out++
// 		}
// 		if inputMap[predLoc[0]][predLoc[1]]=="#"{
// 			rowTravel, colTravel =tmpRow, tmpCol
// 			continue
// 		}
// 		// fmt.Println("new location:", predLoc)
// 		guardLoc=predLoc
// 		directionMap[guardLoc[0]][guardLoc[1]]=direction{rowTravel, colTravel}

// 	}
// 	return out
// }
// func p1(inputMap [][]string, guardLoc []int)[][]bool{
// 	out:= make([][]bool, len(inputMap))
// 	for i:=range(out){
// 		out[i]=make([]bool, len(inputMap[0]))
// 	}
// 	rowTravel := -1
// 	colTravel := 0
// 	out[guardLoc[0]][guardLoc[1]]=true
// 	for (true){
// 		predLoc := []int{guardLoc[0]+rowTravel, guardLoc[1]+colTravel}
// 		if predLoc[0]<0 || predLoc[0]>=len(inputMap) || predLoc[1]<0 ||predLoc[1]>=len(inputMap[0]){
// 			break
// 		}
// 		if inputMap[predLoc[0]][predLoc[1]]=="#"{
// 			rowTravel, colTravel = getGuardDirection(rowTravel, colTravel)
// 			continue
// 		}
// 		// fmt.Println("new location:", predLoc)
// 		guardLoc=predLoc
// 		out[guardLoc[0]][guardLoc[1]]=true

// 	}
// 	return out
// }
func getGuardDirection(input direction)(direction){
	row, col := input.vert, input.horiz
	switch{
	case row==-1:
		return direction{0, 1}
	case col==1:
		return direction{1, 0}
	case row==1:
		return direction{0, -1}
	case col==-1:
		return direction{-1, 0}
	}
	fmt.Println("something went terribly wrong in getting guard new direction")
	return direction{-1, -1}
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
