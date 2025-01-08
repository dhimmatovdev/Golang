/*
Misol Sharti
 a, b, c sonlari berilgan.

Berilgan 3 ta sondan nechtasi musbat ekanligini aniqlovchi dastur yozing va musbat sonlar sonini konsolga chiqaring.

*/

package main
import ("fmt")
func main() {
	var a,b,c int
	var sum int
	fmt.Scan(&a,&b,&c)
	if a>0 {
		sum++
	}
	if b > 0 {
		sum++
	}
	if c> 0{
		sum++
	}

	fmt.Println(sum)
}