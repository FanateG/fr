package main

import (
	"fmt"
	"strconv"
)

var arabian = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}
var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var value1, value2, operator string

// var v1 string = "1"
// var v2 string = "0"
func calculate(value1, value2, ope string) string {
	var result int
	v1, _ := strconv.Atoi(value1)
	v2, _ := strconv.Atoi(value2)
	if v1 <= 10 && v2 <= 10 {

		switch ope {
		case "+":
			result = v1 + v2
		case "-":
			result = v1 - v2
		case "*":
			result = v1 * v2
		case "/":
			result = v1 / v2
		default:
			panic("wrong expression")
		}

	} else {
		panic("fail > 10")
	}
	return strconv.Itoa(result)
}

func converttoArab(x string) string {
	var str string
	m := roman[x]
	if m <= 10 {
		str = strconv.Itoa(m)
	} else {
		panic("fail>10")
	}
	return str
}
func converttoRome(x string) string {
	var str string
	m, _ := strconv.Atoi(x)

	if m > 0 {
		for _, s := range arabian {
			for j := s; j <= m; {
				for st, in := range roman {
					if in == s {
						str += st
						m -= s
					}
				}
			}
		}

	} else {
		panic("в римских нет отрицательных и нулей")
	}
	return str
}
func checkRome(x string) bool {
	var flag bool
	for i := range roman {
		if x == i {
			flag = true
		}
	}
	return flag
}
func checkArab(x string) bool {
	var flag bool
	m, _ := strconv.Atoi(x)
	for _, i := range roman {
		if i == m {
			flag = true
		}
	}
	return flag
}
func calculator(x, y string) string {
	var str string
	var v1 string
	var v2 string
	if len(x) != 0 || len(y) != 0 || len(operator) != 0 {

		if checkRome(x) && checkRome(y) {
			v1 = converttoArab(x)
			v2 = converttoArab(y)
			str = converttoRome(calculate(v1, v2, operator))
		} else {
			if checkArab(x) && checkArab(y) {
				str = calculate(x, y, operator)
			} else {
				panic("fail")
			}
		}

	} else {
		panic("failed")
	}
	return str
}
func main() {
	fmt.Scan(&value1, &operator, &value2)
	fmt.Println(calculator(value1, value2))

}
