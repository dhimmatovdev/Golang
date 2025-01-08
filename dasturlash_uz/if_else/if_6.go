/*
Misol Sharti
Berilgan son hafta kunining qaysi kuniga to'g'ri keladi .
Agar kelgan son 1 ga teng bo'lsa 'Dushanba' ni konsolga chiqaring, 
Agar kegan son ikkiga teng bo'lsa  'Seshanba' ni konsolga chiqaring va hz .., 
Agar hafta kuniga to'g'ri kelmasa 'none' konsolga chiqsin.
*/

package main
import ("fmt")

func main() {
	var a int 
	fmt.Scan(&a)
	// if a==1 {
	// 	fmt.Println("Dushanba")
	// }else if a==2 {
	// 	fmt.Println("Seshanba")
	// }else if a==3 {
	// 	fmt.Println("Chorshanba")
	// }else if a==4 {
	// 	fmt.Println("Payshanba")
	// }else if a==5 {
	// 	fmt.Println("Juma")
	// }else if a==6 {
	// 	fmt.Println("Shanba")
	// }else if a==7 {
	// 	fmt.Println("Yakshanba")
	// }else{
	// 	fmt.Println("Mavjud emas son")
	// }
	switch a {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Seshanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6:
		fmt.Println("Shanba")
	case 7:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("Mavjud emas son")
	}
}	