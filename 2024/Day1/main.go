package main

import (
	"Day1/createList"
	"Day1/totalDist"
	"fmt"
)

func main() {
	// 1st part
	list1, list2 := createList.CreateList()
	dist := totalDist.TotalDist(list1,list2)
	fmt.Println(dist)

	// 2nd part
	var count int
	sim := 0
	for i:=0;i<len(list1);i++{
		for a:=0;a<len(list2);a++{
			if list1[i] == list2[a] {
				count++
			}
		}
		sim = sim + (list1[i]*count)
		count = 0
	}
	fmt.Println(sim)
}
