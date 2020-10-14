package main

import "fmt"

var i, j int = 1, 2

var k = 50 // we are declaring variable outside of function.
// we can't use "k := 50 here. "

func main(){
    var c, python, java = true, false, "no!"
    l := 100 // declaring as short variable
    fmt.Println(i, j, c, python, java, k, l )
}