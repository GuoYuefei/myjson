package myjson


//----------这个是jsonValue.go需要定义的简单结构---------------
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


//----------这个是token.go需要定义的数据-------------------------
//忽然返现定义token前需要先定义状态
type Status int

const(
	//已经绝望的状态  啥都不期待了
	StaNone Status = iota
	//期待key
	StaKey
	//期待Value
	StaValue
	//期待大括号
	StaCloBrace
	//期待双引号
	StaQuotation
	//期待中括号
	StaSquare

)

//这个是包内全局变量，用于记录“\"”是否被期待，因为这个部分前后，想{}[]是不同字符成对出现的
//外部禁止访问，这样会导致不安全
var flagQuo = false





