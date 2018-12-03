package myjson

//终于到组装的底部了，可能零件还需要改动，但总的来说还是没问题的
var s *stackAnaly
var sState *Stack
//var ss *StackAnaly			//辅助栈，出栈的时候倒序
var keyStrs *stackString //永远只有一个key游离在外面

//为什么我感觉还需要一个string的栈啊，一个类型自己写个栈，我都累了，这样重复代码太多了，使用Value类型包装的栈用起来有太麻烦

func init() {
	sState = NewStack()
	s = newStackAnalyByStack(sState)
	//ss = NewStackAnaly()
	keyStrs = newStackString()
}

//为了增强效率，部分Pop函数用deleteN代替 用标志Pop -> deleteN标识

//现在我们需要几个函数，通过关键字可以分析出接下来的东西是key还是value，在数组中还是对象中
//能分析出value的简单类型是number还是bool还是null
/**
类型无非分成三类 	   	容器类型 		数组与对象
					普通类型		true,false,number,null
					”类型		字符串
可以在加上关键字		"关键字类型	{}[]":,
 */

 /**
 最复杂的逻辑写在了这一块
  */

 /**
 delChar这个函数可以分成三块，
 number、true、fasle、null判定块
 string判定块，可以分成key、value两块
 容器判定块，可以分成数组和对象
 三大块内部判定方式相似，但是三大块之间各有不同
  */


