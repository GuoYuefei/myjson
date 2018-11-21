package myjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myjson"
	"testing"
)

func  BenchmarkScannerTestGetJsObjectSmall(b *testing.B) {
	bytes, e := ioutil.ReadFile("./easy.json")
	//bytes := []byte("{\"name\":\"solomon\",\"age\":22,\"HouseholdRegister\":33}")
	if e != nil {
		b.Error("读取文件失败")
	}
	var object myjson.JsObject = make(myjson.JsObject)
	for i := 0; i < b.N; i++ {
		object, _ = myjson.GetJsObject(bytes)
		if e != nil {
			b.Error("测试函数失败")
		}
	}
	b.ReportAllocs()
	fmt.Println(myjson.GetFromJsObject(object, "name"))
}

type Object1 struct {
	Name string
	Age int
	HouseholdRegister int
}

func BenchmarkScannerTestOfficeSmall(b *testing.B) {
	bytes, e := ioutil.ReadFile("./easy.json")
	if e != nil {
		b.Error("读取文件失败")
	}
	var ob = &Object1{}
	for i := 0; i < b.N; i++ {
		e = json.Unmarshal(bytes, ob)
		if e != nil {
			b.Error("测试函数失败")
		}
	}
	b.ReportAllocs()
	fmt.Println(ob)
}

func BenchmarkTestGetJsObjectMid(b *testing.B) {
	bytes, e := ioutil.ReadFile("../testJsonFile/xx.json")
	if e != nil {
		b.Error("读取文件失败")
	}
	var object myjson.JsObject = make(myjson.JsObject)
	for i := 0; i < b.N; i++ {
		object, _ = myjson.GetJsObject(bytes)
		if e != nil {
			b.Error("测试函数失败")
		}
	}
	b.ReportAllocs()
	fmt.Println(myjson.GetFromJsObject(object, "ob.sex"))
}

type Object2 struct {
	Name string
	Addr int
	Array []int
	Ob struct{
		Sex bool
		Cc string
	}
}


func BenchmarkTestOfficeMid(b *testing.B) {
	bytes, e := ioutil.ReadFile("../testJsonFile/xx.json")
	if e != nil {
		b.Error("读取文件失败")
	}
	var ob = &Object2{}
	for i := 0; i < b.N; i++ {
		e = json.Unmarshal(bytes, ob)
		if e != nil {
			b.Error("测试函数失败")
		}
	}
	b.ReportAllocs()
	fmt.Println(ob)
}
