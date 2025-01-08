/*
Misol Sharti
Sonlar o'qida a,b sonlari berilgan. Ular orasidagi masofani toping.
*/

package main 
import ("fmt")
func main() {
	var (
		a = 2
		b = 4
	)
	if a>b {
		fmt.Println(a-b)
	}else{
		fmt.Println(b-a)
	}
}