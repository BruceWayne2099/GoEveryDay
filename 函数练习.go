package main

import (
	"bytes"
	"fmt"
)

func add(a,b int) (int,bool){
	c:= a+b
	fmt.Println(c)
	return  c, false
}

func namedredtvalues() (a,b int){
	a=1
	b=2
	return
}

func joinStrings(slist ...string) string{
	var b bytes.Buffer
	for _,s := range slist{
		b.WriteString(s)
	}
	return b.String()
}

func main(){
	conn, err := add(4,9)
	fmt.Println(conn,err)

	x,y := namedredtvalues()
	fmt.Println(x,y)

	fmt.Println(joinStrings("Wade ","Bosh ","James"))

	players := map[int]string {1:"t-mac"}
	delete(players, 1)
	fmt.Printf("【%v】\n",players == nil)
	fmt.Println(players)

	list := make(map[string]map[int]bool)	// map可以嵌套
	list["curry"] = make(map[int]bool)
	list["curry"][0] = false
	list["Thompson"] = make(map[int]bool)
	list["Thompson"][1] = true
	fmt.Println(list)

}