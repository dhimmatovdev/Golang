/*
Misol Sharti
soni berilgan.

a sonini 5 chi hisoblab konsolga chiqaring
*/
package main
import	("fmt")

func main() {
	var a int 
	fmt.Scan(&a)
	sum:=1
	
	for i:=0; i<5; i++ {
		sum=sum*a
	}
	fmt.Println(sum)
}