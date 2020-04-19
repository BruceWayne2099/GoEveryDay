package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("bulls")
	l.PushBack("lakers")
	l.PushFront(666)

	element := l.PushBack("magic")
	l.InsertBefore("pistons",element)
	l.InsertAfter("bucks",element)

	fmt.Println(l)

	//对for进行遍历，其中i=l.Front()表示初始负值，只会在开始的时候执行一次。每次循环都会进行判断，如果i !=nil 出现了false就会退出循环，反之就会i=l.Next
	for i:= l.Front();i !=nil ;i = i.Next() {
		fmt.Println(i.Value)
	}
}