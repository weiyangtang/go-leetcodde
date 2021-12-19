package github

import (
	"fmt"
	"testing"
)

func TestGithub(t *testing.T) {
	issues, err := SearchIssues([]string{"json", "decoder"})
	if err != nil {
		return
	}
	fmt.Println(issues)
}
