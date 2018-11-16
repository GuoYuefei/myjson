package main

import "fmt"

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
}
