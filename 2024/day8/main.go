package main

import (
	"fmt"
	"os"
	"path/filepath"
)
type locnStruct struct{
	index1 int
	index2 int
}
var locStorage map[locnStruct]int = make(map[locnStruct]int)

func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	input, rawInput := GetInput(inputPath)
	p2(rawInput, input)
}

func p2(rawInput []string, input map[string][][]int)int{
	out:=0
	for _, indices := range(input){
		out+=possAntinodesP2(rawInput, indices)
		// fmt.Println("char: ",c)
		// fmt.Println("the number of possible antinodes:",out)
	}
	fmt.Println("overall antidotes:", out)
	return out
}
func p1(rawInput []string, input map[string][][]int)int{
	out:=0
	for _, indices := range(input){
		out+=possibleAntinodesForChar(rawInput, indices)
		// fmt.Println("char: ",c)
		// fmt.Println("the number of possible antinodes:",out)
	}
	fmt.Println("overall antidotes:", out)
	return out
}


func possAntinodesP2(rawInput []string, input [][]int)int{
	out:=0

	for i, index1 := range(input){
		for j:=i+1; j<len(input); j++{
			index2 := input[j]
			for r:=0; (r<len(rawInput) || r<len(rawInput[0])); r++{
				a1:=[]int{index1[0]+r*findDiffInPosn(index1, index2)[0], index1[1]+r*findDiffInPosn(index1, index2)[1]}
				a2:=[]int{index2[0]+r*findDiffInPosn(index2, index1)[0], index2[1]+r*findDiffInPosn(index2, index1)[1]}
				if checkLocnWithinBounds(rawInput, a1){
					_, present := locStorage[locnStruct{a1[0], a1[1]}]
					if !present{
						fmt.Println(a1)
						out+=1
					}
					locStorage[locnStruct{a1[0], a1[1]}]+=1
				}
				if checkLocnWithinBounds(rawInput, a2){
					_, present := locStorage[locnStruct{a2[0], a2[1]}]
					if !present{
						fmt.Println(a2)
						out+=1
					}
					locStorage[locnStruct{a2[0], a2[1]}]+=1
				}
			}
		}
	}
	return out
}
func possibleAntinodesForChar(rawInput []string, input [][]int)int{
	out:=0

	for i, index1 := range(input){
		for j:=i+1; j<len(input); j++{
			index2 := input[j]
			a1, a2 := getAntiNodeLocns(index1, index2)	
			if checkLocnWithinBounds(rawInput, a1){
				_, present := locStorage[locnStruct{a1[0], a1[1]}]
				if !present{
					out+=1
				}
				locStorage[locnStruct{a1[0], a1[1]}]+=1
			}
			if checkLocnWithinBounds(rawInput, a2){
				_, present := locStorage[locnStruct{a2[0], a2[1]}]
				if !present{
					out+=1
				}
				locStorage[locnStruct{a2[0], a2[1]}]+=1
			}
		}
	}
	return out
}

func checkAntinodeLocnFreeOfAntenna(rawInput []string, index []int)bool{
	return string(rawInput[index[0]][index[1]])=="."
}

func checkLocnWithinBounds(rawInput []string, index []int)bool{
	if index[0]>=0 && index[0]<len(rawInput) && index[1]>=0 && index[1]<len(rawInput[0]){
		return true
	}else{
		return false
	}
}

func getAntiNodeLocns(index1, index2 []int)([]int, []int){
	anti1:=[]int{index1[0]+findDiffInPosn(index1, index2)[0], index1[1]+findDiffInPosn(index1, index2)[1]}
	anti2:=[]int{index2[0]+findDiffInPosn(index2, index1)[0], index2[1]+findDiffInPosn(index2, index1)[1]}
	return anti1, anti2
}

func findDiffInPosn(posn1 []int, posn2 []int)[]int{
	x := posn1[0]-posn2[0]
	y := posn1[1]-posn2[1]
	return []int{x,y}
}
