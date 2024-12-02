package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	reports := GetInput(inputPath)
	result := 0
	for _, report := range reports {
		vals := strings.Split(report, " ")
		for i:=0; i<len(vals); i++{
			valsC := make([]string, len(vals))
			copy(valsC, vals)
			if isReportSafe(append(valsC[:i], valsC[i+1:]...)){
				result+=1
				break
			}

		}
	}
	fmt.Println(result)

}
func isReportSafe(vals []string)bool{
	intVals := MapFuncToList(vals, func(item string)int{
		out, _ := strconv.Atoi(item)
		return out
	})
	c1:=make([]int,len(intVals))
	c2:=make([]int, len(intVals))
	copy(c1, intVals)
	copy(c2, intVals)
	sort.Ints(c1)
	sort.Slice(c2, func(a, b int)bool{return intVals[a]>intVals[b]})
	if reflect.DeepEqual(c1, intVals) || reflect.DeepEqual(c2, intVals){
		for i:=1; i<len(intVals); i++{
			if math.Abs(float64(intVals[i]-intVals[i-1]))>3 || math.Abs(float64(intVals[i]-intVals[i-1]))<1{
				return false
			}
		}
		return true
	}

	return false
}
