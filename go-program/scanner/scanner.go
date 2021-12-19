package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if strings.EqualFold("EOF", line) {
			break
		}
		if _, ok := seen[line]; !ok {
			seen[line] = true
		}
	}
	for k, v := range seen {
		fmt.Printf("k:%v,v:%v\n", k, v)
	}
}
