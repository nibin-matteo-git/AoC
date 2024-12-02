package main

import (
	"log"
	"os"
	"strings"
)

func GetInput(inputFile string) []string {
	contents, err := os.ReadFile(inputFile)
	HandleError(err)
	result := strings.Split(string(contents), "\n")
	out := make([]string, 0)
	for _, num := range result {
		num := strings.TrimSpace(num)
		if num == "" {
			continue
		}
		out = append(out, num)
	}
	return out
}
func HandleError(err error) {
	if err != nil {
		log.Fatalf("Encountered error: %s", err.Error())
	}
}
