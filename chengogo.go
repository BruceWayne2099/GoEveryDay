package main

import (
	"fmt"
	"strconv"
)

var name string = "james"
var number int = 23
var team string = "lakers"
var girl string = "jerry"
const str = ` 武汉
加油
全国人民支持你
一定要战胜病魔
`
const limit  = 512

const (
	x,y int = 1,3
	Cyan = 1
)

func main() {
	var pian = make([]int,5)
	pian[0] = 99
	pian[1] = 77
	pian[2] = 666

	s2 := strconv.Itoa(number)

	fmt.Println("阿里生活的第一天!")
	fmt.Printf("%s is a player of %s,his number is %s. \n",name,team,s2)
	fmt.Println(girl)
	fmt.Println(pian[0],pian[2],pian[1])
	slogen := "武汉，加油！"
	fmt.Println(slogen)
	fmt.Println("文件保存在：'C:\\go\\test\\gogogo.exe'")
	fmt.Println(str)
	chen := &name
	shuo := *chen
	fmt.Printf("%p,%s,%p",chen,shuo,&team)
	fmt.Println(Cyan,limit)


	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagYellow
	)
	fmt.Println(FlagGreen,FlagYellow,FlagRed)

	var girls [3] string
	girls[1] = "lily"
	girls[2] = "lucy"
	girls[0] = "May"

	fmt.Println(girls)

	var players = [...]string {"james","davis","rondo","kuzima","green"}	//这是一个数组
	var players2 [3]string	// 这也是一个数组
	players4 := append(players2[:],"mcgee")
	copy(players2[0:2], players[:])		//通过 [:]讲数组转换成切片,copy 拷贝  players2有一个空值

	players3 := append(players2[:],"shaq","kobe")	//这是一个切片
	fmt.Println(players2,players3,players4)
	fmt.Printf("%T and %T and %T",players,players2,players3)


}