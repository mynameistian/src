package main

import (
	"fmt"
	"reflect"
)

func float64Reflect(b interface{}) {
	rType := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)
	rKind := rVal.Kind()

	fmt.Printf("b type is [%v] Kind is [%s] \n", rType, rKind)
	fmt.Printf("b value is [%v] \n", rVal.Elem())
	rVal.Elem().SetFloat(1.3)
}

func main() {
	var v float64 = 1.2
	float64Reflect(&v)
	fmt.Printf("v is [%v]", v)
}
