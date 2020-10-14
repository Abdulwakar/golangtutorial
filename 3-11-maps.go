package main
import(
    "fmt"
)

type Vertex struct{
    lat, long float64
}

// map map key to values.
var m  map[string]Vertex

// Below is an example for map literals.
var n = map[string]Vertex{
    "Abdul Labs": Vertex{
        12.12921 , 192.4123,
     },
     "Google Labs": Vertex{
        129.2910, -28.2931,
     },
}

func main(){
    // make function return map of given type
    m = make(map[string]Vertex)
    m["Bell Labs x"] = Vertex{
        40.2932, -82.19219,
    }
    fmt.Println(m["Bell Labs x"])
    fmt.Println(n)
}