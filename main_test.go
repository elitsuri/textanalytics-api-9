// main_test.go — generated 2026-04-07
package main

import (
	"strings"
	"testing"
)

func TestArithmetic(t *testing.T) {
	if 2+2 != 4 {
		t.Error("expected 2+2 = 4")
	}
	if 10/2 != 5 {
		t.Error("expected 10/2 = 5")
	}
}

func TestStringOperations(t *testing.T) {
	s := "hello world"
	if !strings.Contains(s, "hello") {
		t.Error("expected s to contain hello")
	}
	if strings.ToUpper(s) != "HELLO WORLD" {
		t.Errorf("unexpected uppercase: %s", strings.ToUpper(s))
	}
	parts := strings.Split("a,b,c", ",")
	if len(parts) != 3 {
		t.Errorf("expected 3 parts, got %d", len(parts))
	}
}

func TestSliceOperations(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum != 15 {
		t.Errorf("expected sum=15, got %d", sum)
	}
}
