package myjson

import (
	"strconv"
	"strings"
)

//在这写一个工具函数，用于判定是否一个字符串的number、true、false、null
func CheckNTFN(str string) (NumberType, *Value){
	if strings.EqualFold(str, "true") {
		//不论true还是True还是truE还是怎么样的，不论大小写，都算是true
		return True, NewVal(true)
	} else if strings.EqualFold(str, "false") {
		return False, NewVal(false)
	} else if strings.EqualFold(str, "null") {
		return Null, NewValue()
	} else {
		if strings.Index(str, ".") >= 0 {
			//就当是float64		为什么是就当呢，因为可能存在内容不是数字的字符串，这时候就不是数字啦，所以这里用“就当”
			//到底是不是数字,应该交给其他人解决
			v, err := strconv.ParseFloat(str, 64)
			if err != nil {
				//发生错误时就当null处理
				return Null, NewValue()
			} else {
				return Float64, NewVal(v)
			}
		} else {
			//就当是整形
			v, err := strconv.Atoi(str)
			if err != nil {
				return Null, NewValue()
			} else {
				return  Int, NewVal(v)
			}
		}
	}
}

func PutValue(sa *StackAnaly,value *Value) {
	if sa.IsSign().GetWT() == TColon {
		//这里是冒号情况下value的情况，需要让：出栈
		if sa.State.GetOOA() {
			//其实一定在对象里的，冒号是不会在数组里出现的
			sa.State.Top().GetAsObjectIgnore()[keyStrs.Pop()] = value
		}
		sa.Pop() //冒号pop出来
	} else if sa.IsSign().GetWT() == TComma && !sa.State.GetOOA() {
		//不在对象中的逗号
		//这里Pop出来有Push回去低效率了，以后看到修复
		a := sa.State.Pop().GetAsSliceIgnore()
		a = append(a, value)
		sa.State.Push(NewVal(a))
		sa.Pop() //逗号pop出来
	} else if sa.IsSign().GetWT() == TSquareL {
		//数组的第一个元素
		a := sa.State.Pop().GetAsSliceIgnore()
		a = append(a, value)
		sa.State.Push(NewVal(a))
		//不用pop，[号是匹配符号
	}
}

//压缩json
//string类型中的空格换行什么的都不能去除
func CompressJson(bs []byte) []byte {
	str := string(bs)
	//有json字符串规律可知，json的字符串中引号不会出现开头与结尾
	after := strings.SplitAfter(str, "\"")
	str = ""
	for i, v := range after {
		if i % 2 == 0 {
			//偶数段一定不是字符串段
			v = strings.Replace(v, " ","", -1)			//去除空格
			//去除换行
			v = strings.Replace(v, "\n","", -1)
			//windows下还可能有回车
			v = strings.Replace(v, "\r", "", -1)
		}
		str = str + v
	}
	return []byte(str)
}

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








