package main

import (
	"fmt"
	"myjson"
)

func main() {
	v := myjson.NewValue()
	v.SetValue(123)
	fmt.Println(v.IsWhichNumType(myjson.Int))
	fmt.Println(v.IsWhichNumType(myjson.Uint))
	fmt.Println(v.IsInt())
	fmt.Println(v.IsWhichNumType(myjson.Float64))
	v.SetAsString("123")
	fmt.Println(v.GetAsString())
	fmt.Println(v.IsString())
	fmt.Println(v.IsWhichNumType(myjson.Uint))
}
