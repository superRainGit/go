package main
import "fmt"

func main() {

	// 测试一下int8的范围
	// 其他的int64 int32 int16类似
	var i int8 = 127
	// 如果使用如下的方式 那么会提示下面的错误码
	// constant 128 overflows int8
	// var i int8 = 128
	fmt.Println("i = ", i)

	// 测试一下uint8的范围
	// 其他的uint16 uint32 uint64类似
	var j uint8 = 255
	// 同样的 如果使用溢出 在运行时也会提示溢出
	// 如果使用如下的方式 那么会提示下面的错误码
	// constant 256 overflows uint8
	// var j uint8 = 256
	fmt.Println("j = ", j)
}