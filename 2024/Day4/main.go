package main

import (
	"fmt"
	"os"
	"bufio"
)

func hor(lines string,i,total int) int  {
	if string(lines[i:i+4]) == "XMAS"{
		total++
	}
	if string(lines[i:i+4]) == "SAMX"{
		total++
	}
	return total
}

func ver(table []string,count,i,total int) int {
	col := []string{}
	xmas := []string{"X","M","A","S"}
	for a:=0;a<count;a++ {
		col = append(col,string(table[a][i]))
	}
	for a:=0;a<count;a++{
		if string(col[a:a+4]) == string(xmas) {
			total++
		}
	}
	return total
}

func diag(table []string,count,i int)  {
	col := []string{}
	for a:=0;a<count;a++ {
		col = append(col,string(table[i][a]))
	}
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
	for i:=0;i<count;i++{
		total = total + ver(table,count,i,total)
		diag(table,count,i)
	}
	fmt.Println(total)
}
