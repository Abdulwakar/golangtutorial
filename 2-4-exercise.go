package main
import (
    "fmt"
   // "math"
    )

func main(){
    fmt.Println(Sqrt(1000000))
}

func Sqrt(x float64) float64{
    z := float64(1)
    for i:=1; i<=10; i++ {
        z = z -(z*z-x)/(2*z)
        fmt.Println(z)
    }
    return z
}