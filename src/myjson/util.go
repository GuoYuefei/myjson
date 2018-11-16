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
func CompressJson(bs []byte) []byte {
	str := string(bs)
	//去除空格
	str = strings.Replace(str, " ", "", -1)
	//去除换行
	str = strings.Replace(str, "\n", "", -1)
	//windows下还可能有回车
	str = strings.Replace(str, "\r", "", -1)
	return []byte(str)
}
