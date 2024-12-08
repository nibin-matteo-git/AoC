package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)


var calibrationResult int = 0

func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	input := GetInput(inputPath)
	p1(input)
	fmt.Println(calibrationResult)
}

func p1(input []myStruct){
	var wg sync.WaitGroup
	
	for i:=0; i<len(input); i++{
		wg.Add(1)
		target := input[i].target
		values := input[i].values
		go func(target int, values []int){
			defer wg.Done()
			p1BruteForce(target, values, 0, 0, &wg)
		}(target, values)
	}
	wg.Wait()
}
// func p2BruteForce(target int, values []int, valIndex int, calcVal int, wg *sync.WaitGroup, operator string, valOnIndex int)bool{
// 	tmpP2Val := 0
// 	newIndex := valIndex+1
// 	if operator == "+"{
// 		calcVal += valOnIndex
// 	}
// 	if operator == "*"{
// 		calcVal*=valOnIndex
// 	}
// 	if operator == "||"{
// 		var _ error
// 		tmpP2Val, _ = strconv.Atoi(strconv.Itoa(valOnIndex)+strconv.Itoa(values[valIndex]))
// 	}
// 	if (operator =="+"||operator=="*") && newIndex<len(values){
// 		tmpP2Val = values[newIndex]
// 	}
// 	if (operator )
// 	if newIndex<len(values){
// 		return p2BruteForce(target, values,  newIndex, calcVal, wg, "+", tmpP2Val)|| p2BruteForce(target, values, newIndex, calcVal, wg, "*", tmpP2Val) || p2BruteForce(target, values, newIndex, calcVal, wg, "||", tmpP2Val)

// 	}else{
// 		if calcVal==target{
// 			calibrationResult+=target
// 			return true
// 		}
// 		return false
// 	}

// }


func p2BruteForce(target int, values []int, valIndex int, calcVal int, wg *sync.WaitGroup, valOnIndex int)bool{
	if valIndex<len(values){
		
		newIndex := valIndex+1
		tmpP2Val, _ := strconv.Atoi(strconv.Itoa(values[valIndex])+strconv.Itoa(values[newIndex]))
		return p2BruteForce(target, values,  newIndex, calcVal*values[valIndex], wg)|| p2BruteForce(target, values, newIndex, calcVal+values[valIndex], wg, 0)||p2BruteForce(target, values, newIndex+1, calcVal, wg, tmpP2Val)||p2BruteForce(target, values, newIndex+1, calcVal, wg, tmpP2Val)
	}else{
		if calcVal==target{
			calibrationResult+=target
			return true
		}
		return false
	}

}


func p1BruteForce(target int, values []int, valIndex int, calcVal int, wg *sync.WaitGroup)bool{
	if valIndex<len(values){
		newIndex := valIndex+1
		return p1BruteForce(target, values,  newIndex, calcVal*values[valIndex], wg)|| p1BruteForce(target, values, newIndex, calcVal+values[valIndex], wg)
	}else{
		if calcVal==target{
			calibrationResult+=target
			return true
		}
		return false
	}

}
