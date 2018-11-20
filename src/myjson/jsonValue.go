package myjson

import (
	"errors"
)

//json是js的对象的序列化
//值应该与js对应
// js.string = go.string
// js.number = go.float64...
// js.null = go.nil
// js.bool = go.bool
// js.array = go.array
// js.object = go.map / object
type Value struct {
	//把value封装起来
	value interface{}
}

// Returns a *Value type representing an empty type
func NewValue() *Value {
	return &Value{nil}
}

// Provide an arbitrary type of parameter 'a'
// Returns a *type type representing Returns a *type type representing the value of any type 'a'
func NewVal(a interface{}) *Value {
	return &Value{a}
}

// Provide an arbitrary type of parameter 'a'
// Then v.value = a
// mean 'v' now starts to represent the value of 'a'
func (v *Value) SetValue(a interface{}) {
	v.value = a
}

//针对string类型，有下类方法

// Determine whether v represents a variable of type string
// Return true if it is a string type, otherwise return false
func (v *Value) IsString() bool {
	if _, ok := v.value.(string); ok {
		return true
	}
	return false
}

//赋值一个string变量
//Receive a parameter of type string
//The variable v will be assigned the value of the parameter
func (v *Value) SetAsString(value string) {
	v.value = value
}

//值作为string返回，因为可能存在该值不是string的情况，所以第二个返回值用于返回是否发生错误
// The value is returned as a string, because there may be cases where the value is not a string,
// so the second return value is used to return whether an error has occurred.
func (v *Value) GetAsString() (string, error) {
	if v, ok := v.value.(string); ok {
		return v, nil
	}
	return "", errors.New("not a string value")
}

//值作为string类型返回，并且忽略错误的发生。若发生错误则返回空字符串
//这个方法是为了方便连续调用
// The value is returned as a string type and the occurrence of the error is ignored. Return an empty string if an error occurs
// This method is to facilitate continuous calls
func (v *Value) GetAsStringIgnore() string {
	if v, ok := v.value.(string); ok {
		return v
	}
	return ""
}

//针对number类型 我本来想针对所有的数字类型的，但是实在要写很多重复的东西了，所以决定还是去掉一些几乎不用的吧

//Determine whether v represents a variable of type Number
//Return true if it is a Number type, otherwise return false
//Only int, uint, int8, uint8(byte), float64 types are supported here.
func (v *Value) IsNumber() bool {
	switch v.value.(type) {
	case int, uint, int8, uint8, float64:
		return true
	default:
		return false
	}
}

//返回类型Number的具体类型，当值v不是Number类型的时候返回第二个返回值返回false，否则第二返回值返回true
//第一个返回值返回该具体数字类型
// Returns the specific type of the type Number, when the value v is not the Number type,
// return the second return value returns false, otherwise the second return value returns true
// The first return value returns the specific number type
func (v *Value) WhichNumType() (NumberType, bool) {
	var result NumberType = Null
	switch v.value.(type) {
	case uint8:
		result = Uint8
	case uint:
		result = Uint
	case int8:
		result = Int8
	case int:
		result = Int
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

//参数是需要判定的类型，如果参数所提供的类型就是值v的类型则返回true，否则返回false
// The parameter is provided with a type tag,
// Return true if the type provided by the argument is the type of the value v, otherwise return false
func (v *Value) IsWhichNumType(numberType NumberType) bool {
	if nt, ok := v.WhichNumType(); ok && nt == numberType {
		return true
	}
	return false
}

// Determine whether v represents a variable of type uint
// Return true if it is a Uint type, otherwise return false
func (v *Value) IsUint() bool {
	return v.IsWhichNumType(Uint)
}

// Determine whether v represents a variable of type int
// Return true if it is a int type, otherwise return false
func (v *Value) IsInt() bool {
	return v.IsWhichNumType(Int)
}

// Determine whether v represents a variable of type byte
// Return true if it is a byte type, otherwise return false
func (v *Value) IsUint8() bool {
	return v.IsWhichNumType(Uint8)
}

// Determine whether v represents a variable of type int8
// Return true if it is a int8 type, otherwise return false
func (v *Value) IsInt8() bool {
	return v.IsWhichNumType(Int8)
}

// Determine whether v represents a variable of type bool
// Return true if it is a bool type, otherwise return false
func (v *Value) IsFloat64() bool {
	return v.IsWhichNumType(Float64)
}

// Because the user-inputted a may not be a numeric type, there may be an assignment error.
// Only accept five types uint int int8 uint8 float64
func (v *Value) SetAsNumber(a interface{}) error {
	var value = NewValue()
	value.SetValue(a)
	if _, ok := value.WhichNumType(); ok {
		v.value = a
		return nil
	}
	return errors.New("a not number type")
}

// Assign a value to the variable v, the assignment type is float64
func (v *Value) SetAsFloat64(f float64) {
	v.value = f
}

// Assign a value to the variable v, the assignment type is int
func (v *Value) SetAsInt(I int) {
	v.value = I
}

// Assign a value to the variable v, the assignment type is uint
func (v *Value) SetAsUint(I uint) {
	v.value = I
}

// Assign a value to the variable v, the assignment type is int8
func (v *Value) SetAsInt8(I int8) {
	v.value = I
}

// Assign a value to the variable v, the assignment type is byte
func (v *Value) SetAsUint8(I uint8) {
	v.value = I
}

// If the variable v can be returned as an int type, the value is returned as an int,
// and the second error return value returns nil
// Returns 0 if the int type cannot be returned, and the second return value returns the corresponding error.
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

// If the variable v can be returned as an int type, the value is returned as an int,
// otherwise the largest integer is returned.
func (v *Value) GetAsIntIgnore() (result int) {
	switch a := v.value.(type) {
	case int:
		result = a
	case int8:
		result = int(a)
	case uint8:
		result = int(a)
	default:
		result = int(^uint(0) >> 1) //返回一个最大值
	}
	return
}

// If the variable v can be returned as an uint type, the value is returned as an uint,
// and the second error return value returns nil
// Returns 0 if the uint type cannot be returned, and the second return value returns the corresponding error.
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

// If the variable v can be returned as an uint type, the value is returned as an uint,
// otherwise the largest unsigned integer is returned.
func (v *Value) GetAsUintIgnore() (result uint) {
	switch a := v.value.(type) {
	case uint:
		result = a
	case uint8:
		result = uint(a)
	default:
		result = ^uint(0) >> 1 //无符号最大值
	}
	return
}

// If the variable v can be returned as an int8 type, the value is returned as an int8,
// and the second error return value returns nil
// Returns 0 if the int8 type cannot be returned, and the second return value returns the corresponding error.
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

// If the variable v can be returned as an byte type, the value is returned as an byte,
// and the second error return value returns nil
// Returns 0 if the byte type cannot be returned, and the second return value returns the corresponding error.
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

// If the variable v can be returned as an float64 type, the value is returned as an float64,
// and the second error return value returns nil
// Returns 0 if the float64 type cannot be returned, and the second return value returns the corresponding error.
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

// If the variable v can be returned as an float64 type, the value is returned as an float64,
// otherwise the largest unsigned integer will be converted to a float64 type return
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
			result = float64(int(^uint(0) >> 1)) //返回一个int下的最大整数吧
		}
	}
	return
}

