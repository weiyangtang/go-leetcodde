package dp

import (
	"fmt"
	"testing"
)

func TestMinStickers(t *testing.T) {
	stickers := minStickers([]string{"with", "example", "science"}, "thehat")
	fmt.Println(stickers)
}
