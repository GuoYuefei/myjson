package myjson

import "testing"

func TestGetSign(t *testing.T) {
	var sign *Sign = GetSign(',',12)
	if TComma != sign.GetWT() {
		t.Error("测试不通过", i)
	}
}
