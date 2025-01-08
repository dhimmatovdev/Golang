/*
Misol Sharti
a, b, c sonlar berilgan.

Agar sonlar o'sish tartibida  joylashgan bo'lsa  1 ni konsolga chiqaring.
Agar kamayish tartibda bo'lsa 2 konsolga chiqaring.
Agar eng kattasi b bo'sa  b ni konsolga chiqaring.
Agar ular teng bo'lsa 5 ni konsolga chiqaring.
Bo'lmasa 0 ni konsolga chiqaring. 


*/

package main 
import ("fmt")

func main() {
	var (
		a = 8
		b = 5
		c = 1
	)
	if a < b && b < c {
		fmt.Println("1")
	}else if a>b && b>c {
		fmt.Println("2")
	}else if a<b && b>c {
		fmt.Println(b)
	}else if a==b && b==c {
		fmt.Println("5")
	}else{
		fmt.Println("0")
	}
}