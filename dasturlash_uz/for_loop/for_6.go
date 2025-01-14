/*
Misol Sharti
a, b sonlar berilgan.

a va b sonlar orasidagi (butun) sonlar   yig'indisini toping (a sonini inobatga oling) 
va yig'indini konsolga chiqaring.
*/

package main
import ("fmt")

func main() {
	var a, b int 
	fmt.Scan(&a,&b)
	sum:=0
	for i:=a; i<b; i++ {
		sum+=i
	}
	fmt.Println(sum)
}