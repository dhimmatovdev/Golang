/*
Misol Sharti
a, b sonlar berilgan.

a va b sonlar uchun 3 ga bo'linadigan sonlarni yig'indisi va sonini konsolga chiqaring.
*/
package main

import (
	"fmt"
)

func main() {
	var a, b int 
	fmt.Scan(&a, &b)
	sum:=0
	count:=0
	for i:=a; i<b; i++ {
		if i%3==0 {
			sum++
			count+=i
		}
	}
	fmt.Println(count, sum)
}