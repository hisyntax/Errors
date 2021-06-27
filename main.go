package main 

import (
    "errors"
    "fmt"
)

func main() {
    err := errors.New("New error message")

    fmt.Println(err)
}