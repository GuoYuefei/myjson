package myjson

import (
	"fmt"
	"github.com/pkg/errors"
)

//终于到组装的底部了，可能零件还需要改动，但总的来说还是没问题的
var s *StackAnaly
var stack *Stack
var ss *StackAnaly			//辅助栈，出栈的时候倒序
var keyStr string			//永远只有一个key游离在外面

func init() {
	stack = NewStack()
	s = NewStackAnalyByStack(stack)
	ss = NewStackAnaly()
	keyStr = ""
}


//先处理一个字符
//如果读到的是json的关键字{}【】等等就返回true
func delChar(i int, b ...byte) {
					//i记录分析到的b的元素的指针
	var sign *Sign = s.IsSign()
	//这一部分是分析栈顶元素
	if !s.IsEmpty() {


	}
	s.Push(b[i])
	//更新下，
	sign = s.IsSign()
	//push完后分析状态
	if sign.GetStatus() == StaCloBrace {
		//当期待}的时候，接下来的东西在对象中
		jsob := make(JsObject)
		//这个stack实际上就是s对象里面的
		//将jsob封装成value放入栈中
		stack.Push(NewVal(jsob))
	}

	if sign.GetStatus() == StaSquare {
		//当期待]的时候，接下来的东西在数组中
		array := make([]*Value,0,20)

		stack.Push(NewVal(array))
	}

	if sign.GetWT() == TBracesR {
		//其实下面的判断一定是成立的，暂且找不出不成立的理由
		if stack.GetOOA() {
			jsob := stack.Pop().GetAsObjectIgnore()
			if stack.GetOOA() {
				stack.Top().GetAsObjectIgnore()[keyStr] = NewVal(jsob)
			} else {
				arr := stack.Top().GetAsSliceIgnore()
				arr = append(arr, NewVal(jsob))
			}
		}
	}

	if sign.GetWT() == TSquareR {
		//也一定成立的
		if !stack.GetOOA() {
			arr := stack.Pop().GetAsSliceIgnore()
			if stack.GetOOA() {
				stack.Top().GetAsObjectIgnore()[keyStr] = NewVal(arr)
			} else {
				arr := stack.Top().GetAsSliceIgnore()
				arr = append(arr,NewVal(arr))
			}
		}
	}

	//s栈标志表明此次push内容在一个字符串中
	if s.flag & 0x40 == 0x40 {
		//json的非关键字，前期我们就认为这个是一个合法的
		if s.IsSign() == nil {
			i++
			delChar(i, b...)
		} else if s.IsSign().GetWT() == TQuotation {
			//直到碰到下一个引号才结束
			s.Pop()
			for s.IsSign() == nil {
				ss.Push(s.Pop())
			}
			s.Pop()
			//循环外的Pop是将"\""清除

			var tempStr string = ""
			for !ss.IsEmpty() {
				tempStr = tempStr + string([]byte{ss.Pop()})
			}
			//关键是我现在怎么知道它是string还是key呢
			//key的情况有两种，一个前面元素也就是现在栈顶是{，一种是不在数组里面的“，”逗号
			sign = s.IsSign()
			if sign.GetStatus() == StaCloBrace || sign.GetStatus() == StaKey {
				keyStr = tempStr
				fmt.Println(keyStr)
			} else {
				if stack.GetOOA() {
					//如果stack栈顶是对象
					stack.Top().GetAsObjectIgnore()[keyStr] = NewVal(tempStr)
				} else {
					//如果栈顶是Array
					arr := stack.Top().GetAsSliceIgnore()
					arr = append(arr, NewVal(tempStr))
				}
			}

		} else {
			//在string里出现json的关键字了，事实上在value中允许的，为了方便暂且不允许
			fmt.Errorf("err" )
			panic(errors.New("err"))
		}
		return
	}
}



