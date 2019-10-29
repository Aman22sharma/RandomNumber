package main
        
import (
    "fmt"
    "math/rand"
)
        
func main() {
    min := 10
    max := 30
    fmt.Println(rand.Intn(max - min) + min)
}