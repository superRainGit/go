package main
import "fmt"

// 同时声明多个全局的变量
var b1 = 100
var b2 = 200
var name = "jerry"
// 设计者认为上述的声明太麻烦 所以使用了如下的方式进行替换
var (
	b3 = 400
	b4 = 600
	name2 = "jack"
)

func main() {
	// 可以一次声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1 = ", n1, "n2 = ", n2, "n3 = ", n3)
	// 一次定义多个变量的方式2
	var m1, m2, m3 = 100, "tom", 88
	fmt.Println("m1 = ", m1, "m2 = ", m2, "m3 = ", m3)
	// 一次性定义多个变量的方式3 同样 类似于单个使用变量 同样可以直接使用:= 进行变量的声明+赋值
	o1, o2, o3 := 200, "哈哈", 300
	fmt.Println("o1 = ", o1, "o2 = ", o2, "o3 = ", o3)
	// 打印全局变量
	fmt.Println("b1 = ", b1, "b2 = ", b2, "name = ", name)
	fmt.Println("b3 = ", b3, "b4 = ", b4, "name2 = ", name2)
}