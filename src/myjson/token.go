package myjson

//这是一个关于token的包
//先定义其数据结构，因为其实json是树形结构的，所以token其实有点像树的节点


//之前优化简化，后来废弃
//已废弃
type sign struct {
	//标识符一般是有自己的期待值的，比如说”{“是会期待”}“。 ","是期待key的出现
	//status Status
	//v byte
	wt WhatsType
}

//KEY OR VALUE 出栈时使用
func newSign(whatsType WhatsType) *sign {
return &sign{ whatsType}
}

func newSign2(whatsType WhatsType) *sign {
	return &sign{whatsType}
}

//func (s *Sign) setV(v byte) {
//	s.v = v
//}

//空状态的应该在栈顶时设置状态，状态为前节点状态
//func (s *Sign) SetSta(sta Status) {
//	s.status = sta
//}

func (s *sign) getWT() WhatsType {
	return s.wt
}

//func (s *Sign) GetV() byte {
//	return s.v
//}

//func (s *Sign) GetStatus() Status {
//	return s.status
//}

//根据sign这个字符来判定是否是Sign类型，返回与标志符想对应的*Sign
//StaNone最终会是继承上一个期待的,
func GetSign(b byte) *sign {
	var s *sign
	switch b {
	//出现频率应该是最高的放在前面
	case ',':
		//if flag & 0x80 == 0x80 {		//当在数组里的时候“，”期待的是下一个数组元素
		//	s = newSign2(StaValue, TComma)
		//} else {
		//	s = newSign2(StaKey, TComma)
		//}
		s = newSign2(TComma)
	case ':':
		s = newSign2(TColon)
	case '"':
		//if flag & 0x40 == 0x40 {		//当栈中有奇数个引号的时候
		//	s = newSign2(StaQuotation, TQuotation)
		//} else {
		//	s = newSign2(StaNone, TQuotation)
		//}
		s = newSign2(TQuotation)
	case '{':
		s = newSign2(TBracesL)
	case '[':
		s = newSign2(TSquareL)
	case '}':
		s = newSign2(TBracesR)
	case ']':
		s = newSign2(TSquareR)
	default:
		s = newSign2(TNone)
	}
	//s.setV(sign)
	return s
}



//还需要一个栈，这个栈需要自定义一些方法，所以感觉可以重开一个文件
