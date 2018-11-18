<p align="center">
    <img 
         src="logo.png"
         width="382" height="351" border="0" alt="MYJSON"/>
    <br/>
</p>





MYJSON is a Go package that provides a fast and simple way to get values from a json document.

## Getting Started

### Install

To start using MyJson,Install Go and run <code>go get</code>:

```bash
$ go get -u github.com/GuoYuefei/myjson
```

This will retrieve the library.

### Usage Example

```go
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
```



## License

Eclipse Public License 2.0

## Preface

The official json parsing library uses the reflection feature, which determines its speed is not too fast. I am still in the first grade of the master's degree, and I just want to challenge myself and write a json analysis library. It is still in the process of perfection, but the deserialization of json has been relatively friendly. Here I declare that I will improve the code, and I will maintain the json parsing library and improve the documentation whenever I have time.

