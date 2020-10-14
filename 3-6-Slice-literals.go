package main
import (
    "fmt"
)

func main(){
    q := []int{2,3,4,5,6,8}
    fmt.Println(q)

    r := []bool{true, false, true, false, true, false}
    fmt.Println(r)

    s := []struct{
        i int
        b bool
    } {
        {2, true},
        {3, false},
        {4, true},
        {5, false},
        {6, true},

    }
    fmt.Println(s[3])
}