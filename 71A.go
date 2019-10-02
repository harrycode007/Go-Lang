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
    var t, i int
    fmt.Scan(&t)
    for i = 0 ; i < t ; i++ {
        var s string
        fmt.Scan(&s)
        if len(s)>10 {
			fmt.Printf("%c%d%c\n", s[0] , len(s) - 2 , s[len(s) - 1])
		} else {
			fmt.Println(s)
		}
    }
}