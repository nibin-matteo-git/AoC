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
	input := GetInput(inputPath)
	part2(input)
}
func part2(input []string){
	var out int
	for _, line := range input {
		pwd := strings.Split(line, ":")[1]
		pwd = strings.TrimSpace(pwd)
		// how to do the above in one line

		policy := strings.Split(line, ":")[0]
		policyChar := strings.Split(policy, " ")[1]
		firstOcc, _ := strconv.Atoi(strings.Split(strings.Split(policy, " ")[0], "-")[0])
		secondOcc, _ := strconv.Atoi(strings.Split(strings.Split(policy, " ")[0], "-")[1])

		bool1 := string(pwd[firstOcc-1])==policyChar
		bool2 := string(pwd[secondOcc-1])== policyChar

		if  (bool1||bool2) && !(bool1 && bool2) {
			out += 1
		}

	}
	fmt.Println(out)

}
func part1(input []string){
	var out int
	for _, line := range input {
		pwd := strings.Split(line, ":")[1]
		pwd = strings.TrimSpace(pwd)
		// how to do the above in one line

		policy := strings.Split(line, ":")[0]
		policyChar := strings.Split(policy, " ")[1]
		policyMin, _ := strconv.Atoi(strings.Split(strings.Split(policy, " ")[0], "-")[0])	
		policyMax, _ := strconv.Atoi(strings.Split(strings.Split(policy, " ")[0], "-")[1])

		if charOcc := strings.Count(pwd, policyChar); charOcc <= policyMax && charOcc >= policyMin {
			out += 1
		}

	}
	fmt.Println(out)

}
