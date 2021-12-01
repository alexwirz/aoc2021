package main

import "testing"

func TestFoo(t *testing.T) {
	res := CheckIncreased([]int{1, 2})
	if res != 0 {
		t.Fatal("expected increase count to be 1")
	}
}
