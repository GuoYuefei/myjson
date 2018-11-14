package myjson

import "fmt"

//这是一个分析栈，json其实就是在分析一个一个的字符
///嗯。。可能我要重复造轮子了，还是打算自己写个栈 针对byte类型

//其实这个栈类型还需要一些其他功能用于分析json，所以叫做分析栈
//分析的时候需要做各种标志，使用一个8位的byte类型做标志
//不知道做出来的时候支不支持中文，可以做测试

//只针对byte的Stack  接口名中的S代表string
//定义了一个基础栈需要的方法
type StackSer interface {
	//移除并取得栈顶元素
	Pop() byte
	//只取得栈顶元素
	Top() byte
	//压入栈
	Push(b byte)
	//清空栈
	Clear()
	//是否为空 空返回true，否则返回false
	IsEmpty() bool
	//获取栈中元素的个数
	Size() int
}

type StackAnaly struct {
	//数据所在容器
	data []byte
	//top元素所在地的指针
	top int
	//这是一个功能标志
	//1111,1111 从左往右一次编号01234567
	//,0号，是否在[]中，1：如果是”，“号所期待的下一个值是value，0：如果不是，则期待key。
	//"1号，栈中是否只有奇数个”\"“,1:如果是则期待的是下一个”，0：如果不是，啥也不期待，继承上一个期待
	flag byte
	//这个栈中存有当前元素应该在对象中还是数组中，实在在什么对象中什么数组中
	State *Stack
}

//忽然发现还需要一个value栈，写在后面了，嗯，重复代码就重复了吧。这里我还是比较想追求一个效率的
//其实有办法可以减少代码量，但是函数调用太多了，byte栈是经常在用的，所以不好牺牲效率

func NewStackAnaly() *StackAnaly {
	//起初的byte先放50个，以后看情况定吧
	return NewStackAnalyByStack(NewStack())
}

func NewStackAnalyByStack(s *Stack) *StackAnaly {
	return &StackAnaly{make([]byte, 0, 50), -1, 0, s}
}

//实现StackSer的过程
func (s *StackAnaly) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

func (s *StackAnaly) Push(b byte) {
	//其实就是len > top+1
	//因为存在pop了，但没删除的情况
	s.top++			//先自加后可以省略两个加法运算
	if len(s.data) != s.top {
		s.data[s.top] = b
	}
	s.data = append(s.data, b)
	switch string([]byte{b}) {
	case "\"":
		//fmt.Println("wo jin lai le 0")
		s.flag = s.flag ^ 0x40			//每次压入一个"时，第二位标志位取反
	case "[":
		s.flag = s.flag | 0x80			//[设标志位
	}
	//fmt.Println(s.IsSign())

}

func (s *StackAnaly) Pop() byte {
	if s.IsEmpty() {
		fmt.Errorf("栈为空！！" )
		return 0
	}
	//必须在top--之前
	if s.IsSign() != nil {
		switch s.IsSign().GetWT() {
		//引号情况居多，引号放前
		case TQuotation:						//如果是引号，标志置空
			s.flag = s.flag ^ 0xBF			//去掉标志位
			//case TSquareL:						//[出栈的时候将标志置空,------------------------------------->但是如果是两层嵌套呢，这里是一个bug，暂不管
			//	s.flag = s.flag & 0x7F
		}

	}

	//当s。state中栈顶不是数组了，就置空标志位
	if s.State.GetOOA() {
		s.flag = s.flag & 0x7F
	}

	b := s.data[s.top]
	s.top-- //只减不删
	return b
}

func (s *StackAnaly) Top() byte {
	if s.IsEmpty() {
		fmt.Errorf("栈为空！" )
		return 0
	}
	return s.data[s.top]
}

func (s *StackAnaly) Size() int {
	return s.top + 1
}

func (s *StackAnaly) Clear() {
	s.top = -1
}

//-----------------接下来需要考虑为这个栈增添什么功能-----------

type StackAnalyser interface {
	StackSer
	Analyser
}

//栈的分析接口
type Analyser interface {
	IsSign() *Sign
	SetFlag(b byte)
	GetFlag() byte
}
//需要分析出标志符{}[],":七个

//查看刚压入的元素是否是标识符 如果是就返回有值的sign，否则就nil
func (s *StackAnaly) IsSign() *Sign {
	return GetSign([]byte{s.Top()}, s.flag)
}

//置位标识符 使用位
func (s *StackAnaly) SetFlag(b byte) {
	s.flag = b
}

//得到标志位
func (s *StackAnaly) GetFlag() byte {
	return s.flag
}


//-----------------这里写value的栈------------------

type StackValuer interface {
	//移除并取得栈顶元素
	Pop() *Value
	//只取得栈顶元素
	Top() *Value
	//压入栈
	Push(b *Value)
	//清空栈
	Clear()
	//是否为空 空返回true，否则返回false
	IsEmpty() bool
	//获取栈中元素的个数
	Size() int
}

//这个类似一个万能的栈了
//主要用来存放js对象和数组的
type Stack struct {
	//可能也会放key的哦
	data []*Value
	top int
	//StackValuer
	//顶层元素在对象里面还是数组里面
	//在对象里置true 数组中就是false
	obOrAr bool
}

func NewStack() *Stack {
	return &Stack{make([]*Value, 0,24), -1, true}
}

func (s *Stack) IsEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *Stack) Push(v *Value) {
	s.top++
	if len(s.data) > s.top {
		s.data[s.top] = v
	} else {
		s.data = append(s.data, v)
	}
	s.setOOA()

}

//根据顶层元素置位，当顶层元素发生变化时必须调用
func (s *Stack)setOOA() {
	v := s.data[s.top]
	//如果是jsob。那么就置obOrAr为true
	if v.IsObject() {
		s.obOrAr = true
	}

	if v.IsSlice() {
		s.obOrAr = false
	}
}



func (s *Stack) Pop() *Value {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[s.top]
	s.top--

	s.setOOA()

	return v
}

func (s *Stack) Top() *Value {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top]
}

func (s *Stack) Size() int {
	return s.top + 1
}

func (s *Stack) Clear() {
	s.top = -1
}
//
//func (s *Stack) SetTrue() {
//	s.obOrAr = true
//}
//
//func (s *Stack) SetFalse() {
//	s.obOrAr = false
//}

func (s *Stack) GetOOA() bool {
	return s.obOrAr
}


