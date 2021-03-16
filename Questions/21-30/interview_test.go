package interview

import (
	"fmt"
	"testing"
)

//Q21:[intermediate] 关于整型切片的初始化，下面正确的是（）
//A. s := make([]int)     			//make([]T,len,cap) ,未定义长度
//B. s := make([]int, 0)
//C. s := make([]int, 5, 10)
//D. s := []int{1, 2, 3, 4, 5}
//参考答案：BCD
func TestSlice(t *testing.T) {
	// s := make([]int)				//Error:missing len argument to make([]int)
	s2 := make([]int, 0)
	s3 := make([]int, 8, 10) //make([]T,len,cap)
	s4 := []int{1, 2, 3, 4, 5}
	fmt.Println(s2, s3, s4)
}

//Q22：[intermediate] 从切片中删除一个元素，下面的算法实现正确的是（）
//A.
//func (s *Slice)Remove(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			if i== len(*s) - 1 {
//				*s = (*s)[:i]
//			}else {
//				*s = append((*s)[:i],(*s)[i + 2:]...)
//			}
//			return nil
//		}
//	}
//	return ERR_ELEM_NT_EXIST
//}

//B.
//func (s *Slice)Remove(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			*s = append((*s)[:i],(*s)[i + 1:])
//			return nil
//		}
//	}
//	return ERR_ELEM_NT_EXIST
//}

//C.
//func (s *Slice)Remove(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			delete(*s, v)
//			return nil
//		}
//	}
//	return ERR_ELEM_NT_EXIST
//}

//D.
//func (s *Slice)Remove(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			*s = append((*s)[:i],(*s)[i + 1:]...)     //整个逻辑就是，找到你要删除的元素，然后把它前面的和后面的组合起来
//			return nil
//		}
//	}
//	return ERR_ELEM_NT_EXIST
//}
//参考答案：D

//Q23:[primary] 对于局部变量整型切片x的赋值，下面定义正确的是（）
//A.
/*
x := []int{
1, 2, 3,
4, 5, 6,
}
//B.
//
x := []int{
1, 2, 3,
4, 5, 6
}
//C.
//
x := []int{
1, 2, 3,
4, 5, 6}
//D.
x := []int{1, 2, 3, 4, 5, 6,}
*/
//参考答案：ACD
func TestSliceDefine(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6}
	x2 := []int{1, 2, 3, 4, 5, 6}

	/*x3 := []int{
		1, 2, 3, 4,
		5, 6, 7, 8  			//Error:missing ',' before newline in composite literal
	}
	*/
	x4 := []int{
		1, 2, 3, 4,
		5, 6, 7, 8,
	}
	x5 := []int{
		1, 2, 3, 4,
		5, 6, 7, 8}
	x6 := []int{
		1, 2, 3, 4,
		5, 6, 7, 8}
	fmt.Println(x, x2, x4, x5, x6)
}

//Q24:[primary] 关于变量的自增和自减操作，下面语句正确的是（）
//A.
//i := 1
//i++
//B.
//
//i := 1
//j = i++     //j都没有初始化变量，什么鬼赋值给j
//C.
//
//i := 1
//++i         //鬼畜必错
//D.
//
//i := 1
//i--
//参考答案：AD

func TestAC(t *testing.T) {
	i := 1
	//++i  			//ERROR:expected statement, found '++'
	i += i
	fmt.Println(i)
}

//Q25:[intermediate] 关于函数声明，下面语法错误的是（）
//A. func f(a, b int) (value int, err error)
//B. func f(a int, b int) (value int, err error)
//C. func f(a, b int) (value int, error)
//D. func f(a int, b int) (int, int, error)
//参考答案：C

func TestFuncDefine(t *testing.T) {
	//用匿名函数测试
	func(a int, b int) (value int, err error) {
		return a + b, nil
	}(1, 2) //OK

	/*
		func(a, b int) (value int,error) { //Error:expected expression
			return a + b, nil
		}(1, 2)
	*/
}

//Q26:[intermediate] 如果Add函数的调用代码为：
//func main() {
//	var a Integer = 1
//	var b Integer = 2
//	var i interface{} = &a
//	sum := i.(*Integer).Add(b)
//	fmt.Println(sum)
//}
//则Add函数定义正确的是（）
//A.
//
//type Integer int
//func (a Integer) Add(b Integer) Integer {
//	return a + b
//}
//B.
//type Integer int
//func (a Integer) Add(b *Integer) Integer {
//	return a + *b
//}
//C.
//
//type Integer int
//func (a *Integer) Add(b Integer) Integer {
//	return *a + b
//}

