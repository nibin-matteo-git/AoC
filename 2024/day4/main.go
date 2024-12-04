package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	reports := GetInput(inputPath)
	fmt.Println(part2(reports))

}
func part2(input [][]string) int {
	out := 0
	for i, char1 := range input {
		for j, char2 := range char1 {
			if char2 == "M" || char2 == "S" {
				if p2ChckXMas(input, i, j) {
					out += 1
				}
			}
		}
	}
	return out
}
func p2ChckXMas(input [][]string, ind1, ind2 int) bool {
	out := 0
	var possXmas []strings.Builder = make([]strings.Builder, 2)
	for i := 0; i < len("MAS"); i++ {
		writeToStringBuilder(input, &possXmas, 0, ind1+i, ind2+i)
		writeToStringBuilder(input, &possXmas, 1, ind1+2-i, ind2+i)
	}
	MapFuncToList(possXmas, func(input strings.Builder) string {
		if input.String() == "MAS" || input.String() == "SAM" {
			out += 1
		}
		return input.String()
	})
	return out == 2
}

func checkXMas(input [][]string) int {
	out := 0
	for i, char1 := range input {
		for j, char2 := range char1 {
			if char2 == "X" {
				//up_d1_right_d2_down_d3_left_d4
				var possXmas []strings.Builder = make([]strings.Builder, 8)
				for k := 0; k < len("XMAS"); k++ {
					writeToStringBuilder(input, &possXmas, 0, i, j-k)
					writeToStringBuilder(input, &possXmas, 1, i+k, j-k)
					writeToStringBuilder(input, &possXmas, 2, i+k, j)
					writeToStringBuilder(input, &possXmas, 3, i+k, j+k)
					writeToStringBuilder(input, &possXmas, 4, i, j+k)
					writeToStringBuilder(input, &possXmas, 5, i-k, j+k)
					writeToStringBuilder(input, &possXmas, 6, i-k, j)
					writeToStringBuilder(input, &possXmas, 7, i-k, j-k)
				}
				MapFuncToList(possXmas, func(input strings.Builder) string {
					if input.String() == "XMAS" {
						out += 1
					}
					return input.String()
				})

			}
		}
	}
	return out
}

func writeToStringBuilder(input [][]string, out *[]strings.Builder, outInd, ind1, ind2 int) {
	if ind1 >= 0 && ind1 < len(input) && ind2 >= 0 && ind2 < len(input[0]) {
		(*out)[outInd].WriteString(input[ind1][ind2])
	}
}
