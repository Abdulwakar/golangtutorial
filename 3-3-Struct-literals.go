package main
import(
    "fmt"
)

type Vertex struct{
    x, y int
}

var(
    v1 = Vertex{5,2}
    v2 = Vertex{x: 1} // This is struct literals, which replace x value from vertex
    v3 = Vertex{} // Struct literals replacing both x & y value to 0.
    p = &Vertex{10,20}
 )


func main(){
    fmt.Println(v1, p, v2, v3)
}