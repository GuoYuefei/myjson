package main

import (
	"fmt"
	"io/ioutil"
	"myjson"
)

func main() {
	bytes, e := ioutil.ReadFile("./src/myjson/yy.json")
	if e != nil {
		panic(e)
	}
	value := myjson.Get(bytes, "date.time.hour")
	//will see &{14} :  14
	fmt.Println(value, ": ", value.GetAsIntIgnore())
	//or you can get JsObeject
	//jsob, _ := myjson.GetJsObject(bytes)
	//value = myjson.GetFromJsObject(jsob, "date.time.minutes")
	//fmt.Println(value, ":", value.GetAsIntIgnore())
	//fmt.Println(myjson.GetFromJsObject(jsob, "date.datetime").GetAsStringIgnore())

}
