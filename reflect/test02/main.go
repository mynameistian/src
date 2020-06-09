package main

import (
	"fmt"
	"reflect"
)

// Buff struct
type Buff struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

//SetBuff  SetValue
func (buff *Buff) SetBuff(name string, age string) {
	buff.Name = name
	buff.Age = age
}

//PutBuff PutBuff
func (buff *Buff) PutBuff() {
	fmt.Printf("Name is [%v] , Age is [%v] \n", buff.Name, buff.Age)
}

//ReflectFunc func
func ReflectFunc(reflectPrt interface{}) {
	rVal := reflect.ValueOf(reflectPrt)
	//rType := reflect.TypeOf(reflectPrt)

	//rKind := rVal.Kind()
	//fieldNum := rType.NumField()
	funcNum := rVal.NumMethod()

	// for i := 0; i < fieldNum; i++ {
	// 	fmt.Printf("FieldName is [%v]\n", rType.Field(i).Name)
	// }
	fmt.Printf("Method Num is [%v]\n", funcNum)
	rVal.Method(1).Call(nil)

}

func main() {
	var buff = Buff{
		Name: "Tianlijun",
		Age:  "24",
	}
	ReflectFunc(buff)

	fmt.Printf("Name is [%s] Age is [%s]", buff.Name, buff.Age)
}
