package main

import(
  "fmt"
  //"github.com/user/stringutil"
)
func main(){
  var a [3] int
  a[0]=1
  a[1]=2
  a[2]=3
  // Returns a slice starting from first param ends excluding second param
  fmt.Println(a[1:2])

  for i, v := range a {
    fmt.Printf("%d %d\n", i, v)
}
}
