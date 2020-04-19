package main

import "fmt"

type guaiwu interface {
	attack()			//做了叫 guaiwu 的接口，接口的方法是attac()
}

type long struct {	// 设定 long 是一个结构体
}

func (dragen long) attack() {	// 传入一个参数叫dragen
	fmt.Println("龙喷火")
}

type yeshou struct {
}

func (tiger yeshou) attack() {	//   传入一个参数tiger
	fmt.Println("野兽抓人")
}

func main(){
	var onikxiya guaiwu	//创建了onikxiya 这个变量，他是属于guaiwu 这个"类"的，所以它有attack的方法

	onikxiya = new(long)	// 当它属于不同的
	onikxiya.attack()


}