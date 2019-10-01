package main
 
import (
	"fmt"
)
 
func main() {
	var t , i int
	fmt.Scan(&t)
	var cnt int
	    cnt = 0
	for i = 0 ; i < t ; i++ {
	    var p , q int
	    fmt.Scan(&p , &q)
	    if q - p >= 2 {
	        cnt++;
	    }
	}
	fmt.Println(cnt)
}