/*
Misol Sharti
a, b sonlar berilgan.

 a va b son chiqarlari musbat sonlar sonini toping va 
 shuni konsolga chiqaring. b sonini ham inobatga oling.


*/
package main
import ("fmt")

func main() {
	var a,b int 
	fmt.Scan(&a, &b)
	sum:=0
	for i:=a; i<=b; i++ {
		if i>0 {
			sum++
		}
	}
	fmt.Println(sum)
}