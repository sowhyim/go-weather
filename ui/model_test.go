package ui

import (
	"fmt"
	"testing"
)

func TestModel(t *testing.T) {
	for key, val := range codes {
		fmt.Println(key, "  ", val)
	}
}
