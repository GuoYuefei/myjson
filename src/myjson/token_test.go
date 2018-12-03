package myjson

import "testing"

func TestGetSign(t *testing.T) {
	var sign *sign = getSign(',')
	if TComma != sign.getWT() {
		t.Error("测试不通过", i)
	}
}
