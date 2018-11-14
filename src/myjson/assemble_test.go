package myjson

import (
	"fmt"
	"testing"
)

//测试

func TestAssemble1(t *testing.T) {
	//str := "{\"name\":\"gyf\",\"age\":\"12\"}"
	str1 := "{\"name\":\"gyf\",\"age\":\"12\",\"ids\":{\"id1\":\"1\",\"id2\":\"2\"}}"
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
	fmt.Println(s.Size(),len(str1))
	fmt.Println(s.Pop(),s.Pop(),s.Pop(),s.Pop(),s.Pop())			//{--123  }--125没单出s 58--:
	//s.Push([]byte("\"")[0])
	//fmt.Println(s.GetFlag() & 0x40)
	//delChar([]byte("{")[0])
	//delChar([]byte("\"")[0])
}

func TestAssemble2(t *testing.T) {
	//,\"age\":\"12\"
	str2 := "{\"name\":\"gyf\",\"age\":\"12\",\"ids\":[\"33\",\"44\"]}"
	var bs []byte = make([]byte,0,20)
	bs = append(bs, []byte(str2)...)
	for _, b := range bs {
		delChar(b)
		//fmt.Println(keyStr)
	}
	fmt.Println(s.State.Top().GetAsObjectIgnore())
	fmt.Println(s.State.Top().GetAsObjectIgnore()["ids"].GetAsSliceIgnore()[0].GetAsStringIgnore())

	fmt.Println(s.data,s.Size(),len(str2))
}