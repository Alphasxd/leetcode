// 通过使用函数名的字符串，调用函数。采用反射的Call方法实现

package main

import (
	"fmt"
    "reflect"
)

type Animal struct{}

func (a *Animal) Eat(){
    fmt.Println("Eat ... ")
}

func main(){
    a := Animal{}
    reflect.ValueOf(&a).MethodByName("Eat").Call([]reflect.Value{})
}