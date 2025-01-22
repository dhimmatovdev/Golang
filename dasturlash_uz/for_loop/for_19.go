/*
Misol Sharti
n  soni berilgan.

Shu sonni raqamlarini yig'indisini toping. 

*/
package main 
import ("fmt")

func main() {
	var n = 1234;
	sum:=0
	for n>0 {
		digits:=n%10
		sum+=digits
		n/=10
	}
	fmt.Println(sum)
}