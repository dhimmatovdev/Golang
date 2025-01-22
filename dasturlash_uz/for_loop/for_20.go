/*
Misol Sharti
n soni berilgan. Shu sonni teskarisini toping va konsolga chiqaring.

Masalan: n = 12345

uning teskarisi 54321; 
*/
package main 
import ("fmt")

func main() {
	var n = 12345
	sum:=0
	for n>0 {
		digit:=n%10
		sum=sum*10+digit
		n/=10
	}
	fmt.Println(sum)
}