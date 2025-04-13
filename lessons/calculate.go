package lessons

import(
	"fmt"
)

func Calc() int {
	var x int
	fmt.Println("input x")
	fmt.Scanf("%d", &x)

	var y int
	fmt.Println("input y")
	fmt.Scanf("%d", &y)
 
	var op string
	fmt.Println("input op")
	fmt.Scanf("%s", &op)

	switch op {
	case "*":
		return x * y
	case "/":
		return x / y
	case "+":
		return x + y
	case "-":
		return x - y
	case "%":
		return x % y
	default:
		panic("unrecognized escape character")
	}
}