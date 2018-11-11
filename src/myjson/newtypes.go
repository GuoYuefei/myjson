package myjson

type NumberType int

const(
	Null NumberType = iota
	Uint
	Int
	Uint8
	Int8
	Float64
	Array
	Object
)

//定义一个别名，这个类型是一个js的对象，实际上就是js对象该有的样子
//有一个点可以确定map的value必须是Value类型的，从外层看，这是一个典型的递归结构
type JsObject = map[string]*Value

//有一个点可以确定，切片中的数据必须是Value类型的，这是一个递归的结构
type Slice = []*Value