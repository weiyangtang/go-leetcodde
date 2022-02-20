package binarySearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLeftBoundBinSearch(t *testing.T) {
	arr := []int{1, 2, 2, 4}
	target := 2
	index := leftBoundBinSearch(arr, target)
	assert.Equal(t, 1, index)

	target = 0
	index = leftBoundBinSearch(arr, target)
	assert.Equal(t, -1, index)

	target = 5
	index = leftBoundBinSearch(arr, target)
	assert.Equal(t, -1, index)

	target = 4
	index = leftBoundBinSearch(arr, target)
	assert.Equal(t, 3, index)
}

func TestRightBoundBinSearch(t *testing.T) {
	arr := []int{1, 2, 2, 4}
	target := 2
	index := rightBoundBinSearch(arr, target)
	assert.Equal(t, 2, index)

	target = 0
	index = rightBoundBinSearch(arr, target)
	assert.Equal(t, -1, index)

	target = 5
	index = rightBoundBinSearch(arr, target)
	assert.Equal(t, -1, index)

	target = 4
	index = rightBoundBinSearch(arr, target)
	assert.Equal(t, 3, index)
}
