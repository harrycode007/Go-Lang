package main
 
import (
	"fmt"
	"sort"
)
 
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a != 1 && b != 1 && c != 1 {
		fmt.Print(a * b * c)
		return
	}
	var ans []int
    ans = append(ans, (a + b) * c)
	ans = append(ans, a * (b + c))
    ans = append(ans, a + b + c)
	sort.Ints(ans)
	fmt.Print(ans[2])
}
