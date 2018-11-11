package myjson

import (
	"fmt"
	"testing"
)

var i int = 0

func TestIsWhichNumType(t *testing.T) {
	var v = NewValue()
	v.SetAsFloat64(123.123)
	if !v.IsWhichNumType(Float64) {
		t.Error("测试失败" ,i)
		i++
	}
	v.SetAsInt8(12)
	if !v.IsWhichNumType(Int8) {
		t.Error("测试失败" ,i)
		i++
	}
	v.SetAsUint8(231)
	if !v.IsWhichNumType(Uint8) {
		t.Error("测试失败" ,i)
		i++
	}
	v.SetAsUint(123243423)
	if !v.IsWhichNumType(Uint) {
		t.Error("测试失败" ,i)
		i++
	}
	v.SetAsInt(2313)
	if !v.IsWhichNumType(Int) {
		t.Error("测试失败" ,i)
		i++
	}
	v.SetValue(23)
	if !v.IsWhichNumType(Int) {
		t.Error("测试失败" ,i)
		i++
	}
	err := v.SetAsNumber("32")
	if err == nil {
		t.Error("测试失败" ,i)
		i++
	}
	err = v.SetAsNumber(2324.343)
	if err != nil {
		t.Error("测试失败" ,i)
		i++
	}
	if !v.IsWhichNumType(Float64) {
		t.Error("测试失败" ,i)
		i++
	}
}

func TestIsTy(t *testing.T) {
	var v *Value = NewValue()
	v.SetAsFloat64(123.23)
	if !v.IsFloat64() {
		t.Error("测试失败" ,i)
		i++
	}
	if !v.IsNumber() {
		t.Error("测试失败" ,i)
		i++
	}
	result, err := v.GetAsFloat64()
	if err != nil {
		t.Error("测试失败" ,i)
		i++
	}else {
		fmt.Println("result", result)
	}
//----------------------------------------------
	v.SetAsInt(123)
	if !v.IsInt() {
		t.Error("测试失败" ,i)
		i++
	}
	if !v.IsNumber() {
		t.Error("测试失败" ,i)
		i++
	}
	result, err = v.GetAsFloat64()
	if err != nil {
		t.Error("测试失败" ,i)
		i++
	}else {
		fmt.Println("result", result)
	}
	result1, err := v.GetAsInt()
	if err != nil {
		t.Error("测试失败" ,i)
		i++
	}else {
		fmt.Println("result1", result1)
	}
//----------------------------------------------
	v.SetAsUint(123)
	if !v.IsUint() {
		t.Error("测试失败" ,i)
		i++
	}
	if !v.IsNumber() {
		t.Error("测试失败" ,i)
		i++
	}
	result, err = v.GetAsFloat64()
	if err != nil {
		t.Error("测试失败" ,i)
		i++
	}else {
		fmt.Println("result", result)
	}
	result4, err := v.GetAsUint()
	if err != nil {
		t.Error("测试失败" ,i)
		i++
	}else {
		fmt.Println("result4", result4)
	}

//---------------------------------------------
	v.SetAsInt8(123)
	if !v.IsInt8() {
		t.Error("测试失败" ,i)
		i++
	}
	if !v.IsNumber() {
		t.Error("测试失败" ,i)
		i++
	}
	result, err = v.GetAsFloat64()
	if err != nil {
		t.Error("测试失败" ,i)
		i++
	}else {
		fmt.Println("result", result)
	}
	result1, err = v.GetAsInt()
	if err != nil {
		t.Error("测试失败" ,i)
		i++
	}else {
		fmt.Println("result1", result1)
	}
	result2, err := v.GetAsInt8()
	if err != nil {
		t.Error("测试失败" ,i)
		i++
	}else {
		fmt.Println("result2", result2)
	}
//--------------------------------------------------
	v.SetAsUint8(23)
	if !v.IsUint8() {
		t.Error("测试失败" ,i)
		i++
	}
	if !v.IsNumber() {
		t.Error("测试失败" ,i)
		i++
	}
	result, err = v.GetAsFloat64()
	if err != nil {
		t.Error("测试失败" ,i)
		i++
	}else {
		fmt.Println("result", result)
	}
	result4, err = v.GetAsUint()
	if err != nil {
		t.Error("测试失败" ,i)
		i++
	}else {
		fmt.Println("result4", result4)
	}
	result3, err := v.GetAsUint8()
	if err != nil {
		t.Error("测试失败" ,i)
		i++
	}else {
		fmt.Println("result3", result3)
	}
}

func TestValue_IsNull(t *testing.T) {
	v := NewValue()
	if !v.IsNull() {
		t.Error("测试失败")
	}
	v.SetAsFloat64(323.23)
	v.SetNull()
	if !v.IsNull() {
		t.Error("测试失败")
	}
}

func TestValue_Bool(t *testing.T) {
	v := NewValue()
	v.SetAsBool(true)
	//第一次测试失败
	if !v.IsBool() {
		t.Error("测试失败")
	}
	result, err := v.GetAsBool()
	if err != nil {
		t.Error("测试失败", v.value)
	}else {
		fmt.Println("result for bool:",result)
	}
}

func TestValue_Object(t *testing.T) {
	ob := make(JsObject)
	ob["12"] = NewVal(123)
	ob["21"] = NewVal("21")
	v := NewValue()
	v.SetAsObject(ob)
	if !v.IsObject() {
		t.Error("test error",ob)
	}
	cc, err := v.GetAsObject()
	if err != nil {
		t.Error("test error",cc)
	}
}

func TestValue_Slice(t *testing.T) {
	sl := make(Slice, 0)
	sl1 := make(Slice,0)
	ob := make(JsObject)
	ob["1"] = NewVal(2312345)
	sl1 = append(sl1, NewVal(32), NewVal(ob))
	sl = append(sl,NewVal(21),NewVal(ob), NewVal(sl1))
	cc := NewValue()
	cc.SetAsSlice(sl)
	//引用变量会改变的哦
	ob["1"] = NewVal(11111111111)
	if !cc.IsSlice() {
		t.Error("test error,",cc.value)
	}
	slll, err := cc.GetAsSlice()
	if err != nil {
		t.Error("test err")
	}
	if obb, err := slll[1].GetAsObject(); err == nil {
		fmt.Println(obb["1"].GetAsInt())
		obb["2"] = NewVal("#2")
	}else {
		t.Error("test err",slll)
	}
	fmt.Println(cc.GetAsSliceIgnore()[1].GetAsObjectIgnore()["2"].GetAsString())

}

func TestValue_String(t *testing.T) {
	v := NewValue()
	v.SetAsString("Hello")
	if !v.IsString() {
		t.Error("test err",v.value)
	}
	if str, err := v.GetAsString(); err == nil {
		fmt.Println(str)
	}else {
		t.Error("test err" )
	}
}
