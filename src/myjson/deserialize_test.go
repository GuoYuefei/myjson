package myjson

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestGetJsObject(t *testing.T) {
	str := "{\"name\":\"solomon\",\"age\":22}"
	jsob, _ := GetJsObject([]byte(str))
	fmt.Println(jsob["age"].GetAsInt())
}

func TestGetJsObjectByReader(t *testing.T) {
	file, err := os.Open("./xx.json")
	defer file.Close()
	if err != nil {
		t.Error("读取文件失败")
	}
	jsob, err := GetJsObjectByReader(file)
	if err != nil {
		t.Error("解析失败")
	}
	fmt.Println(jsob)
}

func TestGetJsObjectByBufReader(t *testing.T) {
	file, err := os.Open("./xx.json")
	defer file.Close()
	if err != nil {
		t.Error("读取文件失败")
	}
	fileReader := bufio.NewReader(file)
	jsob, _ := GetJsObjectByBufReader(fileReader)
	fmt.Println(jsob)
}