// end of number

// begin null

// if v.value equal nil ,return true
func (v *Value) IsNull() bool {
	if v.value == nil {
		return true
	}
	return false
}

// Set the value v to nil
func (v *Value) SetNull() {
	v.value = nil
}

// begin bool

// Return true if the value v can represent a boolean type, otherwise return false
func (v *Value) IsBool() bool {
	_, ok := v.value.(bool)
	return ok
}

// If the value v can represent a Boolean type then the first return value returns the value represented by v,
// and the second error return value returns nil
func (v *Value) GetAsBool() (result bool, err error) {
	if a, ok := v.value.(bool); ok {
		return a, nil
	}
	return false, errors.New("can not get a bool type")
}

// set the value v to b
func (v *Value) SetAsBool(b bool) {
	if b {
		v.value = true
	} else {
		v.value = false
	}
}

//end of bool

//begin Object

// set the value v to a
func (v *Value) SetAsObject(a JsObject) {
	v.value = a
}

// Determine whether v represents a variable of type Object
// Return true if it is a Object type, otherwise return false
func (v *Value) IsObject() bool {
	_, ok := v.value.(JsObject)
	return ok
}

// If the value v can represent a JsObject type then the first return value returns the value represented by v,
// and the second error return value returns nil
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

// User-friendly, this method ignores error
// If the variable v can be returned as an JsObject type, the value is returned as an JsObject,
// otherwise nil will return
func (v *Value) GetAsObjectIgnore() JsObject {
	if a, ok := v.value.(JsObject); ok {
		return a
	}
	return nil
}

//end of Object

//begin Array = Slice
//Determine whether v represents a variable of type Slice
//Return true if it is a Slice type, otherwise return false
func (v *Value) IsSlice() bool {
	_, ok := v.value.([]*Value)
	return ok
}

// If the value v can represent a []*Value type then the first return value returns the value represented by v,
// and the second error return value returns nil
func (v *Value) GetAsSlice() (result Slice, err error) {
	err = nil
	if a, ok := v.value.([]*Value); ok {
		result = a
	} else {
		result = nil
		err = errors.New("can not get a Slice Type")
	}
	return
}

// User-friendly, this method ignores error
// If the variable v can be returned as an []*Value type, the value is returned as an []*Value,
// otherwise nil will return
func (v *Value) GetAsSliceIgnore() Slice {
	if a, ok := v.value.([]*Value); ok {
		return a
	}
	return nil
}

// Set the value v to s
func (v *Value) SetAsSlice(s Slice) {
	v.value = s
}

//end of array = slice

//end of Value type
