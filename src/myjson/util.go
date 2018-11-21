package myjson

import (
	"strconv"
	"strings"
)

//在这写一个工具函数，用于判定是否一个字符串的number、true、false、null
func checkNTFN(str string) (NumberType, *Value){
	if strings.EqualFold(str, "true") {
		//不论true还是True还是truE还是怎么样的，不论大小写，都算是true
		return Bool, NewVal(true)
	} else if strings.EqualFold(str, "false") {
		return Bool, NewVal(false)
	} else if strings.EqualFold(str, "null") {
		return Null, NewValue()
	} else {
		if !(strings.Index(str, ".") >= 0) {
			//就当是整形
			v, err := strconv.Atoi(str)
			if err != nil {
				return Null, NewValue()
			} else {
				return  Int, NewVal(v)
			}
		} else {
			//就当是float64		为什么是就当呢，因为可能存在内容不是数字的字符串，这时候就不是数字啦，所以这里用“就当”
			//到底是不是数字,应该交给其他人解决
			v, err := strconv.ParseFloat(str, 64)
			if err != nil {
				//发生错误时就当null处理
				return Null, NewValue()
			} else {
				return Float64, NewVal(v)
			}
		}
	}
}

//处理value的值，将value按照不同情况识别，
//并按照其情况处理value的去向问题
//这个只处理非容器Value
func putValue(sa *StackAnaly,value *Value) {
	sign := sa.IsSign()
	if sign.GetWT() == TColon {
		//value是对象里的值
		//这里是冒号情况下value的情况，需要让：出栈
		if sa.State.GetOOA() {
			//其实一定在对象里的，冒号是不会在数组里出现的
			sa.State.Top().GetAsObjectIgnore()[keyStrs.Pop()] = value
		}
		//sa.Pop() //冒号pop出来
		sa.DeleteN(1)			// Pop -> DeleteN 注释同上
	} else if sign.GetWT() == TComma && !sa.State.GetOOA() {
		//value是作为数组中的元素的
		//不在对象中的逗号
		//这里Pop出来有Push回去低效率了，以后看到修复
		a := sa.State.Pop().GetAsSliceIgnore()
		a = append(a, value)
		sa.State.Push(NewVal(a))
		//sa.Pop() //逗号pop出来
		sa.DeleteN(1)			//Pop -> DeleteN 注释同上
	} else if sign.GetWT() == TSquareL {
		//value是数组的第一个元素
		a := sa.State.Pop().GetAsSliceIgnore()
		a = append(a, value)
		sa.State.Push(NewVal(a))
		//不用pop，[号是匹配符号
	}
}

//高效版压缩json
func CompressJson(bs []byte) []byte {
	//四个代替法则
	rp := strings.NewReplacer("\t", "",
										" ", "",
										"\n", "",
										"\r", "")
	return []byte(rp.Replace(string(bs)))
}

//还没有完成，将来肯定需要的函数
func FormatJson(bs []byte) []byte {
	str := string(bs)
	//做一些format的事情

	return []byte(str)
}

//分析参数用的
////str1 := "{\"name\":\"gyf\",\"age\":\"12\",\"ids\":{\"id1\":\"1\",\"id2\":\"2\"}}"
//以上的形成的js对象，param 为 name时得到的是以“gyf”为包装的*Value值   ids.id1 就是 “1”的*Value
//所以这里主要是将param字符串中的内容以”.“为分割 分割成string切片返回
func ParseParam(param string) []string {
	return strings.Split(param, ".")
}








