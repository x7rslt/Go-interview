package interview

import (
	"fmt"
	"testing"
)

//Q41:[intermediate] 关于无缓冲和有冲突的channel，下面说法正确的是（）
//A. 无缓冲的channel是默认的缓冲为1的channel  			//为0
//B. 无缓冲的channel和有缓冲的channel都是同步的
//C. 无缓冲的channel和有缓冲的channel都是非同步的
//D. 无缓冲的channel是同步的，而有缓冲的channel是非同步的
//参考答案：D

//Q42:[intermediate] 关于异常的触发，下面说法正确的是（）
//A. 空指针解析
//B. 下标越界
//C. 除数为0
//D. 调用panic函数
//参考答案：ABCD

//Q43:[intermediate] 关于cap函数的适用类型，下面说法正确的是（）
//A. array
//B. slice
//C. map
//D. channel
//参考答案：ABD

// The cap built-in function returns the capacity of v, according to its type:
//	Array: the number of elements in v (same as len(v)).
//	Pointer to array: the number of elements in *v (same as len(v)).
//	Slice: the maximum length the slice can reach when resliced;
//	if v is nil, cap(v) is zero.
//	Channel: the channel buffer capacity, in units of elements;
//	if v is nil, cap(v) is zero.

//Q44:[intermediate] 关于beego框架，下面说法正确的是（）     					//注：beego未接触
//A. beego是一个golang实现的轻量级HTTP框架
//B. beego可以通过注释路由、正则路由等多种方式完成url路由注入
//C. 可以使用bee new工具生成空工程，然后使用bee run命令自动热编译
//D. beego框架只提供了对url路由的处理， 而对于MVC架构中的数据库部分未提供框架支持
//参考答案：ABC

//Q45:[intermediate] 关于goconvey，下面说法正确的是（）							//注：goconvey未接触
//A. goconvey是一个支持golang的单元测试框架
//B. goconvey能够自动监控文件修改并启动测试，并可以将测试结果实时输出到web界面
//C. goconvey提供了丰富的断言简化测试用例的编写
//D. goconvey无法与go test集成
//参考答案：ABC

//Q46:[intermediate] 关于go vet，下面说法正确的是（）							//注：go vet未接触
//A. go vet是golang自带工具go tool vet的封装
//B. 当执行go vet database时，可以对database所在目录下的所有子文件夹进行递归检测
//C. go vet可以使用绝对路径、相对路径或相对GOPATH的路径指定待检测的包
//D. go vet可以检测出死代码

//参考答案：ACD

//Q47:[intermediate] 关于map，下面说法正确的是（）
//A. map反序列化时json.unmarshal的入参必须为map的地址
//B. 在函数调用中传递map，则子函数中对map元素的增加不会导致父函数中map的修改
//C. 在函数调用中传递map，则子函数中对map元素的修改不会导致父函数中map的修改
//D. 不能使用内置函数delete删除map的元素
//
//参考答案：A

func TestMap(t *testing.T) {
	m := make(map[string]string)
	m["地址"] = "杭州"
	m["姓名"] = "Lehehe"
	fmt.Println(m)
	delete(m, "姓名")
	fmt.Println(m)
}

//Q48:[intermediate] 关于GoStub，下面说法正确的是（） 							//注：Gostub未接触
//A. GoStub可以对全局变量打桩
//B. GoStub可以对函数打桩
//C. GoStub可以对类的成员方法打桩
//D. GoStub可以打动态桩，比如对一个函数打桩后，多次调用该函数会有不同的行为
//
//参考答案：ABD

//Q49:[primary] 关于select机制，下面说法正确的是（）
//A. select机制用来处理异步IO问题
//B. select机制最大的一条限制就是每个case语句里必须是一个IO操作
//C. golang在语言级别支持select关键字
//D. select关键字的用法与switch语句非常类似，后面要带判断条件
//
//参考答案：ABC

//Q50:[primary] 关于内存泄露，下面说法正确的是（）
//A. golang有自动垃圾回收，不存在内存泄露
//B. golang中检测内存泄露主要依靠的是pprof包
//C. 内存泄露可以在编译阶段发现									//一般在运行阶段才能发现
//D. 应定期使用浏览器来查看系统的实时内存信息，及时发现内存泄露问题
//
//参考答案：BD
