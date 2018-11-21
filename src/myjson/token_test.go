package myjson

import "testing"

func TestGetSign(t *testing.T) {
	var sign *Sign = GetSign(',',12)
	if StaKey != sign.GetStatus() {
		t.Error("测试不通过", i)
	}
}
