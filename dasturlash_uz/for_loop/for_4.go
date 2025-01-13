/*
Misol Sharti
a soni berilgan. 0 dan a gacha bo'lgan 5 ga bo'linadigan sonlar yig'indisini toping va yig'indini konsolga chiqaring. a sonini ham inobatga oling.

*/

package main
import ("fmt")

func main() {
	var a int 
	sum:=0
	fmt.Scan(&a)
	for i:=0; i<=a; i++ {
		if i%5==0 {
			sum+=i
		}
		
	}
	fmt.Println(sum)
}