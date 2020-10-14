package main
import (
    "fmt"
)
func add(x int, y int) int{
    return x + y
}

// another example for declaring functions.
func add1(x, y int) int{
    return x + y
}

func main(){
    fmt.Println(add(12, 13))
    fmt.Println(add1(13, 14))
  }