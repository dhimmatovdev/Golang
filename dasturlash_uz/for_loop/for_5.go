/*
Misol Shartif
a soni berilgan. 0 dan a gacha bo'lgan juft sonlar soni toping va konsolga chiqaring. a sonini inobatga olmaymiz.


*/
package main 
import ("fmt") 

func main() {
	var a int 
	sum:=0
	fmt.Scan(&a)
	for i:=0; i<a; i++ {
		if i%2==0 {
			sum++
		}
	}
	fmt.Print(sum)
	
}