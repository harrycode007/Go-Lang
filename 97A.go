package main
 
import (
	"fmt"
)
 
func main() {
	var n int
	fmt.Scan(&n)
	var a , b[10000] int
	var i int
	for i = 1 ; i <= n ; i++ {
	    fmt.Scan(&a[i])
	}
	for i = 1 ; i <= n ; i++ {
	    b[a[i] - 1] = i
	}
	for i = 0 ; i < n ; i++ {
	    fmt.Print(b[i])
	    fmt.Print(" ")
	}
}
