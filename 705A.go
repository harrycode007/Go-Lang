package main
 
import "fmt"
 
func main(){
    var n int
	fmt.Scan(&n)
	for i := 1 ; i <= n ; i++ {
		if i == 1 {
			fmt.Printf("I hate")
		} else if i % 2 == 0 {
			fmt.Printf(" that I love")
		} else {
			fmt.Printf(" that I hate")
		}
	}
	fmt.Printf(" it")
}