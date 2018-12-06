package myjson

import (
	"bytes"
	"strconv"
	"strings"
)

// 确定是bool还是null还是数字类型
// NBN NULL Bool Number
func checkNBN(bs []byte) (*Value) {
	if bytes.EqualFold(bs,[]byte{'t','r','u','e'}) {
		return NewVal(true)
	} else if bytes.EqualFold(bs, []byte{'f','a','l','s','e'}) {
		return NewVal(false)
	} else if bytes.EqualFold(bs, []byte{'n','u','l','l'}) {
		return NewValue()
	} else {
		if !(bytes.Index(bs, []byte{'.'}) >= 0) {
			v, err := strconv.Atoi(string(bs))
			if err == nil {
				return NewVal(v)
			} else {
				return NewValue()
			}
		} else {
			v, err := strconv.ParseFloat(string(bs), 64)
			if err == nil {
				return NewVal(v)
			} else {
				return NewValue()
			}
		}
	}
}

//处理value的值，将value按照不同情况识别，
//并按照其情况处理value的去向问题
//这个只处理非容器Value
func putValue(sa *stackAnaly,value *Value) {
	sign := sa.isSign()
	saState := sa.State;
	if sign == TColon {
		//value是对象里的值
		//这里是冒号情况下value的情况，需要让：出栈
		if saState.GetOOA() {
			//其实一定在对象里的，冒号是不会在数组里出现的
			saState.Top().GetAsObjectIgnore()[keyStrs.pop()] = value
		}
		//sa.Pop() //冒号pop出来
		sa.deleteN(1)			// Pop -> deleteN 注释同上
	} else if sign == TComma && !saState.GetOOA() {
		//value是作为数组中的元素的
		//不在对象中的逗号
		//这里Pop出来有Push回去低效率了，以后看到修复
		//a := saState.Pop().GetAsSliceIgnore()
		//a = append(a, value)
		//saState.Push(NewVal(a))

		a := saState.data[saState.top]
		a.AppendSlice(value)


		//sa.Pop() //逗号pop出来
		sa.deleteN(1)			//Pop -> deleteN 注释同上
	} else if sign == TSquareL {
		//value是数组的第一个元素
		//a := saState.Pop().GetAsSliceIgnore()
		//a = append(a, value)
		//saState.Push(NewVal(a))

		a := saState.data[saState.top]
		a.AppendSlice(value)

		//不用pop，[号是匹配符号
	}
}

//高效版压缩json

// Compress json
func CompressJson(bs []byte) []byte {
	//四个代替法则
	rp := strings.NewReplacer("\t", "",
										" ", "",
										"\n", "",
										"\r", "")
	return []byte(rp.Replace(string(bs)))
}

////还没有完成，将来肯定需要的函数
//func FormatJson(bs []byte) []byte {
//	str := string(bs)
//	//做一些format的事情
//
//	return []byte(str)
//}

//分析参数用的
////str1 := "{\"name\":\"gyf\",\"age\":\"12\",\"ids\":{\"id1\":\"1\",\"id2\":\"2\"}}"
//以上的形成的js对象，param 为 name时得到的是以“gyf”为包装的*Value值   ids.id1 就是 “1”的*Value
//所以这里主要是将param字符串中的内容以”.“为分割 分割成string切片返回
// Analysis parameters used
func parseParam(param string) []string {
	return strings.Split(param, ".")
}








