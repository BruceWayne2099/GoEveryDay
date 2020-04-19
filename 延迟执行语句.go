package main

import "fmt"

func name(a,b string)int{
	fmt.Println(a)
	fmt.Println(b)
	return 333
}

func number(x,y int)bool {
	defer name("gaogao","yigyig")
	fmt.Println(x)
	fmt.Println(y)
	return false
}

func main() {
	defer fmt.Println("down机之后的事情1")
	defer fmt.Println("down机之后的事情2")
	name("oleilei","olala")
	defer number(999,007)
	defer fmt.Println("666")
	panic("挂了")
}
