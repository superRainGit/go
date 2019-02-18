package main
import "fmt"

func main() {
	// 该区域的数据值可以在同一类型范围内变化
	var i int = 10
	i = 30
	i = 50
	fmt.Println("i = ", i)
	// 如下写法会提示错误
	// constant 1.2 truncated to integer
	// i = 1.2 

	// 同其他语言一样 变量在同一个作用域(在一个函数或者在代码块)内不能重名
	// 同样的 也不能使用:= 进行同一个变量的声明+赋值
	// i redeclared in this block
	var i int = 59
}