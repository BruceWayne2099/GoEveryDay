package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)


func swap(a,b *int){
	t := *a
	*a = *b
	*b = t
}

func swap222(a,b string){
	m := a + "aaaaaa"
	n := b + "bbbbbb"
	fmt.Println(len(m),utf8.RuneCountInString(n))

	for _,s :=range n{
		fmt.Printf("%c \n",s)
	}
}

func calc(a,b string) string{
	m := a + "ccccc"
	n := b + "ddddd"

	x := m+n
	return x
}

func main(){
	x,y :=   666,999
	swap(&x, &y)
	fmt.Println(x,y)

	girl1,girl2 := "黄颖","高莹"
	swap222(girl1,girl2)

	fmt.Println(calc(girl2,girl1))

	str1 := "superman,batman,wonderwoman"
	str2 := "nba mvp"
	str3 := "第二目标，就是2022年争取升至P8"

	comma := strings.Index(str1,",")
	pos := strings.Index(str1[comma:], "man")

	fmt.Println(comma,pos,str1[comma+pos:])

	fmt.Println(str2[2:6],str3)

	anglebytes := []byte(str2)
	fmt.Println(string(anglebytes))

	hammer := "吃我一锤！"
	sickle := "勇士阵亡了"

	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)
	stringBuilder.WriteString(str3)

	fmt.Println(stringBuilder.String())

	var number = 30
	var target = 50
	title := fmt.Sprintf("已经采集  %v 个草药，要需要 %v 才能完成任务", number,target)
	fmt.Println(title)
}
