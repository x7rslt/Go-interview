package interview

import (
	"encoding/json"
	"fmt"
	"testing"
)

//71 [intermediate] 内置函数delete可以删除数组切片内的元素（）
//参考答案：错的   需要用切片拼接重新赋值的方式

//Slice删除指定值
func Remove(s []int, i int) []int {
	for j, k := range s {
		if k == i {
			return append(s[:j], s[j+1:]...)
		}
		fmt.Printf("Slice no value %d.", i)
	}
	return s
}

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s = Remove(s, 1)
	fmt.Println(s)
}

//72 [primary] 指针是基础类型（）
//参考答案：错的
//基础类型有：
//整数类型： int8、int16等等
//浮点类型：float32、float64
//布尔类型：bool
//复数类型： complex64、complex128
//字符串类型： string
//字符类型： byte、rune
//指针是复杂类型
func TestPointer(t *testing.T) {
	i := 1
	p := &i
	*p = 1
	fmt.Printf("Type: p is %T,*p is %T,&p is %T", p, *p, &p)
}

//73 [primary] interface{}是可以指向任意对象的Any类型（）
//参考答案：对的

//74 [intermediate] 下面关于文件操作的代码可能触发异常（）
//file, err := os.Open("test.go")
//defer file.Close()
//if err != nil {
//fmt.Println("open file failed:", err)
//return
//}
//...
//参考答案：对的

// Open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.

//75 [primary] Golang不支持自动垃圾回收（）
//参考答案：错的，支持的

//76 [primary] Golang支持反射，反射最常见的使用场景是做对象的序列化（）
//参考答案：对的

//76 [primary] Golang支持反射，反射最常见的使用场景是做对象的序列化（）  //不理解
//参考答案：对的

//77 [primary] Golang可以复用C/C++的模块，这个功能叫Cgo（）
//参考答案：工作中使用go ,没有服用过c/c++模块，这个不是很确定，有知道的小伙伴可以打在弹幕上

//77 [primary] Golang可以复用C/C++的模块，这个功能叫Cgo（）
//参考答案：对

//78 [primary] 下面代码中两个斜点之间的代码，比如json:"x"，作用是X字段在从结构体实例编码到JSON数据格式的时候，
//使用x作为名字，这可以看作是一种重命名的方式（）
//type Position struct {
//	X int `json:"x"`
//	Y int `json:"y"`
//	Z int `json:"z"`
//}
//参考答案：对的， 输出的json{"x":1,"y":2,"z":3}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func TestJson(t *testing.T) {
	p := []Position{{1, 2}, {2, 4}}
	j, _ := json.Marshal(p)
	fmt.Printf("%s", j)
}

//79 [primary] 通过成员变量或函数首字母的大小写来决定其作用域（）
//参考答案：对的

//80 [primary] 对于常量定义zero(const zero = 0.0)，zero是浮点型常量（）
//参考答案：错的， 是float64

func TestFloat(t *testing.T) {
	const zero = 0.0
	fmt.Printf("%T", zero)
}
