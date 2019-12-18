package main

import (
	"fmt"
)

func main() {
	//dynamicTypeDeclaration()
	//mixedVariableDeclaration()
	//escapeSequences()
	//constKeyWord()
	//simpleIf()
	//IfElseExample()
	//switchCase(90)
}

func selectStatement() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("recieved ", i1, "from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("recieved ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed")
		}
	default:
		fmt.Printf("no communication")

	}
}

func switchCase(mark int) {
	var grade string = ""
	switch mark {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"

	}
	switch {
	case grade == "A":
		fmt.Printf("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well Done\n")
	case grade == "D":
		fmt.Printf("You Passed\n")
	case grade == "F":
		fmt.Printf("Better Try Again\n")
	default:
		fmt.Printf("Invalid grade0")
	}
	fmt.Printf("Your grade is %s\n", grade)
}
func simpleIf() {
	var a int = 10

	if a < 20 {
		fmt.Printf("a is less than 20\n")
	}
	fmt.Printf("value of a is : %d\n", a)
}

func IfElseExample() {
	var a int = 100
	if a < 20 {
		fmt.Printf("a is less than 20\n")
	} else {
		fmt.Printf("a is not less than 20\n")
	}
	fmt.Printf("value of a is : %d\n", a)
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
