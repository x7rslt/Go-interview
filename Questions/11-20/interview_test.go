package interview

import (
	"fmt"
	"reflect"
	"testing"
)

//Q11:[primary] 关于局部变量的初始化，下面正确的使用方式是（）  //注：这道题对应前面的全局变量初始化
//A. var i int = 10
//B. var i = 10
//C. i := 10
//D. i = 10
//参考答案：ABC

func TestArgs(t *testing.T) {
	var i int = 10
	var j = 10
	k := 10
	fmt.Println(i, j, k)
}

//Q12:[primary] 关于const常量定义，下面正确的使用方式是（）
//A.
//const Pi float64 = 3.14159265358979323846
//const zero = 0.0

//B.
//const (
//	size int64 = 1024
//	eof = -1
//)

//C.    //注：const initializer errors.New("element already exists") is not a constant
//const (
//	ERR_ELEM_EXIST error = errors.New("element already exists")
//	ERR_ELEM_NT_EXIST error = errors.New("element not exists")
//)

//D.
//const u, v float32 = 0, 3
//const a, b, c = 3, 4, "foo"
//参考答案：ABD

func TestConst(t *testing.T) {
	const Pi float64 = 3.14159265358979323846
	const zero = 0.0

	const (
		size int64 = 1024
		eof        = -1
	)

	//const (
	//ERR_ELEM_EXIST    error = errors.New("element already exists") //注：const initializer errors.New("element already exists") is not a constant
	//ERR_ELEM_NT_EXIST error = errors.New("element not exists")
	//)

	const u, v float32 = 0, 3
	const a, b, c = 3, 4, "foo"

}

//Q3:[primary] 关于布尔变量b的赋值，下面错误的用法是（）
//A. b = true
//B. b = 1  		#注：bool值只有true和false
//C. b = bool(1)
//D. b = (1 == 2)
//参考答案：BC

var b1, b2, b3, b4 bool

func TestBool(t *testing.T) {
	b1 = true
	//b2 = 1       //cannot use 1 (type untyped int) as type bool in assignment
	//b3 = bool(1) //cannot convert 1 (type untyped int) to type bool
	b4 = (1 == 2)
	fmt.Println(b1, b2, b3, b4) //b2,b3定义未赋值时，默认为false
}

//Q4:[intermediate] 下面的程序的运行结果是（）
//func main() {
//	if (true) {
//		defer fmt.Printf("1")
//	} else {
//		defer fmt.Printf("2")  //注：不符合条件，无打印
//	}
//	fmt.Printf("3")
//}
//A. 321
//B. 32
//C. 31
//D. 13
//
//参考答案：C
func TestDefer(t *testing.T) {
	if true {
		defer fmt.Printf("1")
	} else {
		defer fmt.Printf("2")
	}
	fmt.Printf("3")
}

//Q15:[primary] 关于switch语句，下面说法正确的有（）
//A. 条件表达式必须为常量或者整数
//B. 单个case中，可以出现多个结果选项
//C. 需要用break来明确退出一个case
//D. 只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case
//参考答案：BD

func TestSwithc(t *testing.T) {
	var x interface{}
	x = 10
	switch x.(type) {
	case int, uint:
		fmt.Println(x)
	case bool:
		fmt.Println(x)
	case string:
		fmt.Println(x)
	case nil:
		fmt.Println(x)
	default:
		fmt.Println("No value.")

	}

	fmt.Println(reflect.TypeOf(x))    //输出x类型
	fmt.Println(fmt.Sprintf("%T", x)) //输出x类型

	switch x {
	case 10:
		fmt.Println("This is a int :", x)
	case nil:
		fmt.Println(x)
	default:
		fmt.Println("No value.")

	}
}

//Q16:[intermediate] golang中没有隐藏的this指针，这句话的含义是（）
//A. 方法施加的对象显式传递，没有被隐藏起来
//B. golang沿袭了传统面向对象编程中的诸多概念，比如继承、虚函数和构造函数    //注：Golang没有继承，只有组合
//C. golang的面向对象表达更直观，对于面向过程只是换了一种语法形式来表达
//D. 方法施加的对象不需要非得是指针，也不用非得叫this
//参考答案：ACD

//Q17:[intermediate] golang中的引用类型包括（）
//A. 数组切片
//B. map
//C. channel
//D. interface
//参考答案：ABCD

//Q18[intermediate] golang中的指针运算包括（）
//A. 可以对指针进行自增或自减运算
//B. 可以通过“&”取指针的地址
//C. 可以通过“*”取指针指向的数据
//D. 可以对指针进行下标运算
//参考答案：BC
func TestPointer(t *testing.T) {
	x := 10
	p := &x
	for i := 1; i < 10; i++ {
		//p += i 			//Error：invalid operation: p += i (mismatched types *int and int)；注：只能值相加，指针不能
		*p += i
	}
	fmt.Println(*p)
}

//Q19:[primary] 关于main函数（可执行程序的执行起点），下面说法正确的是（）
//A. main函数不能带参数
//B. main函数不能定义返回值
//C. main函数所在的包必须为main包
//D. main函数中可以使用flag包来获取和解析命令行参数
//参考答案：ABCD

//Q20:[intermediate] 下面赋值正确的是（）
//A. var x = nil   					//注：不能用nil给一个没有类型的变量
//B. var x interface{} = nil
//C. var x string = nil
//D. var x error = nil
//参考答案：BD

func TestVar(t *testing.T) {
	//var x = nil //error:use of untyped nil
	var y interface{} = nil
	//var z string = nil //error:cannot use nil as type string in assignment
	var a error = nil
	fmt.Println(y, a)
}
