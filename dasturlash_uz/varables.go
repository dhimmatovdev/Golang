/*
:= faqat yangi o'zgaruvchilarni aniqlash uchun ishlatiladi. Agar o'zgaruvchi ilgari e'lon qilingan bo'lsa, 
qayta := ishlatilsa, xato yuzaga keladi.

Go statik tipli til bo'lgani uchun, qiymatlar tipini o'zi aniqlaydi va har bir o'zgaruvchiga mos ravishda tip tayinlaydi.
*/
package main
import ("fmt")

func main() {
  var a, b = 6, "Hello"
  c, d := 7, "World!"

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}