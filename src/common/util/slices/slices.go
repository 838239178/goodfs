package slices

import (
	"common/util/math"
)

func StringsReplace(arr []string, origin string, target string) bool {
	for i, a := range arr {
		if a == origin {
			arr[i] = target
			return true
		}
	}
	return false
}

func Fill[T any](arr []T, v T) {
	if len(arr) == 0 {
		return
	}
	// Preload the first value into the array/slice
	arr[0] = v

	// Incrementally duplicate the value into the rest of the container
	for j := 1; j < len(arr); j *= 2 {
		copy(arr[j:], arr[:j])
	}
}

func FillPattern[T any](arr []T, pattern []T) {
	if len(arr) == 0 || len(pattern) == 0 {
		return
	}
	// Copy the pattern into the start of the container
	copy(arr, pattern)

	// Incrementally duplicate the pattern throughout the container
	for j := len(pattern); j < len(arr); j *= 2 {
		copy(arr[j:], arr[:j])
	}
}

func First[T any](arr []T) T {
	return arr[0]
}

func Last[T any](arr []T) T {
	return arr[len(arr)-1]
}

func Clear[T any](arr *[]T) {
	*arr = (*arr)[:0]
}

func RemoveAll[T any](arr []T) []T {
	return arr[:0]
}

func RemoveFirst[T any](arr *[]T) {
	*arr = (*arr)[1:]
}

func RemoveHead[T any](arr []T) []T {
	return arr[1:]
}

func ExtremalIndex[T any](arr []T, fn func(ex, cur T) bool) int {
	var exIdx int
	for i := range arr {
		if fn(arr[exIdx], arr[i]) {
			exIdx = i
		}
	}
	return exIdx
}

func Extremal[T any](arr []T, fn func(ex, cur T) bool) T {
	return arr[ExtremalIndex(arr, fn)]
}

func Search[T comparable](arr []T, target T) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func RemoveLast[T any](arr *[]T) {
	*arr = (*arr)[:len(*arr)-1]
}

// SafeChunk [start, end], negative number means counting from tail to head
func SafeChunk[T any](arr []T, start, end int) []T {
	// a mod b = a mod (-b)
	// a mod b = - (-a mod b)

	if start < 0 {
		start = math.LogicMod(start, len(arr))
	} else if start >= len(arr) {
		start = len(arr) - 1
	}

	if end < 0 {
		end = math.LogicMod(end, len(arr))
	} else if end >= len(arr) {
		end = len(arr) - 1
	}

	if start > end {
		start, end = end, start
	}

	return arr[start : end+1]
}