//先处理一个字符
//如果读到的是json的关键字{}【】等等就返回true
//这个函数只负责处理keyWord
func delChar(b byte) {
	//不在字符串状态下的空格换行等等字符不判定，不压栈
	if s.getFlag()&0x40 == 0x00 && (b == ' ' || b == '\n' || b == '\t' || b == '\r') {
		return
	}
	s.push(b)

	sign := s.isSign()

	//字符串或者不是标识符时，直接退出函数				en...冒号暂且也没有特殊的事件判断，故而省略
	if sign == TNone  || s.getFlag()&0x40 == 0x40 || sign == TColon {
		return
	}



	/**
	1、判定number、bool、null的方式无非就是在遇到 “, } ]”这三个符号的时候的前面一位的s.isSign().GetWT == TNone
	2、对应的判定这些结束的标志是： ， [
	 */
	if (sign == TComma || sign == TBracesR || sign == TSquareR) && s.isSignN(1) == TNone {
		// 将栈顶标识符弹出
		//v := s.PopWithoutCheck()

		////栈顶数起，第二个元素时json关键字
		//if s.isSignN(1) != TNone {
		//	//当标识符pop出来之后，还有关键字的话就不执行这个if中的任何内容
		//	//s.PushWithoutCheck(v)			//还原s
		//	//如果在，}]前面还有关键字，极大可能是引号，但也有可能是其他的
		//	goto KEYWORD1
		//}

		v := s.popWithoutCheck()

		//这个时候我们确定这里的值就是number，true，false，null了
		//当s.isSign.GetWT ！= TNone就认为value结束了，下一个肯定是标志符了
		//for s.isSign().GetWT() == TNone {
		//	//这里还需要分} ]这两种情况，因为这时候分别是一个对象和数组的结束
		//	//如果是} ]还需要跳转到} ]出执行
		//
		//	//先取出变量的字符串形式
		//	ss.PushWithoutCheck(s.Pop())
		//}
		//for !ss.IsEmpty() {
		//	temp = temp + string([]byte{ss.Pop()})
		//}

		//修改，该代码代替上面两个循环的注释代码，取得最终string
		var i int
		for i = 0; s.isSignN(i) == TNone; i++ {}
		temp := s.popN(i)

		_, value := checkNBN(temp)
		putValue(s, value)
		s.pushWithoutCheck(v)  		//还原} ] ,
		goto KEYWORD2			//如果是] }那么还需要处理
	}

//KEYWORD1:
	//-----------------------string(key-value)----------------------------------

	//如果是引号的话
	//且栈的状态标志在string档位，所以所有的字符进来全是string的一部分
	if sign == TQuotation && s.getFlag() & 0x40 == 0x00 {
		//如果接收到下一个引号了,那么就标志着一个字符串的结束
		//因为第二位标志符号会在第二个”引号来的时候重置

		s.deleteN(1)			//右引号删除
		//在没有碰到左引号时认为都是字符串里的东西
		//for s.isSign().GetWT() != TQuotation {
		//	//当栈顶元素不是引号时一直弹出到ss中
		//	ss.PushWithoutCheck(s.Pop())
		//}
		//for !ss.IsEmpty() {
		//	tempStr = tempStr + string([]byte{ss.Pop()})
		//}

		//!!!修改，该代码代替上面两个循环的注释代码
		var i int
		for i = 0; s.isSignN(i) != TQuotation; i++ {}
		tempStr := string(s.popN(i))


		s.deleteN(1)		//pop -> deleteN 删除左引号
		//接下来就要判定是key还是value了
		//当前栈顶{ ,不在数组中的逗号key
		//当前栈顶: 在数组中的的逗号value 还有[
		//-----------------key-----------------------
		//fmt.Println(s.isSign())
		//fmt.Println(string([]byte{s.Top()}),s.Size())
		if s.isSign() == TBracesL {
			//对象的第一个属性key
			keyStrs.push(tempStr)
		}
		if s.isSign() == TComma && sState.GetOOA() {
			keyStrs.push(tempStr)
			//s.Pop()			//顺便把，逗号pop出来了
			s.deleteN(1)		//pop -> deleteN
		}
		//----------------value---------------------
		putValue(s, NewVal(tempStr))
		return
	}
KEYWORD2:
//--------------------------------------keyWord-------------------------------
	//更新下sign，可能是goto到这的
	//sign = s.isSign()			//貌似没必要
	//push完后分析状态
	//如果是关键字的话

	if sign == TBracesL {
		//开始一个左大括号的时候，接下来的东西在对象中
		jsob := make(JsObject)
		//这个stack实际上就是s对象里面的
		//将jsob封装成value放入栈中
		sState.Push(NewVal(jsob))
		return
	}

	if sign == TSquareL {
		//当开始一个左中括号的时候，接下来的内容在一个数组里面
		//fmt.Println(string([]byte{s.Top()}))
		//当栈顶为[的时候，接下来的东西在数组中
		array := make([]*Value,0,10)
		sState.Push(NewVal(array))
		return
	}

	if sign == TBracesR {
		//其实下面的判断一定是成立的，暂且找不出不成立的理由
		if sState.GetOOA() {
			if sState.Size() == 1 {
				//只剩下最后一个对象了，不就是json全解析完了嘛。。哈哈
				//s.Pop();s.Pop();		//最后两个{}pop出来
				s.deleteN(2)			//注释同上 pop -> deleteN
				return
			}
			jsob := sState.Pop().GetAsObjectIgnore()

			if sState.GetOOA() {
				//对象中的对象需要pop出}{：
				s.deleteN(3)
				sState.Top().GetAsObjectIgnore()[keyStrs.pop()] = NewVal(jsob)
			} else {
				//arr := sState.Pop().GetAsSliceIgnore()
				//arr = append(arr, NewVal(jsob))
				//sState.Push(NewVal(arr))

				//数组中的对象只需要pop出}{
				s.deleteN(2)			//注释同上 ，修改 pop -> deleteN		deleteN的速度更快

				//!!替换以上代码，更新切片是Value类型应该做的事情
				arr := sState.data[sState.top]
				arr.AppendSlice(NewVal(jsob))
			}
		}
		return
	}

	if sign == TSquareR {
		//也一定成立的
		//fmt.Println(string([]byte{s.Top()}))
		if !sState.GetOOA() {
			arr := sState.Pop().GetAsSliceIgnore()
			//s.Pop()   			//]pop
			s.deleteN(1)			//注释同上 Pop -> deleteN
			if sState.GetOOA() {
				//s.Pop();s.Pop() 		//[pop he : pop
				s.deleteN(2)			//注释同上 Pop -> deleteN
				sState.Top().GetAsObjectIgnore()[keyStrs.pop()] = NewVal(arr)
			} else {
				//该数组作为栈顶数组的元素
				//s.Pop();			//[pop
				s.deleteN(1)		//注释同上 Pop -> deleteN
				if s.isSign() == TComma {
					//当s栈顶是，
					//s.Pop()				//把，pop出去
					s.deleteN(1)			//注释同上 Pop -> deleteN
				}
				//temp := sState.Pop().GetAsSliceIgnore()
				//temp = append(temp,NewVal(arr))
				//sState.Push(NewVal(temp))

				//！！！替换以上代码
				temp := sState.data[sState.top]
				temp.AppendSlice(NewVal(arr))
			}
		}
	}
	return
}

func clearAllStack() {
	//ss.Clear()
	keyStrs.clear()
	s.clear()			//sState在内部被Clear了
}
