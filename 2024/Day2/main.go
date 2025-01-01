package main

import (
	"os"
	"encoding/csv"
	"Day2/convRep"
	"fmt"
)

type Report struct {
	Value []int 
	Safe bool
	Increase bool
	Dumpener bool
}

func increase(report Report, i, diff, step int) (int,Report){
	if report.Value[i] < report.Value[i+1] && (diff == 1 || diff == 2 || diff == 3){
		step++
		report.Increase = true
	}
	return step, report
}

func decrease(report Report, i, diff, step int) (int,Report){
	if report.Value[i] > report.Value[i+1] && (diff == -1 || diff == -2 || diff == -3){
		step++
		report.Increase = false
	}
	return step, report
}

func same(report Report, i int) (Report){
		if report.Value[i] == report.Value[i+1] {
			report.Value = append(report.Value[:i], report.Value[i+1:]...)
			report.Dumpener = true
	}
	return report
}

func changeList(report Report, i int) Report {
	report.Value = append(report.Value[:i], report.Value[i+1:]...)
	return report
}

func check(report Report, count int) int { 
	step := 0
	firstGo := false 
	dumpCount := 0
	for i:=0;i<len(report.Value)-1;i++{
		diff := report.Value[i+1] - report.Value[i] 
		if !report.Dumpener{
			if report.Dumpener == true {
				count++
				report.Safe = false
			}
			if dumpCount == 0 {
				same(report,i)
			}
		}
		same(report,i)
		if !firstGo {
			step, report = increase(report,i,diff,step)
			step, report = decrease(report,i,diff,step)
			firstGo = true
		}
		if report.Increase {
			if report.Value[i] < report.Value[i+1] && (diff == 1 || diff == 2 || diff == 3){
				step++
			}
			if !report.Dumpener && (diff == -1 || diff == -2 || diff == -3){
				report = changeList(report,i)
				report.Dumpener = true
				if report.Value[i-1] < report.Value[i] && (diff == 1 || diff == 2 || diff == 3){
					step++
				}
			}
		}
		if !report.Increase {
			if report.Value[i] > report.Value[i+1] && (diff == -1 || diff == -2 || diff == -3){
				step++
			}
			if !report.Dumpener && (diff == 1 || diff == 2 || diff == 3){
				report = changeList(report,i)
				report.Dumpener = true
				if report.Value[i-1] < report.Value[i] && (diff == -1 || diff == -2 || diff == -3){
					step++
				}
			}
		}
		if step == len(report.Value) {
			count++ 
		}
	}
	return count 
}

func main() {
	count := 0
	file, _ := os.Open("input.csv")
	reader := csv.NewReader(file)
	data, _ := reader.ReadAll()

	var report Report

	for i:=0; i<len(data);i++{
		report.Value = convRep.ConvRep(data,i)
		count = check(report,count)
	}
	fmt.Println(count)
}
