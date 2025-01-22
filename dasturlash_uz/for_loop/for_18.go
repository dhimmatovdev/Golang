/*
n soni berilgan.

Shu sonni raqamlarini sonini  konsolga chiqaring.

Namuna
*/

package main 
import ("fmt")
func main() {
	var a = 1234
	temp:=0
	for i:=0; i<a; a/=10 {
		temp++
	}
	fmt.Println(temp)
}