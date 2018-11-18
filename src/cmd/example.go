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
	//or you can get JsObeject that defined in newtypes.go
	jsob, _ := myjson.GetJsObject(bytes)
	//you can use function GetFromJsObject to get information that you want to get
	value = myjson.GetFromJsObject(jsob, "date.time.minutes")
	fmt.Println(value, ":", value.GetAsIntIgnore())
	fmt.Println(myjson.GetFromJsObject(jsob, "date.datetime").GetAsStringIgnore())
	//You can find methods in type Value, and you can quickly use them in name.
}
