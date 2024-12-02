package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	reports := GetInput(inputPath)
	result := 0
	for _, report := range reports {
		if isReportSafe(report, 1) {
			result += 1
		}
	}
	fmt.Println(result)

}
func isReportSafe(report string, tolerate int ) bool {
	vals := strings.Split(report, " ")
	fmt.Println("starting:",vals)
	result := true
	var tbd1 int
	var tbd2 int
	index := 0
	for tbd2-tbd1==0{
		tbd1, _ = strconv.Atoi(vals[index])
		tbd2, _ = strconv.Atoi(vals[index+1])
		if index+=1; index<len(vals)-1{
			if tbd2-tbd1==0{
				return false
			}
			break
		} 
	}
	if tbd2-tbd1 > 0 {
		for i := 1; i < len(vals); i++ {
			a, _ := strconv.Atoi(vals[i])
			b, _ := strconv.Atoi(vals[i-1])
			if a-b < 1 || a-b > 3 {
				if tolerate > 0 {
					tolerate -= 1
					//how does the spread operator work in golang
					//why does this line alter vals value append(vals[:i], vals[(i+1):]...)
					valsC1:=make([]string, len(vals))
					valsC2:=make([]string, len(vals))
					copy(valsC1, vals)
					copy(valsC2, vals)
					return isReportSafe(strings.Join(append(valsC1[:i], valsC1[(i+1):]...)," "),tolerate)||isReportSafe(strings.Join(append(valsC2[:i-1], valsC2[i:]...)," "),tolerate)
				} else {
					fmt.Println("false")
					return false
				}	
			}
		}
	} else {
		for i := 1; i < len(vals); i++ {
			a, _ := strconv.Atoi(vals[i])
			b, _ := strconv.Atoi(vals[i-1])
			if b-a < 1 || b-a > 3 {
				if tolerate > 0 {
					tolerate -= 1
					//how does the spread operator work in golang
					//why does this line alter vals value append(vals[:i], vals[(i+1):]...)
					valsC1:=make([]string, len(vals))
					valsC2:=make([]string, len(vals))
					copy(valsC1, vals)
					copy(valsC2, vals)
					return isReportSafe(strings.Join(append(valsC1[:i], valsC1[(i+1):]...)," "),tolerate)||isReportSafe(strings.Join(append(valsC2[:i-1], valsC2[i:]...)," "),tolerate)
				} else {
					fmt.Println("false")
					return false
				}	
			}
		}
	}
	return result
}
