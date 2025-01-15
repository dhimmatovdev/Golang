/*
Misol Sharti
a , n sonlari berilgan (a soni 1-9 o'rtasida ).

Ketma ketlikni hisoblang. a + aa + aaa + ........  nta a bo'ladi.

Yig'indini konsolga chiqarish kerak. 


*/
package main
import ("fmt") 
func main() {
	var a, n int 
	fmt.Scan(&a, &n)
	sum:=0
	temp:=0

	for i:=0; i<n; i++ {
		temp=temp*10+a
		sum=sum+temp
		
	}
	fmt.Println(sum)
}