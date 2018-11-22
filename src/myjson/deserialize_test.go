package myjson

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetJsObject(t *testing.T) {
	str := "{\"name\":\"solomon\",\"age\":22}"
	jsob, _ := GetJsObject([]byte(str))
	fmt.Println(jsob["age"].GetAsInt())
}

func TestGetJsObjectByReader(t *testing.T) {
	file, err := os.Open("./xx.json")
	defer file.Close()
	if err != nil {
		t.Error("读取文件失败")
	}
	jsob, err := GetJsObjectByReader(file)
	if err != nil {
		t.Error("解析失败")
	}
	fmt.Println(jsob)
}

func TestGetJsObjectByBufReader(t *testing.T) {
	file, err := os.Open("./xx.json")
	defer file.Close()
	if err != nil {
		t.Error("读取文件失败")
	}
	fileReader := bufio.NewReader(file)
	jsob, _ := GetJsObjectByBufReader(fileReader)
	fmt.Println(jsob)
}

func TestGet(t *testing.T) {
	bytes, e := ioutil.ReadFile("./yy.json")
	if e != nil {
		t.Error("读取失败")
	}
	//fmt.Println(string(bytes))
	var v *Value = Get(bytes, "contributor")
	fmt.Println(v.GetAsSliceIgnore()[0].GetAsStringIgnore())
	fmt.Println(Get(bytes, "version").GetAsStringIgnore())
	v = Get(bytes, "date.year")
	fmt.Println(v.GetAsIntIgnore())
	jso, err := GetJsObject(bytes)
	if err != nil {
		t.Error("测试失败")
	}
	fmt.Println(GetFromJsObject(jso,"date.month").GetAsIntIgnore())
	fmt.Println(GetFromJsObject(jso, "date.day").GetAsIntIgnore())
	fmt.Println(GetFromJsObject(jso,"date.datetime").GetAsStringIgnore())
	fmt.Println(GetFromJsObject(jso, "key11").GetAsSliceIgnore()[0].GetAsSliceIgnore()[3])
	fmt.Println(GetFromJsObject(jso, "key12").GetAsSliceIgnore()[1].GetAsObjectIgnore()["key14"])
}
