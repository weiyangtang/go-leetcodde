package link

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	lruCache := Constructor(3)
	lruCache.Put(1, 3)
	lruCache.Put(2, 4)
	lruCache.Put(3, 4)
	lruCache.Put(4, 5)
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(3))
	fmt.Println(lruCache.Get(4))

}
