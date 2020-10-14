package main
import "fmt"

type Vertex struct{
    x int
    Y int
    z float32
}

func main(){

    fmt.Println(Vertex{1,2,31.2})
    v := Vertex{13,14,142.21}
    x := v.Y // Here we are accessing Y value from Struct Fields.
    fmt.Println(x)

    // Pointer to Structs

    p := &v
    p.x = 1e3
    fmt.Println(v)

}