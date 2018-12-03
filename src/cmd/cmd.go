package main

import (
	"fmt"
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
	fmt.Println(strings.SplitAfter("{    \"name\"  :\"qqq\"  }","\""))
	dd()
}

type vv struct {
	a interface{}
}

func (v *vv) setV(a interface{}) {
	v.a = a
}

func (v *vv) getvvSlice() []*vv {
	if a, ok := v.a.([]*vv); ok {
		return a
	}
	return nil
}


type ss struct {
	Data []*vv
}

func dd() {
	//v := &vv{1}
	vs := make([]*vv, 0, 10)
	vs = append(vs, &vv{123})
	//将切片分装成vv类型
	vvs := &vv{vs}

	s := &ss{make([]*vv, 0, 12)}
	s.Data = append(s.Data, vvs)
	fmt.Printf("%p\n", s.Data[0].a)
	fmt.Println(s.Data[0].getvvSlice())
	vvvs := s.Data[0].getvvSlice()
	fmt.Printf("vvvs.poriner1: %p \n", vvvs)
	vvvs = append(vvvs, &vv{2134214})
	vvvs[0] = &vv{23232}
	fmt.Println("vvvs:", vvvs)
	fmt.Printf("vvvs.poriner2: %p \n", vvvs)
	fmt.Printf("%p\n", s.Data[0].a)
	fmt.Println(s.Data[0].getvvSlice())
	fmt.Println(s.Data[0].getvvSlice()[0])
}

