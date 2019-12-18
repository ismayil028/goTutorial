package main

import "fmt"

func main() {
	//dynamicTypeDeclaration()
	//mixedVariableDeclaration()
	//escapeSequences()
	constKeyWord()

}

func constKeyWord() {
	const LENGTH int = 10
	const WITDH int = 5
	var area int
	area = LENGTH * WITDH
	fmt.Printf("value of area : %d", area)
}

func escapeSequences() {
	fmt.Println("\\") // output \
	fmt.Println("\"") // output "
	fmt.Println("\a") // output alert or bell
	fmt.Println("\b") // output alert or bell
	fmt.Println("\f") // form feed
	fmt.Println("\n") // new line
	fmt.Println("\r") // carriage return
	fmt.Println("\t") // horizontal tab
	fmt.Println("\v") // vertical tab
	fmt.Println("\112")
	fmt.Println("\x22U")

}
func dynamicTypeDeclaration() {
	var x float64 = 20.0

	y := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)

}
func mixedVariableDeclaration() {
	var a, b, c = 3, 4, "foo"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)

}
