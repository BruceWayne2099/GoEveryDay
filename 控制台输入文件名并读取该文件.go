package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	var FileName string
	fmt.Println("plz 输入一个文件名:")
	fmt.Scan(&FileName)
	fmt.Println("从控制台交互得到的文件名是：", FileName)
	b, err := ioutil.ReadFile("/Users/chris/Documents/" + FileName)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(b)
	str := string(b)
	fmt.Println(str)
}