package util
import "sort"
type KeyValueSort struct {
	Key   string
	Value int
}
func SortMap(m map[string]int) []KeyValueSort {
	var ss []KeyValueSort
	for k, v := range m {
		ss = append(ss, KeyValueSort{k, v})
	}
	sort.Slice(ss, func(i int, j int) bool {
		return ss[i].Value > ss[j].Value 
	})
	return ss
}
