package goTemplateTools

// RemoveFromSlice removes an element at `pos` from an un-ordered slice.
func RemoveFromSlice[T any](s []T, pos int) []T {
	s[pos] = s[len(s)-1]
	return s[:len(s)-1]
}

// RemoveFromSlice removes an element at `pos` from an ordered slice.
// This is slower than RemoveFromSlice() but maintains the order in
// the slice by copying the tail of the slice to the new slice.
func RemoveFromOrderedSlice[T any](s []T, pos int) []T {
	return append(s[:pos], s[pos+1:]...)
}
