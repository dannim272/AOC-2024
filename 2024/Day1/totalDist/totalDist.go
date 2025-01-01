package totalDist

import (
	"sort"
)

func TotalDist(list1, list2 []int) int {
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	dist := 0
	var diff int

	for i := 0; i < len(list1);i++ {
		diff = list2[i] - list1[i]
		if diff < 0 {
			diff = diff * -1
		}
		dist = dist + diff
	}
	return dist
}
