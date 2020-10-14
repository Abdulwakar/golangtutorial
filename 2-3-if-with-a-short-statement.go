package main
import (
    "fmt"
    "math"
)

func main(){
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 4, 20),
    )

}

func pow(x, y, lim float64) float64{
    if v:= math.Pow(x, y); v< lim{
        return v
    } else{
        fmt.Printf("%g is bigger than %g \n", v, lim)
    }
    return lim

}