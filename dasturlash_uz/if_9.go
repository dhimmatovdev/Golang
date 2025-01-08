/*Misol Sharti
 a,b,c sonlari berilgan. Bu sonlar uchburchakning 3ta tomonidir.  Uchburchak ning shakli qanday ekanligini aniqlang.

Agar teng tomonli bo'lsa  'equilateral' so'zini konsolga chiqaring.
Agar teng yonli bo'lsa  'isosceles' so'zini konsolga chiqaring.
Agar ixtiyoriy  bo'lsa  'scalene'  so'zini konsolga chiqaring.

 */

 package main
 import ("fmt")

 func main() {
	var a,b,c int 
	fmt.Scan(&a,&b,&c)

	if a==b && b==c { //teng tomonli
		fmt.Println("teng tomonli")
	}else if a==b || b==c || c==a { //teng yonli
		fmt.Println("teng yonli")
	}else{
		fmt.Println("scalene")
	}
 }