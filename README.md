# goTemplateTools - a set of tools using go 1.18 templates 


Remove from an unordered slice at location `pos`.

```
func RemoveFromSlice[T any](s []T, pos int) []T {
```

Remove from an ordered slice at location `pos` maintaining order.

```
func RemoveFromOrderedSlice[T any](s []T, pos int) []T {
```
