package myjson

import "fmt"

//终于到组装的底部了，可能零件还需要改动，但总的来说还是没问题的
var s *StackAnaly
var sState *Stack
var ss *StackAnaly			//辅助栈，出栈的时候倒序
var keyStrs *StackString			//永远只有一个key游离在外面

//为什么我感觉还需要一个string的栈啊，一个类型自己写个栈，我都累了，这样重复代码太多了，使用Value类型包装的栈用起来有太麻烦

func init() {
	sState = NewStack()
	s = NewStackAnalyByStack(sState)
	ss = NewStackAnaly()
	keyStrs = NewStackString()
}

//现在我们需要几个函数，通过关键字可以分析出接下来的东西是key还是value，在数组中还是对象中
//能分析出value的简单类型是number还是bool还是null
/**
类型无非分成三类 	   	容器类型 		数组与对象
					普通类型		true,false,number,null
					”类型		字符串
可以在加上关键字		"关键字类型	{}[]":,
 */



//先处理一个字符
//如果读到的是json的关键字{}【】等等就返回true
//这个函数只负责处理keyWord
func delChar(b byte) {
	var sign *Sign
	//这一部分是分析栈顶元素
	if !s.IsEmpty() {


	}
	s.Push(b)
	//更新下，
	sign = s.IsSign()

	if sign == nil {
		//当不是json关键字时就认为是字符串的一部分（此时s.GetFlag() & 0x40==0x40）
		//或者是number、true、false、null

		//暂且不分辨
		return
	}


	//-----------------------string(key-value)----------------------------------

	//如果是引号的话
	if sign.GetWT() == TQuotation {
		fmt.Println("wojin lai le 1")
		//栈的状态标志在string档位，所以所有的字符进来全是string的一部分
		if s.GetFlag() & 0x40 == 0x00 {
			//如果接收到下一个引号了,那么就标志着一个字符串的结束
			//因为第二位标志符号会在第二个”引号来的时候重置
			fmt.Println("wo jin lai le 2")
			var tempStr string = ""
			s.Pop()
			for s.IsSign() == nil {
				//当栈顶元素不是引号时一直弹出到ss中
				ss.Push(s.Pop())
			}
			s.Pop()		//剩余的“pop出来
			for !ss.IsEmpty() {
				tempStr = tempStr + string([]byte{ss.Pop()})
			}
			//接下来就要判定是key还是value了
			//当前栈顶{ ,不在数组中的逗号key
			//当前栈顶: 在数组中的的逗号value 还有[
			//-----------------key-----------------------
			fmt.Println(s.IsSign())
			fmt.Println(string([]byte{s.Top()}),s.Size())
			if s.IsSign().GetWT() == TBracesL {
				//对象的第一个属性key
				keyStrs.Push(tempStr)
			}
			if s.IsSign().GetWT() == TComma && s.State.GetOOA() {
				keyStrs.Push(tempStr)
				s.Pop()			//顺便把，逗号pop出来了
			}


			//----------------value---------------------
			if s.IsSign().GetWT() == TColon {
				//这里是冒号情况下value的情况，需要让：出栈
				if s.State.GetOOA() {
					//其实一定在对象里的，冒号是不会在数组里出现的
					s.State.Top().GetAsObjectIgnore()[keyStrs.Pop()] = NewVal(tempStr)
				}
				s.Pop()		//冒号pop出来
			} else if s.IsSign().GetWT() == TComma && !s.State.GetOOA() {
				//这里Pop出来有Push回去低效率了，以后看到修复
				a := s.State.Pop().GetAsSliceIgnore()
				a = append(a, NewVal(tempStr))
				s.State.Push(NewVal(a))
				s.Pop()			//逗号pop出来
			} else if s.IsSign().GetWT() == TSquareL {
				//数组的第一个元素
				a := s.State.Pop().GetAsSliceIgnore()
				a = append(a, NewVal(tempStr))
				s.State.Push(NewVal(a))
				//不用pop，[号是匹配符号
			}
			return
		}

	}


	//-------------------------分析number-----------------------------
	//首先number是value
	//if s.IsSign()==nil && s.IsSign().GetWT() == TColon || (s.IsSign().GetWT() == TColon && !s.State.GetOOA()) || s.IsSign().GetWT() == TSquareL {
	//
	//}

//--------------------------------------keyWord-------------------------------
	//push完后分析状态
	//如果是关键字的话
	if s.IsSign() != nil {
		if sign.GetStatus() == StaQuotation {
			//刚刚在s中push了一个奇数个引号，其实整个栈中就一个引号
			//return false		//返回一个false，需要一个新的分析函数,就叫字符串分析函数吧，除了解析出字符串以外还需要分辨是key还是value
			//什么也不用做，push的过程已经全部做好了
		}

		if sign.GetStatus() == StaCloBrace {
			//当期待}的时候，接下来的东西在对象中
			jsob := make(JsObject)
			//这个stack实际上就是s对象里面的
			//将jsob封装成value放入栈中
			s.State.Push(NewVal(jsob))
		}

		if sign.GetStatus() == StaSquare {
			//当期待]的时候，接下来的东西在数组中
			array := make([]*Value,0,20)

			s.State.Push(NewVal(array))
		}

		if sign.GetWT() == TBracesR {
			//其实下面的判断一定是成立的，暂且找不出不成立的理由
			if s.State.GetOOA() {
				if s.State.Size() == 1 {
					//只剩下最后一个对象了，不就是json全解析完了嘛。。哈哈
					s.Pop();s.Pop();		//最后两个{}pop出来
					return
				}
				jsob := s.State.Pop().GetAsObjectIgnore()
				s.Pop()				//} pop出来
				s.Pop()				//: or ,  pop出来，属性前面肯定有冒号啊
				s.Pop()				//{ pop出来
				if s.State.GetOOA() {
					s.State.Top().GetAsObjectIgnore()[keyStrs.Pop()] = NewVal(jsob)
				} else {
					arr := s.State.Top().GetAsSliceIgnore()
					arr = append(arr, NewVal(jsob))
				}
			}
		}

		if sign.GetWT() == TSquareR {
			//也一定成立的
			if !s.State.GetOOA() {
				arr := s.State.Pop().GetAsSliceIgnore()
				s.Pop()   			//]pop
				if s.State.GetOOA() {
					s.Pop();s.Pop() 		//[pop he : pop
					s.State.Top().GetAsObjectIgnore()[keyStrs.Pop()] = NewVal(arr)
				} else {
					s.Pop();s.Pop()			//[pop he ,pop
					arr := s.State.Top().GetAsSliceIgnore()
					arr = append(arr,NewVal(arr))
				}
			}
		}

	}

	return

}



