package main
import "fmt"

const Pi = 3.14

func main() {
    const World = "Abdul"
    fmt.Println("Hello %s", World)
    fmt.Println(Pi)
    const Truth = true
    Truth := 14
    fmt.Println("Go rules?", Truth)

}