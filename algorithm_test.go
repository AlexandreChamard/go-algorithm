package algorithm

import (
	"testing"

	. "github.com/AlexandreChamard/go-builtin"
)

func TestAbs(t *testing.T) {

	if Abs(int(-1)) != 1 {
		t.Errorf("expected %v got %v", 1, Abs(int(-1)))
	}
	if Abs(int(1)) != 1 {
		t.Errorf("expected %v got %v", 1, Abs(int(1)))
	}
	if Abs(float64(-0.5)) != 0.5 {
		t.Errorf("expected %v got %v", 0.5, Abs(float64(-0.5)))
	}
	if Abs(float64(0.5)) != 0.5 {
		t.Errorf("expected %v got %v", 0.5, Abs(float64(0.5)))
	}
	if Abs(float64(-0.0)) != 0.0 {
		t.Errorf("expected %v got %v", 0.0, Abs(float64(-0.0)))
	}
}

func TestSort(t *testing.T) {
	sliceInt := []int{42, 12, 4, -5, 0, 69, 20}
	Sort(sliceInt, Less[int])

	intResult := []int{-5, 0, 4, 12, 20, 42, 69}
	if SliceEqualF(sliceInt, intResult, Equal[int]) {
		t.Fatalf("expected %v got %v", intResult, sliceInt)
	}
}

func TestUnique(t *testing.T) {
	type a struct {
		i int
	}
	type b struct {
		i, j int
	}

	eqf := func(a, b b) bool { return a.j == b.j }

	slice1 := []a{{42}, {12}, {4}, {-5}, {0}, {69}, {20}}
	slice2 := []b{{0, 42}, {1, 12}, {2, 4}, {3, -5}, {4, 0}, {5, 69}, {6, 4}, {7, 20}}
	if !Unique(slice1) {
		t.Fatalf("Unique(%v): expected %v got %v", slice1, true, false)
	}
	if !Unique(slice2) {
		t.Fatalf("Unique(%v): expected %v got %v", slice1, true, false)
	}
	if UniqueF(slice2, eqf) {
		t.Fatalf("Unique(%v): expected %v got %v", slice2, false, true)
	}
}
