package go_strc

import (
	"math/rand"
)

// RandomSubset takes a random sub-slice of length n from a given slice s
func RandomSubset[T any](s []T, n int) []T {
	if n > len(s) {
		n = len(s)
	}

	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })

	return s[:n]
}

// Reverse reverses a slice
func Reverse[T any](s []T) []T {
	reversed := make([]T, len(s))
	copy(reversed, s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return reversed
}

// Equal compares two slices contents and returns true if they are the same, false if they are not
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CountDuplicates[T comparable](s []T) int {
	duplicates := 0
	counted := Counts(s)
	for c := range counted {
		if counted[c] > 1 {
			duplicates++
		}
	}
	return duplicates
}

func Counts[V comparable](vs []V) map[V]int {
	return CountsFunc(vs, func(v V) V { return v })
}

func CountsFunc[V any, K comparable](vs []V, key func(V) K) map[K]int {
	h := make(map[K]int)
	for _, v := range vs {
		h[key(v)]++
	}
	return h
}
