/*
Misol Sharti
a,b,c sonlari berilgan . Shu sonlar uchburchakning tomonlari ekanligini aniqlang.
Agar kelgan sonlar uchburchakning tomonlari bo'lsa ,true  konsolga  chiqaring,  bo'lmasa false konsolga chiqaring.

Demak  a,b,c sonlari uchburchakning tomoni bo'lishi uchun ular dan  xohlagan 2tasini yig'indisi 3chisidan katta bo'lishi kerak.
*/
package main
import ("fmt")

func main() {
	var a,b,c int
	fmt.Scan(&a,&b,&c)

	if a+b>c && a+c>b && c+b>a  {
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
}