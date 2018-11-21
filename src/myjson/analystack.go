package myjson

import (
	"fmt"
)

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
	//"1号，栈中是否只有奇数个”\"“,1:如果是则期待的是下一个” 处于接收string的状态下，0：如果不是，不在接收string的状态
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
	return &StackAnaly{make([]byte, 0, 35), -1, 0, s}
}

//实现StackSer的过程
func (s *StackAnaly) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

func (s *StackAnaly) Push(b byte) {
	s.PushWithoutCheck(b)
	switch b {
	case '"':
		//fmt.Println("wo jin lai le 0")
		s.flag = s.flag ^ 0x40			//每次压入一个"时，第二位标志位取反
	case '[':
		s.flag = s.flag | 0x80			//[设标志位
	}
	//fmt.Println(s.IsSign())
}

func (s *StackAnaly) PushWithoutCheck(b byte) {
	//其实就是len > top+1
	//因为存在pop了，但没删除的情况
	s.top++			//先自加后可以省略两个加法运算
	if len(s.data) != s.top {
		s.data[s.top] = b
	}
	s.data = append(s.data, b)
}

func (s *StackAnaly) Pop() byte {
	//当s。state中栈顶不是数组了，就置空标志位
	if s.State.GetOOA() {
		s.flag = s.flag & 0x7F
	}
	return s.PopWithoutCheck()
}

//标志位暂且没有作用，这个PopN也写成不会withoutCheck的
func (s *StackAnaly) PopN(n int) []byte {
	if s.top >= n-1 && n > 0{
		result := s.data[s.top-n+1:s.top+1]
		s.top = s.top - n
		return result
	}
	panic("PopN: top < n-1")
}

func (s *StackAnaly) PopWithoutCheck() byte {
	if s.IsEmpty() {
		fmt.Errorf("栈为空！！" )
		return 0
	}
	b := s.data[s.top]
	s.top--
	return b
}

//从栈顶删除元素，在不需要元素返回的时候可以调用这个函数，效率快
func (s *StackAnaly) DeleteN(n int) {
	if s.top >= n-1 {
		s.top = s.top - n
		return
	}
	panic("DeleteN: top < n-1")
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
	s.State.Clear()
	s.flag = 0x00
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
	//return GetSign([]byte{s.Top()}, s.flag)
	return s.IsSignN(0)
}

//n=0时与IsSign同
//n=1时就是除栈顶以外的第一个元素
func (s *StackAnaly) IsSignN(n int) *Sign {
	//flag还是用当前栈顶状态的值
	//是一个危险的函数，这个可能会越界，没有越界检查。之后看效率要不要添加检查
	return GetSign(s.data[s.top-n], s.flag)
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
	return &Stack{make([]*Value, 0, NumLayer), -1, true}
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
	s.obOrAr = true
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


type StackStringer interface {
	IsEmpty() bool
	Push(s string)
	Pop() string
	Top() string
	Clear()
	Size() int
}


//不管了，我在写一个string的栈吧，这样执行速度上将比上面的更加快写
type StackString struct {
	data []string
	top int
}

func NewStackString() *StackString {
	return &StackString{make([]string, 0, NumLayer), -1}
}

func (s *StackString)IsEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *StackString) Push(v string) {
	s.top++
	if len(s.data) > s.top {
		s.data[s.top] = v
	} else {
		s.data = append(s.data, v)
	}
}

func (s *StackString) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	v := s.data[s.top]
	s.top--
	return v
}

func (s *StackString) Top() string {
	if s.IsEmpty() {
		return ""
	}
	return s.data[s.top]
}

func (s *StackString) Size() int {
	return s.top + 1
}

func (s *StackString) Clear() {
	s.top = -1
}


