package main

import (
	"fmt"
	"myjson"
	"strings"
)

func main() {
	//v := myjson.NewValue()
	//v.SetValue(123)
	//fmt.Println(v.IsWhichNumType(myjson.Int))
	//fmt.Println(v.IsWhichNumType(myjson.Uint))
	//fmt.Println(v.IsInt())
	//fmt.Println(v.IsWhichNumType(myjson.Float64))
	//v.SetAsString("123")
	//fmt.Println(v.GetAsString())
	//fmt.Println(v.IsString())
	//fmt.Println(v.IsWhichNumType(myjson.Uint))
	s := make([]byte, 0, 2)
	//s[0] = 0
	//s[1] = 1
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(cap(s), len(s), s)
	fmt.Println(s[1:2])
	fmt.Println(int(^uint(0)>>1))
	fmt.Println(myjson.ParseParam("name.age"))
	fmt.Println(strings.SplitAfter("{    \"name\"  :\"qqq\"  }","\""))
	dd()
}

type vv struct {
	A int
}

type ss struct {
	Data *[]*vv
}

func dd() {
	v := &vv{1}
	vs := make([]*vv,0,10)
	s := &ss{&vs}
	*s.Data = append(*s.Data, v)
	fmt.Println((*(s.Data))[0].A)
}
