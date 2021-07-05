package util

// RemoveListItem removes an Item from a list without changing the order
func RemoveListItem(list []string, key string) ([]string, bool) {
	i := 0
	for ; i < len(list); i++ {
		if list[i] == key {
			return append(list[:i], list[i+1:]...), false
		}
	}
	return list, true
}
