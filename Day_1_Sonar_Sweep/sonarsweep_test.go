package main

import "testing"

func TestSecondGreaterThanFirst(t *testing.T) {
	res := CheckIncreased([]int{1, 2})
	if res != 1 {
		t.Fatal("expected increase count to be 1")
	}
}

func TestSecondSameAsFirst(t *testing.T) {
	res := CheckIncreased([]int{1, 1})
	if res != 0 {
		t.Fatal("expected increase count to be 0")
	}
}

func TestSecondLessThanFirst(t *testing.T) {
	res := CheckIncreased([]int{12, 1})
	if res != 0 {
		t.Fatal("expected increase count to be 0")
	}
}

func TestThirdGreaterThatFirst(t *testing.T) {
	res := CheckIncreased([]int{12, 1, 14})
	if res != 1 {
		t.Fatal("expected increase count to be 0")
	}
}

func TestWave(t *testing.T) {
	res := CheckIncreased([]int{12, 1, 14, 10, 42})
	if res != 2 {
		t.Fatal("expected increase count to be 0")
	}
}
