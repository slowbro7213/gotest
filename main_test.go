package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	sumRet := sum(1, 3)

	if sumRet != 3 {
		t.Error("sum is not 3")
	}
}
