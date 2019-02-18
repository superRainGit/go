package main
import "fmt"

func main() {
	// 若声明了变量但是没有进行赋值 那么使用默认值
	// int的默认值是 0 
	var i int
	fmt.Println("i = ", i) // 这句话会输出 i = 0

	// 根据变量的值进行类型推断
	var num = 10.11
	fmt.Println("num = ", num) // 这句话会输出 num = 10.11

	// 省略var 此时不能使用= 而是使用:= 进行变量的书写
	// 注意：:= 左侧的变量不能是已经声明过的变量 否则会编译报错
	// 使用这种写法 等价 -> var name string     name = "tom" 
	name := "tom"
	fmt.Println("name = ", name)
	
	// 如果使用如下的方式书写 在编译时会有如下的提示信息
	// # command-line-arguments
 	// ./main2.go:20:7: no new variables on left side of :=
	// name := "123"
	// fmt.Println("name = ", name)
}