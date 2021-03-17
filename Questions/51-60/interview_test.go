package interview

//51 [primary] 声明一个整型变量i__________
//参考答案：var i int
//
//52 [primary] 声明一个含有10个元素的整型数组a__________
//参考答案：var a [10]int
//
//53 [primary] 声明一个整型数组切片s__________
//参考答案：var s []int
//
//54 [primary] 声明一个整型指针变量p__________
//参考答案：var p *int
//
//55 [primary] 声明一个key为字符串型value为整型的map变量m__________
//参考答案：var m map[string]int
//
//56 [primary] 声明一个入参和返回值均为整型的函数变量f__________     //错：func f(x int)int{}，这个是f函数了，不是f变量函数
//参考答案：var f func(a int) int
//
//57 [primary] 声明一个只用于读取int数据的单向channel变量ch__________
//参考答案：var ch <-chan int
//
//58 [primary] 假设源文件的命名为slice.go，则测试文件的命名为__________
//参考答案：slice_test.go
//
//59 [primary] go test要求测试函数的前缀必须命名为__________
//参考答案：Test
//
//60 [intermediate] 下面的程序的运行结果是__________					//错：0 1 2 3 4，注意defer使前面的值最后输出
//for i := 0; i < 5; i++ {
//defer fmt.Printf("%d ", i)
//}
//参考答案：4 3 2 1 0
