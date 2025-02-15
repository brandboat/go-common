package gocommon

import "slices"

// SliceContentCmp compares two slices and returns true if they have the same content with any order.
func SliceContentCmp[T comparable](x, y []T) bool {
	if len(x) == 0 && len(y) == 0 {
		return true
	}
	if len(x) != len(y) {
		return false
	}

	yMap := make(map[T]int, len(y))
	for _, item := range y {
		yMap[item]++
	}

	for _, xItem := range x {
		if counter, exist := yMap[xItem]; exist {
			if counter == 0 {
				return false
			}
			yMap[xItem]--
		} else {
			return false
		}
	}
	return true
}

// SliceDedupe removes duplicated items and return a non-deplucated items slice.
func SliceDedupe[T comparable](x []T) []T {
	if len(x) == 0 {
		return x
	}

	result := make([]T, 0)
	for _, item := range x {
		if !slices.Contains(result, item) {
			result = append(result, item)
		}
	}
	return result
}
