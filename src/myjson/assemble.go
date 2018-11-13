package myjson

//终于到组装的底部了，可能零件还需要改动，但总的来说还是没问题的
var s *StackAnaly
var stack *Stack

func init() {
	s = NewStackAnaly()
	stack = NewStack()

}


//先处理一个字符
func delChar(b byte) {
	if !s.IsEmpty() {
		//上一个元素的状态
		var sign *Sign = s.IsSign()
		if sign == nil {
			//如果sign为nil，那么sign就是key or value的一部分

		}

		if sign.GetStatus() == StaCloBrace {
			//接下来的东西在对象中

		}

		if sign.GetStatus() == StaSquare {
			//接下来的东西在数组中
		}

		if sign.GetStatus() == StaQuotation {
			//接下来的东西应该在“中，所以要么是字符串value，要么是key

		}
	}
	s.Push(b)

}



