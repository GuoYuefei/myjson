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
type stackSer interface {
	//移除并取得栈顶元素
	pop() byte
	//只取得栈顶元素
	front() byte
	//压入栈
	push(b byte)
	//清空栈
	clear()
	//是否为空 空返回true，否则返回false
	isEmpty() bool
	//获取栈中元素的个数
	size() int
}



type stackAnaly struct {
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

func newStackAnaly() *stackAnaly {
	//起初的byte先放50个，以后看情况定吧
	return newStackAnalyByStack(NewStack())
}

func newStackAnalyByStack(s *Stack) *stackAnaly {
	return &stackAnaly{make([]byte, 0, 35), -1, 0, s}
}

//实现StackSer的过程
func (s *stackAnaly) isEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

func (s *stackAnaly) push(b byte) {
	s.pushWithoutCheck(b)
	switch b {
	case '"':
		//fmt.Println("wo jin lai le 0")
		s.flag = s.flag ^ 0x40			//每次压入一个"时，第二位标志位取反

		//发现这并没用到
	//case '[':
	//	s.flag = s.flag | 0x80			//[设标志位
	}
	//fmt.Println(s.IsSign())
}

func (s *stackAnaly) pushWithoutCheck(b byte) {
	//其实就是len > top+1
	//因为存在pop了，但没删除的情况
	s.top++			//先自加后可以省略两个加法运算
	if len(s.data) != s.top {
		s.data[s.top] = b
	}
	s.data = append(s.data, b)
}

func (s *stackAnaly) pop() byte {
	//当s。state中栈顶不是数组了，就置空标志位
	if s.State.GetOOA() {
		s.flag = s.flag & 0x7F
	}
	return s.popWithoutCheck()
}

//标志位暂且没有作用，这个PopN也写成不会withoutCheck的
func (s *stackAnaly) popN(n int) []byte {
	if s.top >= n-1 && n > 0{
		result := s.data[s.top-n+1:s.top+1]
		s.top = s.top - n
		return result
	}
	panic("PopN: top < n-1")
}

func (s *stackAnaly) popWithoutCheck() byte {
	if s.isEmpty() {
		fmt.Errorf("栈为空！！" )
		return 0
	}
	b := s.data[s.top]
	s.top--
	return b
}

//从栈顶删除元素，在不需要元素返回的时候可以调用这个函数，效率快
func (s *stackAnaly) deleteN(n int) {
	if s.top >= n-1 {
		s.top = s.top - n
		return
	}
	panic("DeleteN: top < n-1")
}

func (s *stackAnaly) front() byte {
	if s.isEmpty() {
		fmt.Errorf("栈为空！" )
		return 0
	}
	return s.data[s.top]
}

func (s *stackAnaly) size() int {
	return s.top + 1
}

func (s *stackAnaly) clear() {
	s.top = -1
	s.State.Clear()
	s.flag = 0x00
}

//-----------------接下来需要考虑为这个栈增添什么功能-----------

type stackAnalyser interface {
	stackSer
	analyser
}

//栈的分析接口
type analyser interface {
	isSign() WhatsType
	setFlag(b byte)
	getFlag() byte
}
//需要分析出标志符{}[],":七个

//查看刚压入的元素是否是标识符 如果是就返回有值的sign，否则就nil
func (s *stackAnaly) isSign() WhatsType {
	//return GetSign([]byte{s.Top()}, s.flag)
	return s.isSignN(0)
}

func (s *stackAnaly) isSignN(n int) (wt WhatsType) {
	switch s.data[s.top-n] {
	//出现频率应该是最高的放在前面
	case ',':
		wt = TComma
	case ':':
		wt = TColon
	case '"':
		wt = TQuotation
	case '{':
		wt = TBracesL
	case '[':
		wt = TSquareL
	case '}':
		wt = TBracesR
	case ']':
		wt = TSquareR
	default:
		wt = TNone
	}
	return
}

//置位标识符 使用位
func (s *stackAnaly) setFlag(b byte) {
	s.flag = b
}

//得到标志位
func (s *stackAnaly) getFlag() byte {
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

// This is a stack that can store any type
// Depends on the Value type
type Stack struct {
	//可能也会放key的哦
	data []*Value
	top int
	//StackValuer
	//顶层元素在对象里面还是数组里面
	//在对象里置true 数组中就是false
	obOrAr bool
}

// New a Stack
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
		return
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

// If it is judged whether the Value type that Push enters represents this array or an object,
// the default is the object.
func (s *Stack) GetOOA() bool {
	return s.obOrAr
}


type stackStringer interface {
	isEmpty() bool
	push(s string)
	pop() string
	front() string
	clear()
	size() int
}

// 不管了，我在写一个string的栈吧，这样执行速度上将比上面的更加快写
type stackString struct {
	data []string
	top int
}

func newStackString() *stackString {
	return &stackString{make([]string, 0, NumLayer), -1}
}

func (s *stackString)isEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *stackString) push(v string) {
	s.top++
	if len(s.data) > s.top {
		s.data[s.top] = v
	} else {
		s.data = append(s.data, v)
	}
}

func (s *stackString) pop() string {
	if s.isEmpty() {
		return ""
	}
	v := s.data[s.top]
	s.top--
	return v
}

func (s *stackString) front() string {
	if s.isEmpty() {
		return ""
	}
	return s.data[s.top]
}

func (s *stackString) size() int {
	return s.top + 1
}

func (s *stackString) clear() {
	s.top = -1
}


