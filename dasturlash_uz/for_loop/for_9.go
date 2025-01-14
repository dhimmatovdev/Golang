/*
Misol Sharti
a, b sonlar berilgan.

a va b sonlar uchun 2 va 3 ga bo'linadigan sonlarning sonini konsolga chiqaring. 
*/

package main
import ("fmt") 

func main() {
	var a, b int 
	fmt.Scan(&a, &b)
	sum:=0
	for i:=a; i<b; i++ {
		if i%2==0 && i%3==0 {
			sum++
		}
	}
	fmt.Println(sum)
}