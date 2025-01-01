package main

import (
	"fmt"
	"os"
	"regexp"
	"Day3/match"
)

func main() {
	var do bool
	var dont bool
	data, err := os.ReadFile("input.csv")
	if err != nil {
		fmt.Println(err)
	}
	allowedChars := `[^mul(1234567890,)don't]`
	re := regexp.MustCompile(allowedChars)
	result := re.ReplaceAll(data, []byte(" "))
	total := 0

	for i:=0;i<len(result)-8;i++{
		if string(result[i]) == "d" &&string(result[i+1]) == "o" &&string(result[i+2]) == "n" &&string(result[i+3]) == "'" &&string(result[i+4]) == "t" &&string(result[i+5]) == "(" &&string(result[i+6]) == ")"{
			fmt.Println(string(result[i:i+7]))
			dont = true
			do = false
		}
		if string(result[i]) == "d" &&string(result[i+1]) == "o" &&string(result[i+2]) == "(" &&string(result[i+3]) == ")"{
			fmt.Println(string(result[i:i+4]))
			do = true
			dont = false
		}
		if do || (!dont && !do){
			total = total + match.Match(result,i)
		}
	}
fmt.Println(total)
}
