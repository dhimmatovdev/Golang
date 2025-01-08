/* Misol Sharti
 Berilgan son 3 va 4 ga bo'linadimi? Agar bo'linsa konsolga 'true'  bo'lmasa falseni chiqaring.
*/

package main 
import ("fmt")

func main() {
	var a int
	fmt.Scan(&a)
	// 3 va 4 ga qoldiqsiz bo'linsa true aks holda false chiqaradi va biz % belgidan foydlanamiz
	if (a%3==0) && (a%4==0) { 
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}