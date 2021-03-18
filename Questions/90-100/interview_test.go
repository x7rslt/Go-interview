package interview

import (
	"fmt"
	"testing"
)

//91 [primary] 同级文件的包名不允许有多个（）
//参考答案：对的
//
//92 [intermediate] 可以给任意类型添加相应的方法（）
//参考答案：错的
//
//93 [primary] golang虽然没有显式的提供继承语法，但是通过匿名组合实现了继承（）
//参考答案：对的

//94 [primary] 使用for range迭代map时每次迭代的顺序可能不一样，因为map的迭代是随机的（）
//参考答案：对的
//
//95 [primary] switch后面可以不跟表达式（）
//参考答案：对的

//96 [intermediate] 结构体在序列化时非导出变量（以小写字母开头的变量名）不会被encode，
//因此在decode时这些非导出变量的值为其类型的零值（）
//参考答案：对的

//97 [primary] golang中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以NewXXX来命名（）
//参考答案：对的

//98 [intermediate] 当函数deferDemo返回失败时，并不能destroy已create成功的资源（）
//func deferDemo() error {
//	err := createResource1()
//	if err != nil {
//		return ERR_CREATE_RESOURCE1_FAILED
//	}
//	defer func() {
//		if err != nil {
//			destroyResource1()
//		}
//	}()
//
//	err = createResource2()
//	if err != nil {
//		return ERR_CREATE_RESOURCE2_FAILED
//	}
//	defer func() {
//		if err != nil {
//			destroyResource2()
//		}
//	}()
//
//	err = createResource3()
//	if err != nil {
//		return ERR_CREATE_RESOURCE3_FAILED
//	}
//	return nil
//}
//参考答案：错的，有defer啊
//
//99 [intermediate] channel本身必然是同时支持读写的，所以不存在单向channel（）
//参考答案：错的
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func TestSingleChannel(t *testing.T) {
	n := make(chan int)
	s := make(chan int)
	go counter(n)
	go squarer(s, n)
	printer(s)
}

//
//100 [primary] import后面的最后一个元素是包名（）   //元素的话是路径，最后一个字段是包名
//参考答案：错的  import后面跟的是包的路径，而不是包名。
