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
