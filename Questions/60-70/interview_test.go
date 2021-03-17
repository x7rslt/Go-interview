package interview

import (
	"fmt"
	"testing"
)

//61 [intermediate] 下面的程序的运行结果是__________
//func main() {
//	x := 1
//	{
//		x := 2
//		fmt.Print(x)
//	}
//	fmt.Println(x)
//}
//参考答案：21   就近原则，谁离的近，就找谁，和找对象一样，你工作和哪个小姐姐，小哥哥经常接触，就最有机会
//
//62 [intermediate] 下面的程序的运行结果是__________
//func main() {
//	strs := []string{"one", "two", "three"}
//
//	for _, s := range strs {
//		go func() {
//			time.Sleep(1 * time.Second)
//			fmt.Printf("%s ", s)
//		}()
//	}
//	time.Sleep(3 * time.Second)
//}
//参考答案：three three three   闭包问题，当函数执行的时候，s的值都是three;go func 并不能保证立即执行，在它还没有执行的时候，item值已经被for循环改变了。
//
//63 [intermediate] 下面的程序的运行结果是__________
func TestRun63(t *testing.T) {
	x := []string{"a", "b", "c"}
	for v, i := range x {
		fmt.Println(v, i)
	}
}

//参考答案：
//0 a
//1 b
//2 c
//
//64 [intermediate] 下面的程序的运行结果是__________
//func main() {
//	x := []string{"a", "b", "c"}
//	for _, v := range x {
//		fmt.Print(v)
//	}
//}
//参考答案：abc
//
//65 [primary] 下面的程序的运行结果是__________
//func main() {
//	i := 1
//	j := 2
//	i, j = j, i
//	fmt.Printf("%d%d\n", i, j)
//}
//参考答案：21
//
//66 [primary] 下面的程序的运行结果是__________
//func incr(p *int) int {
//	*p++
//	return *p
//}
//func main() {
//	v := 1
//	incr(&v)
//	fmt.Println(v)
//}
//参考答案：2
//
//67 [primary] 启动一个goroutine的关键字是__________
//参考答案：go
//
//68 [intermediate] 下面的程序的运行结果是__________
type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	//fmt.Println(*s, elem)
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
func TestSlice(t *testing.T) {
	s := NewSlice()
	defer s.Add(1).Add(2)
	//fmt.Println(s)
	s.Add(3)
}

//参考答案：132     //不理解

//69 [primary] 数组是一个值类型（）
//参考答案：对的
//
//70 [primary] 使用map不需要引入任何库（）
//参考答案：对的
