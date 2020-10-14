package main
import "fmt"

var pow = []int{1,2,4,8,16,64,128}

func main(){
    for i, v:= range pow{ //ranging over slice, two values are returned for each iteration
        fmt.Printf("2**%d = %d\n",i, v) // First is index and second is copy of element which is index.
    }

    fmt.Println("Skipping index or values")
    pow := make([]int, 10)
    for i := range pow{ // this loop will assign values in array
        pow[i] = 2<<uint(i)
    }
    fmt.Println("***********")
    for _, value := range pow{ // skip index value from range
        fmt.Printf("%d\n", value)
    }
}