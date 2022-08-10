package algorithm

import (
	. "github.com/AlexandreChamard/go-builtin"
)

func Abs[N Number](n N) N {
	if n < 0.0 {
		return -n
	} else {
		return n
	}
}

// Returns the lowest element
func Min[T Ord](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

// Returns the lowest element
func MinN[T Ord](xs ...T) T {
	min := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] < min {
			min = xs[i]
		}
	}
	return min
}

// Returns the biggest element
func Max[T Ord](a, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}
}

// Returns the biggest element
func MaxN[T Ord](xs ...T) T {
	max := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > max {
			max = xs[i]
		}
	}
	return max
}

// Generic sort with a predicate
func Sort[T any](slice []T, comp Ordf[T]) {
	QuickSort(slice, comp)
}

// Generic QuickSort with a predicate
func QuickSort[T any](slice []T, comp Ordf[T]) {
	quickSort(slice, comp, 0, len(slice)-1)
}

func quickSort[T any](slice []T, comp Ordf[T], low, high int) {
	if len(slice) <= 1 {
		return
	}
	if low < high {
		// pi is partitioning index, slice[p] is now at right place
		pi := Partition(slice, comp, low, high)
		quickSort(slice, comp, low, pi-1)  // Before pi
		quickSort(slice, comp, pi+1, high) // After pi
	}
}

func Partition[T any](slice []T, comp Ordf[T], low, high int) int {
	i := (low - 1)       // index of smaller element
	pivot := slice[high] // pivot

	for j := low; j < high; j++ {
		// if current element is smaller than or equal to pivot
		if comp(slice[j], pivot) {
			// increment index of smaller element
			i = i + 1
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

// Returns true if all the elements are equals
func SliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] == b[i] {
			return false
		}
	}
	return true
}

func SliceEqualF[T any](a, b []T, comp Compf[T]) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if comp(a[i], b[i]) {
			return false
		}
	}
	return true
}

// Returns true if ALL pred(element) = true
func All[T any](slice []T, pred Eqf[T]) bool {
	for _, a := range slice {
		if !pred(a) {
			return false
		}
	}
	return true
}

// Returns true if AT LEASE ONE pred(element) = true
func Any[T any](slice []T, pred Eqf[T]) bool {
	for _, a := range slice {
		if pred(a) {
			return true
		}
	}
	return false
}

// Returns true if all elements are unique
func Unique[T comparable](slice []T) bool {
	m := make(map[T]bool, len(slice))
	for _, a := range slice {
		if m[a] {
			return false
		}
		m[a] = true
	}
	return true
}

// Same as Unique with a predicate
func UniqueF[T any](slice []T, pred Compf[T]) bool {
	for i := range slice {
		for j := i + 1; j < len(slice); j++ {
			if pred(slice[i], slice[j]) {
				return false
			}
		}
	}
	return true
}
