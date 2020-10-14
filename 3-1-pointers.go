package main
import(
    "fmt"
)

func main(){
    i, j := 10, 21

    p := &i // point to i
    fmt.Println(*p) // read i through pointer
    *p = 11 // set i through pointer
    fmt.Println(i) // see value of i

    q := &j
    *q = *q / 3
    fmt.Println(j)
}