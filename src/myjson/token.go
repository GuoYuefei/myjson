package myjson

//这是一个关于token的包
//先定义其数据结构，因为其实json是树形结构的，所以token其实有点像树的节点


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
type Sign struct {
	//标识符一般是有自己的期待值的，比如说”{“是会期待”}“。 ","是期待key的出现
	status Status

}

func NewSign(sta Status) *Sign {
	return &Sign{sta}
}

func (s *Sign)GetStatus() Status {
	return s.status
}

//根据sign这个字符来判定是否是Sign类型，返回与标志符想对应的*Sign
func GetSign(sign string) *Sign {
	var s *Sign
	switch sign {
	case "{": s = NewSign(StaCloBrace)
	case "[": s = NewSign(StaSquare)
	case "\"":
		if flagQuo {
			s = NewSign(StaQuotation)
		} else {
			s = NewSign(StaNone)
		}
		flagQuo = !flagQuo			//遇到一次”就取反一次
	case ":": s = NewSign(StaValue)
	case ",": s = NewSign(StaKey)
	case "}","]": s = NewSign(StaNone)
	default:
		s = &Sign{-1}
	}
	return s
}



type Token struct {
	//前节点，这是一个双亲表示法
	parent *Token
	//当前节点的值
	value Value
	//当前节点的状态，这个状态取决于前一个关键字也就是前一个Sign对象，特别注意{ “ ： ，[
	status Status
}


func (t *Token)GetParent() *Token {
	return t.parent
}

func (t *Token)GetValue() *Value {
	return &t.value
}

func (t *Token)GetStatus() Status {
	return t.status
}
