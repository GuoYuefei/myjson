package myjson

import (
	"errors"
)

//json是js的对象的序列化
//值应该与js对应
//js.string = go.string
//js.number = go.float64
//js.null = go.nil
//js.bool = go.bool
//js.array = go.array
//js.object = go.map or object

type Value struct {
	//把value封装起来
	value interface{}
}

func NewValue() *Value {
	return &Value{nil}
}

func NewVal(a interface{}) *Value {
	return &Value{a}
}

func (v *Value) SetValue(a interface{}) {
	v.value = a
}

//针对string类型，有下类方法

//当是string类型时返回true否则为false
func (v *Value) IsString() bool {
	if _, ok := v.value.(string); ok {
		return true
	}
	return false
}

//赋值一个string变量
func (v *Value) SetAsString(value string) {
	v.value = value
}

//值作为string返回，因为可能存在该值不是string的情况，所以第二个返回值用于返回是否发生错误
func (v *Value) GetAsString() (string, error) {
	if v, ok := v.value.(string); ok {
		return v, nil
	}
	return "", errors.New("not a string value")
}

func (v *Value) GetAsStringIgnore() string {
	if v, ok := v.value.(string); ok {
		return v
	}
	return ""
}

//针对number类型 我本来想针对所有的数字类型的，但是实在要写很多重复的东西了，所以决定还是去掉一些几乎不用的吧

func (v *Value) IsNumber() bool {
	switch v.value.(type) {
	case int, uint, int8, uint8, float64:
		return true
	default:
		return false
	}
}

func (v *Value) WhichNumType() (NumberType, bool) {
	var result NumberType = Null
	switch v.value.(type) {
	case int:
		result = Int
	case uint:
		result = Uint
	case int8:
		result = Int8
	case uint8:
		result = Uint8
	case float64:
		result = Float64
	//数组类型找不到一个合适的表示方式，但是可以放在defult中
	//case []interface{}: result = Slice
	//这里不妨Object和Array的判定了，还是只判定数字类型比较好
	//case JsObject: result = Object
	default:
		result = Null
	}
	if result == Null {
		return Null, false //如果不是数字就返回false
	}
	return result, true
}

//参数是需要判定的类型，如果v是这个类型正确就返回true
func (v *Value) IsWhichNumType(numberType NumberType) bool {
	if nt, ok := v.WhichNumType(); ok && nt == numberType {
		return true
	}
	return false
}

func (v *Value) IsUint() bool {
	return v.IsWhichNumType(Uint)
}

func (v *Value) IsInt() bool {
	return v.IsWhichNumType(Int)
}

func (v *Value) IsUint8() bool {
	return v.IsWhichNumType(Uint8)
}

func (v *Value) IsInt8() bool {
	return v.IsWhichNumType(Int8)
}

func (v *Value) IsFloat64() bool {
	return v.IsWhichNumType(Float64)
}

//因为用户传入的a可能不是数字类型，所以可能存在赋值错误的情况
//只接受五个类型uint int int8 uint8 float64
func (v *Value) SetAsNumber(a interface{}) error {
	var value = NewValue()
	value.SetValue(a)
	if _, ok := value.WhichNumType(); ok {
		v.value = a
		return nil
	}
	return errors.New("a not number type")
}

func (v *Value) SetAsFloat64(f float64) {
	v.value = f
}

func (v *Value) SetAsInt(I int) {
	v.value = I
}

func (v *Value) SetAsUint(I uint) {
	v.value = I
}

func (v *Value) SetAsInt8(I int8) {
	v.value = I
}

func (v *Value) SetAsUint8(I uint8) {
	v.value = I
}

//会自动将int8，uint8类型转换，float和uint不支持，因为存在危险
func (v *Value) GetAsInt() (result int, err error) {
	err = nil
	switch a := v.value.(type) {
	case int:
		result = a
	case int8:
		result = int(a)
	case uint8:
		result = int(a)
	default:
		result = 0
		err = errors.New("can not get a int number")
	}
	return
}

