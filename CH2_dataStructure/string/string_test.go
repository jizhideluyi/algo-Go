package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//字符串大小写转换
	fmt.Println(strings.ToUpper("peter"))

	formatPrint()
	trim()
	printExpression()
}

/*格式化输出*/
func formatPrint() {
	fmt.Println(strings.Join([]string{"hello", "world"}, " "))
	fmt.Println(fmt.Sprintf("%s:%s", "peter", "13000000000"))

	user3 := bytes.Buffer{}
	user3.WriteString("hello")
	user3.WriteString(" world")
	fmt.Println(user3.String())
}

/*字符串替换*/
func trim() {
	fmt.Println(strings.Trim("    去除空格    ", " "))
	fmt.Println(strings.Trim("米饭，馒头，面条", "米饭"))
}

//字符串拼接
func printExpression() {
	fmt.Println((1 + 2) * 3)
	fmt.Println("我的年龄是:" + strconv.Itoa(23))
}
