/*
Misol Sharti
n soni berilgan.

Shu son tub (prime) ekanligini aniqlang. Agar tub bo'lsa true 
 bo'lmasa false konsolga chiqaring.

Tubson  1 ga va o'ziga bo'linadigan sonlar.  Masalan 7. U 1 ga va o'ziga bo'linadi u 
boshqa songa bo'linmaydi. Uning bo'linuvchilari 2ta.
*/
package main 
import ("fmt")

func main() {
	var n int 
	fmt.Scan(&n)
	sum:=0
	for i:=1; i<=n; i++ {
		if n%i==0 {
			sum++
		}
	}
	if sum==2 {
		fmt.Println("true")

	}else{
		fmt.Println("false")
	}
}