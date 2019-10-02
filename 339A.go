package main
 
import(
    "fmt"
    "sort"
    "strings"
)
 
func main(){
    var s string
	fmt.Scan(&s)
	ans := strings.Split(s, "+")
	sort.Strings(ans)
	fmt.Println(strings.Join(ans, "+"))
}