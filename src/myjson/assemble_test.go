package myjson

import (
	"fmt"
	"testing"
)

//测试

func TestAssemble(t *testing.T) {
	str := "{\"name\":\"gyf\",\"age\":\"12\"}"
	var bs []byte = make([]byte,0,20)
	bs = append(bs, []byte(str)...)
	for _, b := range bs {
		delChar(b)
		//fmt.Println(keyStr)
	}
	fmt.Println(s.State.Top().GetAsObjectIgnore()["name"].GetAsStringIgnore())
	fmt.Println(s.Size(),len(str))
	fmt.Println(s.Pop(),s.Pop())			//{123}125没单出s
	//s.Push([]byte("\"")[0])
	//fmt.Println(s.GetFlag() & 0x40)
	//delChar([]byte("{")[0])
	//delChar([]byte("\"")[0])
}
