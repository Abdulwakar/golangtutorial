package main
import "fmt"

func main(){

    s := []int{2,3,4,5,6,7}
    fmt.Println(s)

    s = s[1:4]
    fmt.Println(s)
    printSlice(s) // calculating length & capacity of slice

    s = s[:3]
    fmt.Println(s)

    s = s[1:]
    fmt.Println(s)

    a := make([]int, 5) // this will create empty array of 5
    SliceMake("a", a)

    b := make([]int, 0,5) // specify ending in "make" create 5 capacity array
    SliceMake("b", b)

    c := b[:2]
    SliceMake("c", c)

    d := c[2:5]
    SliceMake("d", d)

    // append slice
    e := append(d, 1)
    printSlice(e)


}

func printSlice(s []int){
    fmt.Printf("len=%d cap=%d %v\n",len(s), cap(s), s)
}

func SliceMake(s string, x []int){
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)

}