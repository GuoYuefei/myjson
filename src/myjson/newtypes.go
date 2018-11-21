package myjson

//----------这个是jsonValue.go需要定义的简单结构---------------
type NumberType int

//Array, Object, True, False as a special number type
const (
	Null NumberType = iota
	Uint
	Int
	Uint8
	Int8
	Float64
	Array
	Object
	Bool
	Str
)

// Define an alias, this type is a js object, in fact, is what the js object should look like
// There is a point to determine that the value of the map must be of type Value. In fact,
// you can think of this type as a recursive structure.
type JsObject = map[string]*Value

// There is a point to determine that the data in the slice must be of type Value
// it can be thought of as a recursive structure.
type Slice = []*Value

//----------这个是token.go需要定义的数据-------------------------

//当前字符期待哪个字符的出现
//A status flag required to analyze a node, indicating which character the current node value expects to appear
type Status = int

const (
	//已经绝望的状态  啥都不期待了
	StaNone Status = iota
	//expects key
	StaKey
	//expects Value
	StaValue
	//expects }
	StaCloBrace
	//expects "
	StaQuotation
	//expects ]
	StaSquare
)



// As a unique identifier for a keyword, you can distinguish {}[]:,"
type WhatsType = int
const (
	TNone WhatsType = iota
	//左大括号
	TBracesL
	//右大
	TBracesR
	//中左
	TSquareL
	//中右括号
	TSquareR
	//引号
	TQuotation
	//冒号
	TColon
	//逗号
	TComma
)