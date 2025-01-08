/*
Misol Sharti
c char berilgan. Berilgan  char katta alfavit yoki kichik alfavit ekanlini aniqlang .

Agar alfavit kichik bo'lsa 'lowerCase' so'zini konsolga chiqaring,  Agar alfavit katta bo'lsa 'upperCase' ni chiqaring,  agar alfavit  bo'lmasa  'none' ni  chiqaring.

*/

package main 
import ("fmt")

func main() {
	var c = 'A'

	if c >= 'a' && c <= 'z' {
		fmt.Println("lowerCase")
	}else{
		fmt.Println("upperCase")
	}

}