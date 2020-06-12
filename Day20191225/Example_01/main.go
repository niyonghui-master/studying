package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	//fmt.Println("type:", reflect.TypeOf(x), "kind:", reflect.TypeOf(x).Kind())
	//fmt.Println("value:", reflect.ValueOf(x).String(), "kind:", reflect.TypeOf(x).Kind())

	//v := reflect.ValueOf(x)
	//
	//fmt.Println("type:", v.Type())
	//fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	//fmt.Println("value:", v.Float())

	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}
