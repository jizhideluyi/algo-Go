package main

import "fmt"

//全局变量
var msg = "去外面吃点好的。"

func main() {
	//局部变量
	msg := "main函数内"
	{
		msg := "在家吃点好的"
		fmt.Println(msg)
	}

	fmt.Println(msg)
}
