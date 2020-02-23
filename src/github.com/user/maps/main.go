package main

import(
	"fmt"
)

func main(){
	mymap:= make(map[string]int)

	mymap["sam"] = 1

	fmt.Println(mymap["sam"])
}