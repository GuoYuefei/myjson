package myjson

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//书写示例文件


func Example() {
	file, _ := os.Open("yy.json")
	jso, err := GetJsObjectByReader(file)

	if err != nil {
		//do something
	}

	v := GetFromJsObject(jso, "date.time.hour")
	fmt.Print(v.GetAsIntIgnore())

	//or

	bs, _ := ioutil.ReadFile("yy.json")
	v = Get(bs, "date.time.minutes")
	fmt.Println(" :", v.GetAsIntIgnore())

	// Output:
	// 14 : 26
}

func ExampleGetJsObject() {
	str := "{\"name\": \"solomon\",\"Addr\": 3304,\"Array\":[1,2,3,4],\"ob\": {\"sex\":true,\"cc\":\"cc\"}}"
	jso, err := GetJsObject([]byte(str))
	if err != nil {
		//do something
	}
	fmt.Println(jso["name"].GetAsStringIgnore())

	// Output:
	// solomon
}

func ExampleGetJsObjectByReader() {
	file, _ := os.Open("xx.json")
	jso, err := GetJsObjectByReader(file)
	if err != nil {
		//do something
	}
	fmt.Println(jso["key1"].GetAsIntIgnore())

	// Output:
	// 12
}

func ExampleGetJsObjectByBufReader() {
	file, _ := os.Open("xx.json")
	buf := bufio.NewReader(file)
	jso, err := GetJsObjectByBufReader(buf)
	if err != nil {
		//do something
	}
	fmt.Println(jso["key1"].GetAsIntIgnore())

	// Output:
	// 12
}

// It is recommended to use this function more.
func ExampleGetFromJsObject() {
	str := "{\"name\": \"solomon\",\"Addr\": 3304,\"Array\":[1,2,3,4],\"ob\": {\"sex\":true,\"cc\":\"cc\"}}"
	jso, err := GetJsObject([]byte(str))
	if err != nil {
		//do something
	}
	rst := GetFromJsObject(jso, "ob.sex")
	if rst.IsBool() {
		fmt.Println(rst.GetAsBoolIgnore())
	}

	// Output:
	// true
}

func ExampleGet() {
	str := "{\"name\": \"solomon\",\"Addr\": 3304,\"Array\":[1,2,3,4],\"ob\": {\"sex\":true,\"cc\":\"cc\"}}"
	value := Get([]byte(str), "ob.cc")
	fmt.Println(value.GetAsStringIgnore())

	// Output:
	// cc
}


