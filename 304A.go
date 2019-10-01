package main
 
import (
    "fmt"
)
 
func main(){
    var k , n , w int
    fmt.Scan(&k , &n, &w)
    var ans int 
    ans = (w * (w + 1)/2 * k)
    if ans > n {
        fmt.Print(ans - n)
    }else{
        fmt.Print("0")
    }
}