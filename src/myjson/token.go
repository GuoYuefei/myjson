package myjson

//这是一个关于token的包
//先定义其数据结构，因为其实json是树形结构的，所以token其实有点像树的节点

type Sign struct {
	//标识符一般是有自己的期待值的，比如说”{“是会期待”}“。 ","是期待key的出现
	status Status
	v []byte
	wt WhatsType
}

//KEY OR VALUE 出栈时使用
func NewSign(sta Status, v []byte, whatsType WhatsType) *Sign {
	return &Sign{sta, v, whatsType}
}

func newSign2(sta Status, whatsType WhatsType) *Sign {
	return &Sign{sta, make([]byte,0,1), whatsType}
}

func (s *Sign) setV(v []byte) {
	s.v = v
}

//空状态的应该在栈顶时设置状态，状态为前节点状态
func (s *Sign) SetSta(sta Status) {
	s.status = sta
}

func (s *Sign) GetWT() WhatsType {
	return s.wt
}

func (s *Sign) GetV() []byte {
	return s.v
}

func (s *Sign) GetStatus() Status {
	return s.status
}

//根据sign这个字符来判定是否是Sign类型，返回与标志符想对应的*Sign
//StaNone最终会是继承上一个期待的,
func GetSign(sign []byte,flag byte) *Sign {
	var s *Sign
	switch string(sign) {
	case "{":
		s = newSign2(StaCloBrace, TBracesL)
	case "[":
		s = newSign2(StaSquare, TSquareL)
	case "\"":
		if flag & 0x40 == 0x40 {		//当栈中有奇数个引号的时候
			s = newSign2(StaQuotation, TQuotation)
		} else {
			s = newSign2(StaNone, TQuotation)
		}
		flagQuo = !flagQuo //遇到一次”就取反一次
	case ":":
		s = newSign2(StaValue, TColon)
	case ",":
		if flag & 0x80 == 0x80 {		//当在数组里的时候“，”期待的是下一个数组元素
			s = newSign2(StaValue, TComma)
		}
		s = newSign2(StaKey, TComma)
	case "}":
		s = newSign2(StaNone, TBracesR)
	case "]":
		s = newSign2(StaNone, TSquareL)
	default:
		s = nil
	}
	if s != nil {
		s.setV(sign)
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
