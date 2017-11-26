package main

import (
	"fmt"
	"reflect"
)

func main() {
	//反射三法则  https://blog.go-zh.org/laws-of-reflection
	a := 4
	t := reflect.TypeOf(a)
	fmt.Printf("(1)接口值到反射对象类型: %T , %v \n", t, t)
	fmt.Println(t.Kind().String(), t.Name(), t.String())
	v := reflect.ValueOf(a)
	fmt.Printf("(1)接口值到反射对象值:%T , %v \n", v, v)
	a1 := v.Interface().(int)
	fmt.Println("(2)反射对象到接口值:", a1)

	var x = 3.4
	vf := reflect.ValueOf(x)
	//vf.SetFloat(7.1) // Error: will panic.
	fmt.Println("(3)基本类型等值是不可修改的！:", vf)
	//pvf := reflect.ValueOf(&x)
	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	if p.Kind() == reflect.Ptr {
		v := p.Elem()
		v.SetFloat(7.1)
		fmt.Println("Value set:", v.Interface(), x)
	}
}
