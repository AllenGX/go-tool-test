package main

import (
	"fmt"
	"reflect"
	"test/reflect/file"
)

func test(j interface{}) {
	vIn := reflect.ValueOf(j)
	new := reflect.New(reflect.TypeOf(j))
	vOut := reflect.Indirect(new)
	for i := 0; i < vIn.NumField()-1; i++ {

		fmt.Println(vIn.Type().Field(i).Name + " is ok")
		// logs.Info(vOut.Field(i).String())
		vOut.Field(i).SetInt(int64(1))

		fmt.Println(vOut)
		// vOut.Field(i).Set(reflect.ValueOf(endpoint))
	}
}

func main() {
	i := file.Endpoints{}
	test(i)
}
