package main
 
import (
    "fmt"
)
 
func main(){
    var n int 
    fmt.Scan(&n)
    if n == 2 || n % 2 == 1 {
        fmt.Println("NO")
    }else {
        fmt.Println("YES")
    }
}
