package main

import (
	"fmt"
	"os"
	"encoding/csv"
)

func extractLine(line []string, line1, line2 string) {
	for i:=0;i<11;i++ {
		if i == 0{
			line1 = line[i]
		}
		if i == 1{
			line2 = line[i]
		}
	}
	fmt.Println(line)
}

func main() {
	var line1 string
	var line2 string
	file, err := os.Open("input.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	line,_ := reader.Read()
	extractLine(line,line1,line2)
}
