/*
Misol Sharti
a, b, c sonlar berilgan.

Agar shu solarning ikkitasi o'zaro teng bo'lsa  teng bo'lmagan sonni konsolga chiqaring, aks holda 0 ni chiqaring.


*/

package main

import (
	"fmt"
)

func main() {
	var (
		a = 3
		b = 10
		c = 8
	)
	// if a != b && a != c {
	// 	fmt.Println(a)
	// } else if b != a && b != c {
	// 	fmt.Println(b)
	// } else {
	// 	fmt.Println("0")
	// }
	if a == b {
		fmt.Println(c)
	} else if a == c {
		fmt.Println(b)
	} else if b==c {
		fmt.Println(a)
	}else{
		fmt.Println("0")
	}
}
