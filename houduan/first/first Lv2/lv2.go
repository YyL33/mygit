package main

import "fmt"

func calculator1() {
	var a int
	var b int
	var s string
	fmt.Println("请输入计算式：")
	fmt.Scan(&a, &s, &b) //格式：a回车s回车b回车，或把回车替换为空格
	if s == "+" {
		fmt.Println(a, s, b, "=", a+b)
	} else if s == "-" {
		fmt.Println(a, s, b, "=", a-b)
	} else if s == "*" {
		fmt.Println(a, s, b, "=", a*b)
	} else if s == "/" {
		fmt.Println(a, s, b, "=", a/b)
	}
}
func calculator2() {
	var a int
	var b int
	var s string
	fmt.Println("请输入计算式:")
	fmt.Scan(&a,&s,&b)      //格式：a回车s回车b回车，或把回车替换为空格
	switch s{
	case "+":fmt.Println(a,s,b,"=",a + b)
	case "-":fmt.Println(a,s,b,"=",a - b)
	case "*":fmt.Println(a,s,b,"=",a * b)
	case "/":fmt.Println(a,s,b,"=",a / b)
	default:
		fmt.Println("错误，请重试")
	}
}

func main() {
	for {
		calculator1()
	}
}