func (v *Value) GetAsIntIgnore() (result int) {
	switch a := v.value.(type) {
	case int:
		result = a
	case int8:
		result = int(a)
	case uint8:
		result = int(a)
	default:
		result = int(^uint(0) >> 1)		//返回一个最大值
	}
	return
}

func (v *Value) GetAsUint() (result uint, err error) {
	err = nil
	switch a := v.value.(type) {
	case uint:
		result = a
	case uint8:
		result = uint(a)
	default:
		result = 0
		err = errors.New("can not get a uint number")
	}
	return
}

func (v *Value) GetAsUintIgnore() (result uint) {
	switch a := v.value.(type) {
	case uint:
		result = a
	case uint8:
		result = uint(a)
	default:
		result = ^uint(0) >> 1			//无符号最大值
	}
	return
}

func (v *Value) GetAsInt8() (result int8, err error) {
	err = nil
	switch a := v.value.(type) {
	case int8:
		result = a
	default:
		result = 0
		err = errors.New("can not get a int8 number")
	}
	return
}

func (v *Value) GetAsUint8() (result uint8, err error) {
	err = nil
	switch a := v.value.(type) {
	case uint8:
		result = a
	default:
		result = 0
		err = errors.New("can not get a uint8 number")
	}
	return
}

func (v *Value) GetAsFloat64() (result float64, err error) {
	err = nil
	if _, ok := v.WhichNumType(); ok {
		switch a := v.value.(type) {
		case float64:
			result = a
		case uint:
			result = float64(a)
		case int:
			result = float64(a)
		case uint8:
			result = float64(a)
		case int8:
			result = float64(a)
		default:
			result = 0
			err = errors.New("can not get a float64 number")
		}
	}
	return
}

func (v *Value) GetAsFloat64Ignore() (result float64) {
	if _, ok := v.WhichNumType(); ok {
		switch a := v.value.(type) {
		case float64:
			result = a
		case uint:
			result = float64(a)
		case int:
			result = float64(a)
		case uint8:
			result = float64(a)
		case int8:
			result = float64(a)
		default:
			result = float64(int(^uint(0)>>1))			//返回一个int下的最大整数吧
		}
	}
	return
}

//end of number

//begin null

//if v.value equal nil ,return true
func (v *Value) IsNull() bool {
	if v.value == nil {
		return true
	}
	return false
}

func (v *Value) SetNull() {
	v.value = nil
}

//begin bool

func (v *Value) IsBool() bool {
	_, ok := v.value.(bool)
	return ok
}

func (v *Value) GetAsBool() (result bool, err error) {
	if a, ok := v.value.(bool); ok {
		return a, nil
	}
	return false, errors.New("can not get a bool type")
}

func (v *Value) SetAsBool(b bool) {
	if b {
		v.value = true
	} else {
		v.value = false
	}
}

//end of bool

//begin Object
func (v *Value) SetAsObject(a JsObject) {
	v.value = a
}

func (v *Value) IsObject() bool {
	_, ok := v.value.(JsObject)
	return ok
}

func (v *Value) GetAsObject() (result JsObject, err error) {
	err = nil
	if a, ok := v.value.(JsObject); ok {
		result = a
	} else {
		result = nil
		err = errors.New("can not get a JsObject type")
	}
	return
}

//这个是为了方便用户调用，忽略error
//只有容器才提供这样的方法
func (v *Value) GetAsObjectIgnore() JsObject {
	if a, ok := v.value.(JsObject); ok {
		return a
	}
	return nil
}

//end of Object

//begin Array = Slice
func (v *Value) IsSlice() bool {
	_, ok := v.value.([]*Value)
	return ok
}

func (v *Value) GetAsSlice() (result []*Value, err error) {
	err = nil
	if a, ok := v.value.([]*Value); ok {
		result = a
	} else {
		result = nil
		err = errors.New("can not get a Slice Type")
	}
	return
}

//忽略error方便用户调用
//只有容器才提供这样的方法
func (v *Value) GetAsSliceIgnore() []*Value {
	if a, ok := v.value.([]*Value); ok {
		return a
	}
	return nil
}

func (v *Value) SetAsSlice(s []*Value) {
	v.value = s
}

//end of array = slice

//end of Value type
