package myjson

//这是一个关于token的包
//先定义其数据结构，因为其实json是树形结构的，所以token其实有点像树的节点

type Sign struct {
	//标识符一般是有自己的期待值的，比如说”{“是会期待”}“。 ","是期待key的出现
	status Status
}

func NewSign(sta Status) *Sign {
	return &Sign{sta}
}

func (s *Sign) GetStatus() Status {
	return s.status
}

//根据sign这个字符来判定是否是Sign类型，返回与标志符想对应的*Sign
//StaNone最终会是继承上一个期待的,
func GetSign(sign []byte) *Sign {
	var s *Sign
	switch string(sign) {
	case "{":
		s = NewSign(StaCloBrace)
	case "[":
		s = NewSign(StaSquare)
	case "\"":
		if flagQuo {
			s = NewSign(StaQuotation)
		} else {
			s = NewSign(StaNone)
		}
		flagQuo = !flagQuo //遇到一次”就取反一次
	case ":":
		s = NewSign(StaValue)
	case ",":			//这个需要修改，如果在[]里就不一样了
		s = NewSign(StaKey)
	case "}", "]":
		s = NewSign(StaNone)
	default:
		s = nil
	}
	return s
}

type Token struct {
	//前节点，这是一个双亲表示法
	parent *Token
	//当前节点的值
	value *Value
	//当前节点的状态，这个状态取决于前一个关键字也就是前一个Sign对象，特别注意{ “ ： ，[
	status Status
}

func NewToken(p *Token, v *Value, s Status) *Token {
	return &Token{p, v, s}
}

func (t *Token) GetParent() *Token {
	return t.parent
}

func (t *Token) GetValue() *Value {
	return t.value
}

func (t *Token) GetStatus() Status {
	return t.status
}

//还需要一个栈，这个栈需要自定义一些方法，所以感觉可以重开一个文件
