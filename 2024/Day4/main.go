package main

import (
	"fmt"
	"os"
	"bufio"
)

func hor(lines string,i,total int) int  {
	if string(lines[i:i+4]) == "XMAS"{
		fmt.Println(string(lines[i:i+4]))
		total++
	}
	if string(lines[i:i+4]) == "SAMX"{
		fmt.Println(string(lines[i:i+4]))
		total++
	}
	return total
}

func ver(table []string,count,total int) int {
	xmas := 0
	for i:=0;i<count;i++ {
		if string(table[i][9]) == "X" {
			xmas++
		} else if string(table[i][9]) == "M" && xmas == 1{
			xmas++
		} else if string(table[i][9]) == "A" && xmas == 2{
			xmas++
		} else if string(table[i][9]) == "S" && xmas == 3{
			xmas++
		}
	}
	if xmas == 4 {
		total++
	}
	return total
}

func main() {
	total := 0
	count := 0
	table := []string{}
	file, _ := os.Open("prac.csv")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		for i:=0;i<len(lines)-3;i++{
			total = hor(lines,i,total)
		}
		table = append(table,lines) 
		count++
	}
	total = total + ver(table,count,total)
	fmt.Println(total)
}
