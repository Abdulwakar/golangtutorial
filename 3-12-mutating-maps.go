package main
import "fmt"

func main(){
    m := make(map[string]int)

    m["Answer"] = 32
    fmt.Println("The value: ", m["Answer"])

    m["Answer"] = 49
    fmt.Println("The value: ", m["Answer"])

    delete(m,"Answer")
    fmt.Println("The value: ", m["Answer"])

    v, ok := m["Answer"]
    fmt.Println("The value: ", v ,"Present?", ok)
}