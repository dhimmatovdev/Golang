package main

import (
	"fmt"

)

func main() {
	n := 1634
	temp := n
	sum := 0
	numberLength := 0

	// Berilgan sonning uzunligini aniqlaymiz
	for temp != 0 {
		numberLength++
		temp /= 10
	}

	// numberLength - bu n raqamining xonalari soni
	temp = n

	// Ketma-ketlikni hisoblaymiz
	for temp != 0 {
		m := temp % 10 // Oxirgi raqamni olamiz
		pow := 1

		// m raqamining numberLength-darajasini hisoblash
		for j := 0; j < numberLength; j++ {
			pow *= m
		}

		// Yig'indini oshiramiz
		sum += pow
		temp /= 10 // Raqamni 10 ga bo'lib, oxirgi xonani olib tashlaymiz
	}

	// Agar yig'indi n ga teng bo'lsa, bu Armstrong raqam
	if sum == n {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
