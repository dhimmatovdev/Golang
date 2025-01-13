/*
Misol Sharti
a soni berilgan. 0 dan  a gacha bo'lgan toq sonlar yig'indisini toping  va yig'indini konsolga chiqaring. a sonini ham inobatga oling.


*/
package main 
import ("fmt")

func main() {
	var a int 
	fmt.Scan(&a)
	sum:=0
	// for i:=1; i<=a; i+=2 {
	// 	sum=i+sum
	// }
	for i:=0; i<=a; i++ {
		
		if i%2!=0 {
			sum+=i;
		}
	}
	fmt.Println(sum)
}