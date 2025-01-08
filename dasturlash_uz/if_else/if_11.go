/*
Misol Sharti
a, b sonlari berilgan.
Berilgan ikkita sondan  kichigini  toping  va shu sonni konsolga chiqaring. Agar ular teng bo'lsa 'teng' so'zini konsolga chiqaring.

*/

package main 
import ("fmt")

func main() {
	var (
		a = 4
		b = 1
	)
	//fmt.Scan(&a,&b)
	if a<b  {
		fmt.Println(a)
	}else if b<a {
		fmt.Println(b)
	}else{
		fmt.Println("Ular teng")
	}
}