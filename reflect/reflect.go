package main

import (
	"fmt"
	"reflect"
)

type interStruct struct {
	C string
}

type A struct {
	a  string
	BB string
	C  interStruct
}

func main() {
	a := A{a: "a", BB: "b", C: interStruct{C: "cc"}}
	b := A{}
	copyStruct(&a, &b)
	fmt.Println(b)
}

func copyStruct(src, dst interface{}) {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()
	for i := 0; i < srcVal.NumField(); i++ {
		field, value := srcVal.Type().Field(i), srcVal.Field(i)
		if !field.IsExported() { // 判断是否大写开头的字段
			continue
		}
		dstField := dstVal.FieldByName(field.Name)
		if dstField.IsValid() {
			dstField.Set(value)
		}
	}
}