//D.
//type Integer int
//func (a *Integer) Add(b *Integer) Integer {
//	return *a + *b
//}
//参考答案：AC

//B和D，都是要接受Integer指针类型的参数，所以不对

//如果将一个接口类型变量断言成一个指针类型的变量，在断言成功的前提下，两个变量将共享内存空间,
//而且add方法，可以是struct和它的指针类型调用，所以AC是可以的
type Integer int

func (a *Integer) Add(b Integer) Integer {
	return *a + b
}

func (a Integer) Add2(b Integer) Integer {
	return a + b
}
func TestFuncPointer(t *testing.T) {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b)
	sum2 := i.(*Integer).Add(b)
	fmt.Println(sum, sum2)
}

//Q27:[intermediate] 如果Add函数的调用代码为：
//func main() {
//	var a Integer = 1
//	var b Integer = 2
//	var i interface{} = a
//	sum := i.(Integer).Add(b)
//	fmt.Println(sum)
//}
//则Add函数定义正确的是（）
//A.
//type Integer int
//func (a Integer) Add(b Integer) Integer {
//	return a + b
//}
//B.
//type Integer int
//func (a Integer) Add(b *Integer) Integer {
//	return a + *b
//}
//C.
//type Integer int
//func (a *Integer) Add(b Integer) Integer {
//	return *a + b
//}
//D.
//
//type Integer int
//func (a *Integer) Add(b *Integer) Integer {
//	return *a + *b
//}
//参考答案：A

//A 看定义是可以的，很顺眼，
//C,定义是指针类型，不能调用
//BD参数都是指针，直接完犊子了

//Q28:[intermediate] 关于GetPodAction定义，下面赋值正确的是（）
type TransInfo string

type Fragment interface {
	Exec(transInfo *TransInfo) error
}
type GetPodAction struct {
}

func (g GetPodAction) Exec(transInfo *TransInfo) error {
	fmt.Println(g, transInfo)
	return nil
}

//A. var fragment Fragment = new(GetPodAction) 			//A new可以，因为它返回指针类型
//B. var fragment Fragment = GetPodAction				//B 结构体不能作赋值
//C. var fragment Fragment = &GetPodAction{}			//C 很明显了，人家都用&符号了
//D. var fragment Fragment = GetPodAction{}				//D 一个结构体对象赋值给接口变量 ok 啊
//
//参考答案：ACD
func TestInterfaceDefine(t *testing.T) {
	var fragment Fragment = &GetPodAction{}
	fmt.Printf("%T", fragment)
	//var fragment2 Fragment = *GetPodAction{}  			//invalid indirect of GetPodAction literal (type GetPodAction)
	fmt.Println(fragment)
}

//补充知识：接口重定义
type Stringer interface {
	String() string
}
type S string

func (s S) String() string {
	return "hahaha"
}
func TestInterface(t *testing.T) {
	s := S("lehehe")
	fmt.Println(s)
}

//Q29:[intermediate] 关于GoMock，下面说法正确的是（）//注：打桩知识点无知，待补习
//A. GoMock可以对interface打桩
//B. GoMock可以对类的成员函数打桩
//C. GoMock可以对函数打桩
//D. GoMock打桩后的依赖注入可以通过GoStub完成

//参考答案：AD

//Q30:[intermediate] 关于接口，下面说法正确的是（）
//A. 只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等价的，可以相互赋值
//B. 如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A
//C. 接口查询是否成功，要在运行期才能够确定
//D. 接口赋值是否可行，要在运行期才能够确定
//参考答案：ACD

//A儿子  B爹   父给子转换

type Father interface {
	DoWork()
}

type Son interface {
	Father
	Happy()
}

type Human struct {
	Name string
}

func (h Human) DoWork() {
	fmt.Println(h.Name + " 工作...")
}

type Person struct {
	Name string
}

func (p Person) DoWork() {
	fmt.Println(p.Name + "是富二代，无需工作。")
}

func (p Person) Happy() {
	fmt.Println(p.Name + "在玩")
}

func TestInterfaceTransfer(t *testing.T) {
	var father Father
	var son Son
	father = Human{Name: "老王"}

	//son = father     					//Error:cannot use father (type Father) as type Son in assignment Father does not implement Son (missing Happy method)

	fmt.Println(son, father)
}
