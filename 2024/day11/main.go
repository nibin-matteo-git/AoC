package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
)

var GOMAXPROCS = runtime.NumCPU()


func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	input := GetInput(inputPath)
	blinks := 75
	fmt.Println("no of cores: ", GOMAXPROCS)
	for i:=0; i<blinks; i++{
		input = p1(input)
		// fmt.Println(input)
		fmt.Println("iter: ", i, ";len: ", len(input))
	}
	fmt.Println(len(input))
}


func p1(input []string)[]string{
	chunkSize := (len(input)/GOMAXPROCS)
	if chunkSize==0{
		chunkSize=1
	}
	noOfRoutines := (len(input)/chunkSize)+1
	out := make([][]string, noOfRoutines) 
	// var mutex sync.Mutex
	var wg sync.WaitGroup
	for i:=0; i<noOfRoutines; i++{
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			start:=i*chunkSize
			end := (i+1)*chunkSize
			if end>len(input){
				end = len(input)
			}
			localOut := make([]string, 0, chunkSize)
			for j:=start; j<end; j++{
				num := input[j]
				if num=="0"{
					localOut = append(localOut, strconv.Itoa(1))
				}else if len(num)%2==0{
					stone1 := num[:len(num)/2]
					stone2 := num[len(num)/2:]
					if num[0]=='0'{
						stone1 = removeLeadZeros(stone1)
					}
					if num[len(num)/2]=='0'{
						stone2 = removeLeadZeros(stone2)
					}
					localOut = append(localOut, stone1)
					localOut = append(localOut, stone2)
				}else{
					localOut = append(localOut, strconv.Itoa(convertStringToInt(num)*2024))
				}
			}
			// mutex.Lock()
			out[i]=localOut
			// mutex.Unlock()
		}(i)
	}

	wg.Wait()

	finalOut := combineNestedLoops(out, (len(input)))


	return finalOut
}

func removeLeadZeros(input string)string{
	for i, val := range(input){
		if val!='0'{
			return input[i:]
		}
	}
	return "0"
}
