package main
import "fmt"

func split(sum int) (x, y int){
    x = sum * 4/9
    y = sum - x
    return // This is known as naked return, as it doesn't return anything.
}

func main(){
    fmt.Println(split(120))
}