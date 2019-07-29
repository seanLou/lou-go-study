package main

import "fmt"

func main() {
	a := 10

	// if-else
	if a < 20 {
		fmt.Printf("a小于20\n")
	} else {
		fmt.Printf("a大于20\n")
	}
	fmt.Printf("a的值是：%d\n", a)

	// else-if
	if a > 20 {
		fmt.Printf("a大于20\n")
	} else if a > 10 {
		fmt.Printf("a大于10, 小于等于20\n")
	} else {
		fmt.Printf("a小于等于10\n")
	}

	// 初始化子句, 变量b只在if代码快中有效
	if b := 15; b < 20 {
		fmt.Printf("b小于20\n")
	} else {
		fmt.Printf("b大于20\n")
	}

	// switch 用法 1. 表达式 switch

	marks := 87
	grade := "B"

	// 单值判断
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	fmt.Printf("你的评分是%s\n", grade)

	marks = 97

	// 范围判断
	switch {
	case marks >= 90:
		grade = "A"
	case marks >= 80:
		grade = "B"
	case marks >= 70:
		grade = "C"
	case marks >= 60:
		grade = "D"
	default:
		grade = "E"
	}

	switch {
	case grade == "A":
		fmt.Printf("成绩优秀")
	case grade == "B":
		fmt.Printf("B")
	// 多表达是写法
	//case grade == "C", grade == "D":
	//	fmt.Printf("成绩一般继续努力")
	case grade == "C":
		// fallthrough 可以把当期case的控制权交给下一个case判断
		// fallthrough 跟多表达式 效果一样
		fallthrough
	case grade == "D":
		fmt.Printf("成绩一般继续努力")
	default:
		fmt.Printf("不及格")
	}

	fmt.Printf("你的成绩是%s\n", grade)

	// 2. 类型 switch

}
