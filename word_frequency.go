package main

import "fmt"

import (
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := make(map[string]int)
	for _, word := range words {
		count[word]++
	}
	return count
}

func main() {
	fmt.Println(WordCount("She sells sea shells on the sea shore."))
}

