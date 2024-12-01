package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	cwd, _ := os.Getwd()
	inputPath := filepath.Join(cwd, "input.txt")
	input1, input2 := GetInput(inputPath)

}
