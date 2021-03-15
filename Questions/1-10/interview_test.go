package basic_test

import (
	"fmt"
	"testing"
)

/*
Q1:[primary] 下面属于关键字的是（）
A. func
B. def  		python的
C. struct
D. class 		java的

参考答案：AC
*/

//Q2：[primary] 定义一个包内全局字符串变量，下面语法正确的是 （）
//A. var str string
//B. str := ""
//C. str = ""
//D. var str = ""
//参考答案 AD

var str string
var str2 = ""

func TestVar(t *testing.T) {
	str3 := "" //注：只能定义的局部变量
	fmt.Println(str, str2, str3)
}

//Q3:[primary] 通过指针变量 p 访问其成员变量 name，下面语法正确的是（）
//A. p.name
//B. (*p).name
//C. (&p).name
//D. p->name
//
//参考答案：AB
type Person struct {
	name string
}

func TestPointer(t *testing.T) {
	p := Person{"lehehe"}
	po := &p
	fmt.Println(po, *po) //注：&是指针，*是变量值

	p2 := &Person{"hahaha"}
	fmt.Println(p2.name) //指针成员变量可直接引用，编译器会自动隐形p转为（*p2）
	fmt.Println((*p2).name)
}

//Q4:[primary] 关于接口和类的说法，下面说法正确的是（）
//A. 一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口
//B. 实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理
//C. 类实现接口时，需要导入接口所在的包   										//测试：将interface独立分包，不需要导入即可引用
//D. 接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口
//
//参考答案：ABD

//这里说的类就是Go的结构体
type le interface {
	ha() string
}
type kaixin struct {
	name string
}

func (l *kaixin) ha(s string) string {
	return fmt.Sprintf("%s %s.", l.name, s)
}

func TestInterface(t *testing.T) {
	name := kaixin{"xiao"}
	status := name.ha("happy.")
	fmt.Println(status)
}

//Q5:[primary] 关于字符串连接，下面语法正确的是（）
//A. str := ‘abc’ + ‘123’
//B. str := "abc" + "123"
//C. str ：= '123' + "abc"
//D. fmt.Sprintf("abc%d", 123)
//
//参考答案：BD
func TestStr(t *testing.T) {
	str := "123" + `abc`
	fmt.Println(str)
}

//Q6:[primary] 关于协程，下面说法正确是（）
//A. 协程和线程都可以实现程序的并发执行
//B. 线程比协程更轻量级              	注：协程更轻量级
//C. 协程不存在死锁问题					注：同样有死锁问题
//D. 通过channel来进行协程间的通信
//
//参考答案：AD

//Q7:[intermediate] 关于init函数，下面说法正确的是（）
//A. 一个包中，可以包含多个init函数
//B. 程序编译时，先执行导入包的init函数，再执行本包内的init函数
//C. main包中，不能有init函数    							注：这个可以有
//D. init函数可以被其他函数调用 							 注：init函数是自动被调用的
//
//参考答案：AB

//Q8：[primary] 关于循环语句，下面说法正确的有（）
//A. 循环语句既支持for关键字，也支持while和do-while
//B. 关键字for的基本使用方法与C/C++中没有任何差异
//C. for循环支持continue和break来控制循环，但是它提供了一个更高级的break，可以选择中断哪一个循环
//D. for循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量
//
//参考答案：CD

func TestFor(t *testing.T) {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
	for i := 0; i < 100; i++ {
		fmt.Println("lehehe")
	}
}

//Q9:[intermediate] 对于函数定义：
//
//func add(args ...int) int {
//	sum := 0
//	for _, arg := range args {
//		sum += arg
//	}
//	return sum
//}
//下面对add函数调用正确的是（）
//A. add(1, 2)
//B. add(1, 3, 7)
//C. add([]int{1, 2})    		//要的是int,int切片不符合
//D. add([]int{1, 3, 7}...) 	//调用时...三个小点点，相当于遍历每一个int
//参考答案：ABD

func Add(args ...int) {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	fmt.Println(sum)
}

func TestAdd(t *testing.T) {
	Add(1, 2)
	Add([]int{1, 2, 3}...)
	//Add([]int{1, 2, 3}) //Error:cannot use []int literal (type []int) as type int in argument to Add
}

//Q10:[primary]关于类型转化，下面语法正确的是（）
//A.
//type MyInt int
//var i int = 1
//var j MyInt = i

//B.
//type MyInt int
//var i int = 1
//var j MyInt = (MyInt)i

//C.
//type MyInt int
//var i int = 1
//var j MyInt = MyInt(i)

//D.
//type MyInt int
//var i int = 1
//var j MyInt = i.(MyInt)

//参考答案：C
