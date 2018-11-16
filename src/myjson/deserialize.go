package myjson

import (
	"bufio"
	"io"
)

/**
反序列化的各个函数接口
 */

func GetJsObject(json []byte) (JsObject, error) {
	clearAllStack()
	json = CompressJson(json)			//必须先压缩，在识别
	for _,b := range json {
		delChar(b)
	}
	return sState.Top().GetAsObject()
}

func GetJsObjectByBufReader(reader *bufio.Reader) (JsObject, error) {
	p := make([]byte, 16)			//每次取16字节数据
	json := make([]byte, 0, 112)
	for {
		n, err := reader.Read(p)
		json = append(json, p[:n]...)
		if err == io.EOF {
			break
		}
	}
	return GetJsObject(json)
}

func GetJsObjectByReader(reader io.Reader) (JsObject, error) {
	bufReader := bufio.NewReader(reader)
	return GetJsObjectByBufReader(bufReader)
}



