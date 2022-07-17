package main

import "fmt"

//全局变量
var msg = "去外面吃点好的。"

/**
作用域：公开、模块、包、局部

小写：包
大写：公开
*/
func main() {
	/*
		局部变量
	*/
	msg := "main函数内"

	{
		msg := "在家吃点好的"
		fmt.Println(msg)
	}

	fmt.Println(msg)
}
