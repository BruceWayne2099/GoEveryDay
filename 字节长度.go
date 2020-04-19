package main

import (
	"fmt"
	"sort"
)

func test1() {
	nums := []int{1, 2, 3}

	_ = append(nums[:3:3], 4)
	fmt.Println("test1:", nums)

}


func main() {
	aaa := make([]string,3,5)
	_ = append(aaa, "james","bosh","wade")
	fmt.Println(aaa)

	//test1()

	seq := [] string{"a","b","c","d","e"}
	index := 2
	seq = append(seq[:index],seq[index+1:]...)
	fmt.Println(seq)

	sence := make(map[string]int)

	sence["route"]=66
	sence["james"]=23
	sence["wade"]=3
	sence["bosh"]=1
	fmt.Println(sence)

	sence["wade"]=9
	delete(sence,"route")		//删除
	fmt.Println(sence)

	sence = make(map[string]int)	//清空，这里是 =
	fmt.Println(sence)

	var playerlist []string
	var numlist []int
	for k,v := range sence {
		playerlist = append(playerlist, k)
		numlist = append(numlist, v)
	}

	sort.Strings(playerlist)
	sort.Ints(numlist)
	//sort.Sort(sort.Reverse(sort.IntSlice{numlist}))

	fmt.Println(playerlist,numlist)
}