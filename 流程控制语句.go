package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	Name string
	Number int
	input string
	err	error
)

func main() {
	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("请输入一个名字：")

		input, err = f.ReadString('\n')
		//读取器对象提供一个方法 ReadString(delim byte) ，该方法从输入中读取内容，直到碰到 delim 指定的字符，然后将读取到的内容连同 delim 字符一起放到缓冲区。
		if len(input) == 1 {
			continue
		}
		fmt.Printf("你的名字是%s", input)

		fmt.Sscan(input,&Name,&Number)
		switch input {
		case "james\n":
			fmt.Printf("you are a loser!")
		case "curry\n":
			fallthrough //使用下一个条件句
		case "harden\n":
			fmt.Println("your 3pt is very COOL!")
		case "gorden\n", "zimuge\n":
			fmt.Println("your dunk is very COOL!")
		case "kobe\n":
			fmt.Println("i am sorry for you dead!")
		default:
			fmt.Println("who the hell are you,get out of there!")
		}
	}
}
