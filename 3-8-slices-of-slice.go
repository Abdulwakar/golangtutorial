package main
import (
    "fmt"
   "strings"
)

func main() {
    //Create a board
    board := [][]string{
        []string{"_","_","_"},
        []string{"_","_","_"},
        []string{"_","_","_"},

    }
    //The player take turns.

    board[0][0] = "X"
    board[2][2] = "0"
    board[1][2] = "a"
    board[1][0] = "0"
    board[0][2] = "0"

    fmt.Println(len(board))
    fmt.Println(board[1])
    for i:=0; i< len(board); i++{
        fmt.Printf("%s\n", strings.Join(board[i],""))
    }
}