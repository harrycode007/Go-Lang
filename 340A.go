package main
 
import (
    "fmt"
)
 
func main(){
     var n , ans int
     fmt.Scan(&n)
     if n % 5 == 0 {
         ans = n/5
     }else{
         ans = n/5 + 1
     }
     fmt.Println(ans)
}