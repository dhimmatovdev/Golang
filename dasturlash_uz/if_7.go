/*
Misol Sharti
a,b,c sonlari berilgan . Shu sonlar uchburchakning ichki burchaklari ekanligini aniqlang.
Agar berilgan  sonlar uchburchakning burchaklari bo'lsa ,  true konsolga chiqaring bo'lmasa false konsolga chiqaring.
*/

package main 
import ("fmt")

func main() {
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	// Uchburchakning ichki burchaklari a,b,c ga teng bo'lsa true aks holda false chiqaradi
	if a+b+c==180 { 
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}