package main

import (
	"fmt"
	"reflect"
)

func ADD(a,b string) int{
	resulta := reflect.TypeOf(a)
	resultb := reflect.TypeOf(b)

	fmt.Println(resulta == resultb)

	return 666
}

func main()  {
	ADD("james","shaq")

}