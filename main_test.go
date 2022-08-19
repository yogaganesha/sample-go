package main

import (
	"testing"
)

func TestForTest(t *testing.T) {
	r := forTest(1, 1)
	if r != 11 {
		t.Error("expected return value of 100, got ", r)
	}
}

func TestForTest2(t *testing.T) {
	r := forTest(0, 1)
	if r != 10 {
		t.Error("expected return value of 100, got ", r)
	}
}
