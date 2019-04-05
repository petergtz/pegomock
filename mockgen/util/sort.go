package util

import "sort"

// SortedKeys returns the keys of a map in alphabetical order.
func SortedKeys(m map[string]bool) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
