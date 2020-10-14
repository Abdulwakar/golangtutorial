package main
import "fmt"
func main(){
    sum := 0
    // for loop has, initials, conditions & post statements.
    for i := 0; i < 10; i++{
        fmt.Printf("Current value of I is %d", i)
        sum += i
        fmt.Println("current value is: ", sum)
    }
    fmt.Println(sum)

    // The init and post statements are optional in For loop
    total := 5
    for ; total<15;{
        fmt.Println(total)
        total += total
    }
    fmt.Println(total)
}