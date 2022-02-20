package dfs

import (
	"fmt"
	"testing"
)

func TestLongestDiverseString(t *testing.T) {
	diverseString := longestDiverseString(7, 1, 0)
	fmt.Println(diverseString)
}
