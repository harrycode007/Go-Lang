package main
 
import (
    "fmt"
)

func sum(a []int) int {
	var i, sum int
	for i = 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}
 
func main(){
    var q , i int
    for i = 0 ; i < q ; i++ {
        var n int 
        fmt.Scan(&n)
        a := make([]int, n)
        var ans int
        for i = 0 ; i < n ; i++ {
            fmt.Scan(&a[i])
        }   
        ans = sum(a)
        if ans % n != 0 {
            fmt.Println(ans/n + 1)
        }else{
            fmt.Println(ans/n)
        }
     }
}
