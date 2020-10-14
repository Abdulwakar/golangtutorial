package main
import (
    "fmt"
)

func main (){
    defer fmt.Println("First Sentence")
    // defer execute command at very last.
    fmt.Println("Second Sentence")


}