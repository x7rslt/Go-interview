//Q28:[intermediate] 关于GetPodAction定义，下面赋值正确的是（）
//###部分不理解
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
	//var fragment2 Fragment = *GetPodAction{}  			//具体什么意思？为什么不能*赋值；invalid indirect of GetPodAction literal (type GetPodAction)
	fmt.Println(fragment)
}

//Q28:[intermediate] 关于GoMock，下面说法正确的是（）
//###注：打桩知识点无知，待补习
//A. GoMock可以对interface打桩
//B. GoMock可以对类的成员函数打桩
//C. GoMock可以对函数打桩
//D. GoMock打桩后的依赖注入可以通过GoStub完成

//参考答案：AD

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

//Q48:[intermediate] 关于GoStub，下面说法正确的是（） 							//注：Gostub未接触
//A. GoStub可以对全局变量打桩
//B. GoStub可以对函数打桩
//C. GoStub可以对类的成员方法打桩
//D. GoStub可以打动态桩，比如对一个函数打桩后，多次调用该函数会有不同的行为
//
//参考答案：ABD
