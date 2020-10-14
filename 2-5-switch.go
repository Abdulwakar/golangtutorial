package main
import(
    "fmt"
    "runtime"
   "time"
)

func main(){
    fmt.Print("Go runs on ")
    switch os:= runtime.GOOS ; os{ //here we defined runtime.GOOS as condition
        case "darwin":
            fmt.Println("Os X. ")
        case "Linux":
            fmt.Println("Linux. ")
        default:
            fmt.Printf("%s \n", os)
    }

    // another example fo switch case
    fmt.Println("When is Saturday?")
    today := time.Now().Weekday()
    switch time.Saturday { //Here we defined time.Saturday as condition
        case today+0:
            fmt.Println("Today.")
        case today+1:
            fmt.Println("Tomorrow.")
        case today+2:
            fmt.Println("Day after tomorrow")
        default:
            fmt.Println("Too far away")
    }

    //Another example without switch condition
    t := time.Now()
    switch {
    case t.Hour()<12:
        fmt.Println("Good Morning")
    case t.Hour()<17:
        fmt.Println("Good Afternoon")
    default:
        fmt.Println("Good Evening")
    }
}