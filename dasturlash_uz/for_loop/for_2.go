/*
Misol Sharti
a soni berilgan. 0 dan a gacha bo'lgan juft  sonlar yig'indisini toping  va yig'indini konsolga chiqaring
*/
package main
import ("fmt")
func main() {
	var a int
	fmt.Scan(&a)
	var sum int
	for i:=0; i<=a; i+=2 {
		sum+=i
	}
	fmt.Print(sum)
}