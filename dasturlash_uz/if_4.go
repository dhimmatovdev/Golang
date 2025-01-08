/*Misol Sharti
c char berilgan. Berilgan  char alfavit  yoki raqam ekanligini aniqlang. Agar alfavit bo'lsa 'alpha' ni  konsolga chiqaring,  agar son bo'lsa 'son' ni  konsolga chiqaring.*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var char rune
	fmt.Print("Belgini kiriting: ")
	fmt.Scanf("%c", &char)

	if unicode.IsLetter(char) {
		fmt.Println("alpha")
	} else if unicode.IsDigit(char) {
		fmt.Println("son")
	} else {
		fmt.Println("Notanish belgi")
	}
}
