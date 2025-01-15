/*
Misol Sharti
a soni berilgan.

a sonini n chi darajasini hisoblaydigan
dastur yozing  va hosil bo'lgan sonni konsolga chiqaring.
*/
package main 
import ("fmt")

func main() {
	var a, n int 
	fmt.Scan(&a, &n)
	sum:=1
	for i:=0; i<n; i++ {
		sum=sum*a
	}
	fmt.Println(sum)
}