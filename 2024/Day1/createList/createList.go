package createList

import (
	"fmt"
	"os"
	"encoding/csv"
	"strings"
	"strconv"
)

func CreateList() ([]int, []int) {
	file, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("something wrong")
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	var count int = 1
	list1 := []int{}
	list2 := []int{}

	for _, inner := range records {
		for _, str := range inner {
			numbers := strings.Split(str, " ")
			for _, number := range numbers {
				if count % 2 != 0 {
					num, _ := strconv.Atoi(number)
					if num != 0 {
						list1 = append(list1,num)
					}
				}else {
					num, _ := strconv.Atoi(number)
					if num != 0 {
						list2 = append(list2,num)
					}
				}
				count++
			}
		}
	}
	return list1, list2
}
