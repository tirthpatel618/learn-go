package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	dict := map[string]int{}
	fields := strings.Fields(s)
	for _, val := range fields {
		if _, ok := dict[val]; ok {
			dict[val] += 1
		} else {
			dict[val] = 1
		}
	}	
	return dict
}

func main() {
	wc.Test(WordCount)
}
