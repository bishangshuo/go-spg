package utils

import (
	"fmt"
	"testing"
)

func TestGetArgBool(t *testing.T) {
	reindex, _ := GetArgBool("reindex")
	fmt.Println(reindex)
}
