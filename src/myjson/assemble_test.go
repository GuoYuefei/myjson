package myjson

import (
	"fmt"
	"io/ioutil"
	"testing"
)

//测试

func TestAssemble1(t *testing.T) {
	//str := "{\"name\":\"gyf\",\"age\":\"12\"}"
	str1 := "{\"name\":\"gyf\",\"age\":\"12\",\"ids\":{\"id1\":\"1234567890\",\"id2\":\"2\"}}"
	var bs []byte = make([]byte,0,20)
	bs = append(bs, []byte(str1)...)
	for _, b := range bs {
		delChar(b)
		//fmt.Println(keyStr)
	}
	fmt.Println(s.State.Top().GetAsObjectIgnore())
	fmt.Println(s.State.Top().GetAsObjectIgnore()["ids"].GetAsObjectIgnore())
	//这里可以发现应该是ids的地方变成了id2 //ps已修复
	fmt.Println(s.State.Top().GetAsObjectIgnore()["ids"].GetAsObjectIgnore()["id1"].GetAsStringIgnore())
	fmt.Println(s.State.Top().GetAsObjectIgnore()["name"].GetAsStringIgnore())
	fmt.Println(s.size(),len(str1))
	//fmt.Println(s.Pop(),s.Pop(),s.Pop(),s.Pop(),s.Pop())			//{--123  }--125没单出s 58--:
	//s.Push([]byte("\"")[0])
	//fmt.Println(s.getFlag() & 0x40)
	//delChar([]byte("{")[0])
	//delChar([]byte("\"")[0])
}

func TestAssemble2(t *testing.T) {
	//,\"age\":\"12\"
	str2 := "{\"name\":\"gyf\",\"age\":\"12\",\"ids\":[\"1234567890\",\"44\"]}"
	clearAllStack()
	var bs []byte = make([]byte,0,20)
	bs = append(bs, []byte(str2)...)
	for _, b := range bs {
		delChar(b)
		//fmt.Println(keyStr)
	}
	fmt.Println(s.State.Top().GetAsObjectIgnore())
	fmt.Println(s.State.Top().GetAsObjectIgnore()["ids"].GetAsSliceIgnore()[0].GetAsStringIgnore())

	fmt.Println(s.data,s.size(),len(str2))
}

func TestAssemble3(t *testing.T) {
	str3 := "{\"name\":\"gyf\",\"age\":12,\"sex\":true,\"abc\":null,\"ids\":[\"33\",\"44\"],\"ids1\":{\"id1\":32,\"id2\":111}}"
	clearAllStack()
	var bs []byte = make([]byte,0,20)
	bs = append(bs, []byte(str3)...)
	for _, b := range bs {
		delChar(b)
		//fmt.Println(keyStr)
	}
	fmt.Println(s.State.Top().GetAsObjectIgnore())
	fmt.Println(s.State.Top().GetAsObjectIgnore()["age"].GetAsInt())
	fmt.Println(s.data,s.size(),len(str3))
	v, err := s.State.Top().GetAsObjectIgnore()["sex"].GetAsBool()
	if err != nil {
		t.Error(err)
	}
	if v {
		fmt.Println(v)
	}
	fmt.Println(s.State.Top().GetAsObjectIgnore()["abc"].IsNull())
}

//测试字符串来自文件
func TestAssembleFile(t *testing.T) {
	clearAllStack()
	bs, err := ioutil.ReadFile("./xx.json")
	if err != nil {
		t.Error("读取文件失败")
		panic(err)
	}
	//bs = CompressJson(bs)
	//fmt.Println(string(bs))

	for _, b := range bs {
		delChar(b)
	}
	fmt.Println(s.State.Top().GetAsObjectIgnore())
	fmt.Println("key3 = ",s.State.Top().GetAsObjectIgnore()["key3"].GetAsSliceIgnore())
	fmt.Println("key4.key5 = ", s.State.Top().GetAsObjectIgnore()["key4"].GetAsObjectIgnore()["key5"].GetAsIntIgnore())
	//fmt.Println(s.State.Top().GetAsObjectIgnore()["key11"])
	fmt.Println(s.State.Top().GetAsObjectIgnore()["key11"].GetAsSliceIgnore()[1].GetAsSliceIgnore()[1].GetAsInt())
	fmt.Println(s.State.Top().GetAsObjectIgnore()["key11"].GetAsSliceIgnore()[0].GetAsSliceIgnore()[3].GetAsStringIgnore())
	fmt.Println("key12 = ", s.State.Top().GetAsObjectIgnore()["key12"].GetAsSliceIgnore()[0].GetAsObjectIgnore()["key13"].GetAsIntIgnore())

	fmt.Println(s.State.Top().GetAsObjectIgnore()["key6"].GetAsObjectIgnore()["key8"].GetAsObjectIgnore()["key9"].GetAsStringIgnore())
	fmt.Println(s.data,s.size(),len(string(bs)))
	fmt.Println(sState.data,s.size())
	fmt.Println(keyStrs.data,keyStrs.size())
}
