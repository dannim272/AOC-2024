package match

import (
	"fmt"
	"strconv"
)

func Match(result []byte,i int) int{
	var num string
	var num1 string
	var mult int = 0
	if string(result[i]) == "m" && string(result[i+1]) == "u" && string(result[i+2]) == "l" && string(result[i+3]) == "(" {
		if string(result[i+5]) == "," {
			num = string(result[i+4])
			if string(result[i+7]) == ")" {
				num1 = string(result[i+6])
			} else if string(result[i+8]) == ")" {
				num1 = string(result[i+6]) + string(result[i+7])
			} else if string(result[i+9]) == ")" {
				num1 = string(result[i+6]) + string(result[i+7]) + string(result[i+8])
			}
		}
		if string(result[i+6]) == "," {
			num = string(result[i+4]) + string(result[i+5])
			if string(result[i+8]) == ")" {
				num1 = string(result[i+7])
			} else if string(result[i+9]) == ")" {
				num1 = string(result[i+7]) + string(result[i+8])
			} else if string(result[i+10]) == ")" {
				num1 = string(result[i+7]) + string(result[i+8]) + string(result[i+9])
			}
		}
		if string(result[i+7]) == "," {
			num = string(result[i+4]) + string(result[i+5]) + string(result[i+6])
			if string(result[i+9]) == ")" {
				num1 = string(result[i+8])
			} else if string(result[i+10]) == ")" {
				num1 = string(result[i+8]) + string(result[i+9])
			} else if string(result[i+11]) == ")" {
				num1 = string(result[i+8]) + string(result[i+9]) + string(result[i+10])
			}
		}
		num2, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
		}
		num3, err1 := strconv.Atoi(num1)
		if err1 != nil {
			fmt.Println(err)
		}
		mult = num2 * num3
		fmt.Println(num," - ",num1)
	}
	return mult
}
