package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var num int = 10

	if num > 10 {
		fmt.Println("bigger than 10!")
	} else {
		fmt.Println("not bigger than 10!")
	}

	var step int
	for ; step < num; step ++{
		string := strconv.Itoa(step)
		fmt.Println("骑士攻击"+ string + "次")
		time.Sleep(1000000000)		//等待1秒
	}

	playlist := map[string]int{
		"黄颖": 83,
		"高莹": 90,
		"杜楠": 85,
		"陈默": 88,
		"王雨": 88,
	}

	for name,score := range playlist{
		fmt.Println(name,score)
	}

	a := 10
	switch a {
	case 11:
		fmt.Println("dayu 11!")
	case 9:
		fmt.Println("xiaoyu 10")
	default:
		fmt.Println("dengyu 101")
	}

	fmt.Println(strings.Index("woshinidaye","a"))

	i := 0
	HERE:
		fmt.Println(i)
	i++
	if i < 10{
		goto HERE
	}
}


