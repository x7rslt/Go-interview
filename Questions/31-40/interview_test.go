package interview

import (
	"fmt"
	"testing"
	"time"
)

//Q31:[primary] 关于channel，下面语法正确的是（）
//A. var ch chan int
//B. ch := make(chan int)
//C. <- ch
//D. ch <- 					//赋值，未见值
//参考答案：ABC

//Q32:[primary] 关于同步锁，下面说法正确的是（）
//A. 当一个goroutine获得了Mutex后，其他goroutine就只能乖乖的等待，除非该goroutine释放这个Mutex
//B. RWMutex在读锁占用的情况下，会阻止写，但不阻止读
//C. RWMutex在写锁占用情况下，会阻止任何其他goroutine（无论读和写）进来，整个锁相当于由该goroutine独占
//D. Lock()操作需要保证有Unlock()或RUnlock()调用与之对应
//参考答案：ABC

//D  RUnlock方法应该和读写锁对应

//Q33:[intermediate] golang中大多数数据类型都可以转化为有效的JSON文本，下面几种类型除外（）
//A. 指针
//B. channel
//C. complex
//D. 函数
//参考答案：BCD

//JSON 格式是借这个样子滴 {"name":"小红"} ，指针是没问题的

//BCD怎么变成上面的格式？ 不知道怎么搞了，那就选A吧

//Q34:[intermediate] 关于go vendor，下面说法正确的是（）
//A. 基本思路是将引用的外部包的源代码放在当前工程的vendor目录下面
//B. 编译go代码会优先从vendor目录先寻找依赖包
//C. 可以指定引用某个特定版本的外部包
//D. 有了vendor目录后，打包当前的工程代码到其他机器的$GOPATH/src下都可以通过编译
//参考答案：ABD，参考官方文档

//Q35:[primary] flag是bool型变量，下面if表达式符合编码规范的是（）
//A. if flag == 1
//B. if flag
//C. if flag == false
//D. if !flag
//参考答案：BCD

//Q36:[primary] value是整型变量，下面if表达式符合编码规范的是（）
//A. if value == 0
//B. if value
//C. if value != 0
//D. if !value
//参考答案：AC

//Q37:[intermediate] 关于函数返回值的错误设计，下面说法正确的是（）
//A. 如果失败原因只有一个，则返回bool
//B. 如果失败原因超过一个，则返回error
//C. 如果没有失败原因，则不返回bool或error
//D. 如果重试几次可以避免失败，则不要立即返回bool或error
//参考答案：ABCD

//函数返回值，尽量详细告诉调用者，函数执行后的结果，应该有错误返回值，例如channel中

//Q38:[intermediate] 关于异常设计，下面说法正确的是（）
//A. 在程序开发阶段，坚持速错，让程序异常崩溃
//B. 在程序部署后，应恢复异常避免程序终止
//C. 一切皆错误，不用进行异常设计
//D. 对于不应该出现的分支，使用异常处理
//参考答案：ABD

//Q39:[intermediate] 关于slice或map操作，下面正确的是（）
//A.
//var s []int
//s = append(s,1)
//B.
//
//var m map[string]int
//m["one"] = 1
//C.
//
//var s []int
//s = make([]int, 0)
//s = append(s,1)
//D.
//
//var m map[string]int
//m = make(map[string]int)
//m["one"] = 1
//参考答案：ACD

//B 没有初始化，所以m是nil，赋值，直接抛了,其他的都Ok

func TestMap(t *testing.T) {
	var m map[string]int //Error:panic: assignment to entry in nil map [recovered]
	//注：没有初始化，所以m是nil，赋值，直接抛了
	//m["one"] = 1
	fmt.Println(m)

	m2 := make(map[string]int)
	m2["one"] = 1
	fmt.Println(m2)
}

//Q40:[intermediate] 关于channel的特性，下面说法正确的是（）
//A. 给一个 nil channel 发送数据，造成永远阻塞
//B. 从一个 nil channel 接收数据，造成永远阻塞			//注AB必须有发有接，否则就会一直堵塞
//C. 给一个已经关闭的 channel 发送数据，引起 panic
//D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
//参考答案：ABCD

func TestChannel(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 1 //只发送，未接收，一直堵塞
	}()
	//fmt.Println(<-c) 					//通道接收，解除堵塞
	time.Sleep(1 * time.Second)
	fmt.Println("Channel test")
}
