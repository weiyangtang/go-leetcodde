package tree

import (
	"fmt"
	"testing"
)

func TestSerialize(t *testing.T) {
	codec := Constructor()
	serializeStr := "1,2,3,#,#,4,5"
	root := codec.deserialize(serializeStr)
	serialize := codec.serialize(root)
	fmt.Println(serialize)
}
