package main
import(
    "fmt"
)

func main(){
    var a [2] string //declare array as variable 'a'
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2,3,4,5,6,2} //declare 'prime' as array
    fmt.Println(primes)
    fmt.Println(primes[2])

    //Slices
    var s[]int = primes[1:3]
    fmt.Println(s)
}