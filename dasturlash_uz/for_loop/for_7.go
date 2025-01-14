/*
Misol Sharti
a, b sonlar berilgan

a va b sonlar juft sonlarni sonini toping
 (a sonini inobatga oling) va natijani konsolga chiqaring.
*/
package main
import ("fmt")

func main() {
	var a, b int 
	fmt.Scan(&a,&b)
	sum:=0
	for i:=a; i<b; i++ {
		if i%2==0 {
			sum++
		}
	}
	fmt.Println(sum)
}