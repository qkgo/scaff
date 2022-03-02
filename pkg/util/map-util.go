package util

import "sort"

type ValSorter struct {
	KeyList []string
	ValList []int
}

func NewValSorter(m map[string]int) *ValSorter {
	vs := &ValSorter{
		KeyList: make([]string, 0, len(m)),
		ValList: make([]int, 0, len(m)),
	}
	for k, v := range m {
		vs.KeyList = append(vs.KeyList, k)
		vs.ValList = append(vs.ValList, v)
	}
	return vs
}
func (vs *ValSorter) Sort() {
	sort.Sort(vs)
}
func (vs *ValSorter) Len() int           { return len(vs.ValList) }
func (vs *ValSorter) Less(i, j int) bool { return vs.ValList[i] < vs.ValList[j] }
func (vs *ValSorter) Swap(i, j int) {
	vs.ValList[i], vs.ValList[j] = vs.ValList[j], vs.ValList[i]
	vs.KeyList[i], vs.KeyList[j] = vs.KeyList[j], vs.KeyList[i]
}

//**
// INPUT [{},{}]   ATTR1  [1,2]
// OUT    [{ATTR1:1},{ATTR1:2}]
func MergeToMap(mergeMap []map[string]int, keyStr string, keyList []int) []map[string]int {
	for i := 0; i < len(mergeMap); i++ {
		if mergeMap[i] == nil {
			mergeMap[i] = map[string]int{}
		}
		if mergeMap[i] != nil {
			mergeMap[i][keyStr] = keyList[i]
		}
	}
	return mergeMap
}

// INPUT [{a:*},{b:*}]    [{ATTR1:1},{ATTR1:2}]
// OUTPUT [{a:{ATTR1:1}},{b:{ATTR1:2}}]
func MergeMapToKey(jobMap map[string]interface{}, amalgamate []map[string]int) map[string]interface{} {
	num := 0
	for key, _ := range jobMap {
		jobMap[key] = amalgamate[num]
		num++
	}
	return jobMap
}
