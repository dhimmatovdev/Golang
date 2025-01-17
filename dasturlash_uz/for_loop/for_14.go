/*
Misol Sharti
 a soni berilgan.

 Shu son mukammal raqam yoki yo'qmi shuni aniqlang.

 Agar mukammal raqam bo'lsa true bo'lmasa false konsolga chiqaring.

Mukammal raqam deb bo'linuvchilar  yig'indisiga teng bo'ladigan songa (o'zidan tashqari) aytiladi.

Deylik

 a = 6, uning bo'linuvchilari 1,2,3
 6 = 1 + 2 + 3   demak 6 raqami murakkab raqam hisoblanadi.

 a = 15, uning bo'linuvchilari 1,3,5,
 1 + 3 + 5 = 8    bo'linuvchilari yig'indisi 8ga teng 15 ga emas. Shuning uchun bu mukammal son emas.

*/

package main
import ("fmt") 

func main() {
	var a int 
	fmt.Scan(&a)
	sum:=0
	for i:=1; i<a; i++ {
		if a%i==0 {
			sum=sum+i
		}
	}
	if a == sum {
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
}