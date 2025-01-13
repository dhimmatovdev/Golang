/*
Misol Sharti
a soni berilgan.  0 dan a gacha bo'lgan sonlaryig'indisini 
toping  va yig'indini konsolga chiqaring. a sonini ham hisobga oling.
*/

package main
import ("fmt")

func main() {
	var a int
	fmt.Scan(&a)
	var sum int
	for i := 0; i <= a; i += 2 {
		sum += i
	}
	fmt.Print(sum)
}
