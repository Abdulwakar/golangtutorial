package main
import(
    "fmt"
    "math/cmplx"
)
 var (
        ToBe bool = false
        MaxInt uint64 = 1<<3-1
        z complex128 = cmplx.Sqrt(3 + 4i)
    )
func main(){

    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)

    // Zero variables: Variable declared without an explicit values are given their zero values
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q \n",i, f, b, s)
}