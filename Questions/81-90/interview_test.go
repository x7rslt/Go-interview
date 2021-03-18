package interview

import (
	"fmt"
	"testing"
	"time"
)

//81 [primary] 对变量x的取反操作是~x（）
//参考答案：错的 ， !

//82 [primary] 下面的程序的运行结果是xello（）
//func main() {
//	str := "hello"
//	str[0] = 'x'
//	fmt.Println(str)
//}
//参考答案：错的， str[0] = 'x' 这句话编译都不能通过，Go字符串是不可变的
func TestString(t *testing.T) {
	s := "hello"
	//s[0] = "x"
	fmt.Println(s)
}

//83 [primary] golang支持goto语句（）
//参考答案：对的

//84 [primary] 下面代码中的指针p为野指针，因为返回的栈内存在函数结束时会被释放（）
//type TimesMatcher struct {
//	base int
//}
//func NewTimesMatcher(base int) *TimesMatcher{
//	return &TimesMatcher{base:base}
//}
//func main() {
//	p := NewTimesMatcher(3)
//	...
//}
//参考答案：错的，NewTimesMatcher函数后不会被释放

//85 [primary] 匿名函数可以直接赋值给一个变量或者直接执行（）
//参考答案：对的
//
//86 [primary] 如果调用方调用了一个具有多返回值的方法，但是却不想关心其中的某个返回值，
//可以简单地用一个下划线“_”来跳过这个返回值，该下划线对应的变量叫匿名变量（）
//参考答案：对的

//87 [primary] 在函数的多返回值中，如果有error或bool类型，则一般放在最后一个（）
//参考答案：对的
//
//88 [primary] 错误是业务过程的一部分，而异常不是（）
//参考答案：对的
func TestPanic(t *testing.T) {
	defer fmt.Println("Test")
	for i := 0; i < 100; i++ {
		if i == 10 {
			panic("defer can't run.")
		}
		fmt.Println("1")
		time.Sleep(time.Second)
	}
}
