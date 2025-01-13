//Multiple constants can be grouped together into a block for readability:

package main
import ("fmt")
const (
	A int = 1
	B = 3.14
	C = "Hello"
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

}