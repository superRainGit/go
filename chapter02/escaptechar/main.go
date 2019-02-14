package main

import "fmt"

func main() {
	// 测试转义字符的使用
	fmt.Println("tom\tjack")
	fmt.Println("hello\nworld")
	fmt.Println("Tom said \"I love U\"")
	// 回车就是从当前行的最前面开始输出 覆盖掉以前的内容
	fmt.Println("天龙八部雪山飞狐\r张飞")
}