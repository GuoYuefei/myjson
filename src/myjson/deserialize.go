package myjson

import (
	"bufio"
	"io"
)

/**
反序列化的各个函数接口
 */

//基础函数，底层函数，也比较好用
//根据字节数组解析成相应的JsObject，成功error就返回nil，否则返回相应错误

// GetJsObject Basic functions, the underlying functions, also easier to use.
// According to the byte array parsed into the corresponding JsObject, if successful, the return value of error is nil, otherwise return the corresponding error.
func GetJsObject(json []byte) (JsObject, error) {
	clearAllStack()

	for _,b := range json {
		delChar(b)
	}
	return sState.Top().GetAsObject()
}

// GetJsObjectByBufReader Get Js object according to bufio.Reader.
// Return situation is the same as func GetJsObject.
func GetJsObjectByBufReader(reader *bufio.Reader) (JsObject, error) {
	p := make([]byte, 32)			//每次取32字节数据
	json := make([]byte, 0, 224)
	for {
		n, err := reader.Read(p)
		json = append(json, p[:n]...)
		if err == io.EOF {
			break
		}
	}
	return GetJsObject(json)
}

// GetJsObjectByReader Get Js Object according to io.Reader.
// Return situation is the same as func GetJsObject.
func GetJsObjectByReader(reader io.Reader) (JsObject, error) {
	bufReader := bufio.NewReader(reader)
	return GetJsObjectByBufReader(bufReader)
}

//根据param得到JsObject中的内容
//参照Get函数的注释资料

// GetFromJsObject Get the contents of JsObject according to param.
// Refer to the annotation data of the Get function.
func GetFromJsObject(object JsObject, param string) *Value {
	params := parseParam(param)
	var jso JsObject = object
	for i, v := range params {
		if i == len(params) - 1 {
			//最后一个可能不是object
			break
		}
		jso = jso[v].GetAsObjectIgnore()
	}
	return jso[params[len(params) - 1]]
}


//根据para参数获取Value的值
//str1 := "{\"name\":\"gyf\",\"age\":\"12\",\"ids\":{\"id1\":\"1\",\"id2\":\"2\"}}"
//比如以上的json字符串作为[]byte参数
//“ids.id1”作为param参数
//得到的结果就是NewVal("1")的指针

// Get Get the value of Value according to the para parameter,
// para := "{\"name\":\"gyf\",\"age\":\"12\",\"ids\":{\"id1\":\"1\", \"id2\":\"2\"}}",
// Such as the above json string as a [] byte parameter,
// "ids.id1" as the param parameter,
// The result is a pointer to NewVal("1").
// If you need to get the final result, just refer to the Document of the Value type.
func Get(json []byte, param string) *Value{
	object, e := GetJsObject(json)
	if e != nil {
		panic(e)		//暂时不处理
	}
	return GetFromJsObject(object, param)
}


