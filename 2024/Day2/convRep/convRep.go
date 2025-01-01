package convRep

import (
	"strings"
	"strconv"
)

func ConvRep(data [][]string,i int) []int {
	report := []int{}
	for a:=0;a<len(data[i]);a++ {
		nums := strings.Split(data[i][a]," ")
		for _, num := range nums {
			number, _ := strconv.Atoi(num)
			report = append(report,number)
		}
	}
	return report
}
