/* Misol Sharti
 a soni berilgan. Berilgan son 5 ga bo'linadimi? Agar bo'linsa  konsolga true bo'lmasa false chiqaring.*/

 package main
 import ("fmt")

 func main() {
	var a int
	fmt.Scan(&a)
	// 5 ga qoldiqsiz bo'linsa true aks holda false chiqaradi va biz % belgidan foydlanamiz
	if (a%5==0) { 
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

 }