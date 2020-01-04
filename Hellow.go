package main

import (
	"fmt"
	"math"
)

func main() {
	//dynamicTypeDeclaration()
	//mixedVariableDeclaration()
	//escapeSequences()
	//constKeyWord()
	//simpleIf()
	//IfElseExample()
	//switchCase(90)
	//forLoopTest()
	//nestedLoops()
	//breakStatement()
	//continueStatement()
	//goTo()
	//maxValueTest()
	//callByValueTest()
	//callByReferenceTest()
	//funcAsValue()

}
func functionClosureTest() {
	//nextNumber here is a function that i as 0
	nextNumber := getSequence()
	//so we increase i by 1 and return same

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	//so creating new sequence
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func funcAsValue() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Print(getSquareRoot(9))
}
func callByReferenceTest() {
	var a int = 100
	var b int = 200
	fmt.Printf("Before swap , value of a : %d\n", a)
	fmt.Printf("Before swap , value of b : %d\n", b)

	/* calling a function to swap the values.
	 * &a indicates pointer to a ie. address of variable a and
	 * &b indicates pointer to b ie. address of variable b.
	 */
	swap1(&a, &b)
	fmt.Printf("After swap , value of a : %d\n", a)
	fmt.Printf("After swap , value of b : %d\n", b)
}
func callByValueTest() {
	var a int = 100
	var b int = 200
	fmt.Printf("Before swap , value of a : %d\n", a)
	fmt.Printf("Before swap , value of b : %d\n", b)

	swap(a, b)
	fmt.Printf("After swap , value of a : %d\n", a)
	fmt.Printf("After swap , value of b : %d\n", b)
	//It shows that there is no change in the values though they had been changed inside the function.
}
func swap(x int, y int) int {
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}

func swap1(x *int, y *int) int {
	var temp int
	temp = *x
	*x = *y
	*y = temp
	return temp
}
func maxValueTest() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)
	fmt.Printf("Max Value is %d\n", ret)
}
func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}

}
func goTo() {
	var a int = 10

LOOP:
	for a <= 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}
func continueStatement() {
	var a int = 10
	for a <= 20 {
		if a == 15 {
			a = a + 1
			continue
		}
		fmt.Printf("value of a %d\n", a)
		a++
	}
}
func breakStatement() {
	var a int = 10
	for a < 20 {
		fmt.Printf("value of a: %d\n", a)
		a++
		if a > 15 {
			break
		}

	}
}
func nestedLoops() {
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j < (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d is prime\n", i)
		}
	}
}
func forLoopTest() {
	var b = 15
	var a int
	numbers := [6]int{1, 2, 3, 5}
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a %d\n", a)
	}
	for a < b {
		a++
		fmt.Printf("value of a : %d\n", a)
	}
	for i, x := range numbers {
		fmt.Printf("value of x %d at %d\n", x, i)
	}
}
func selectStatement() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("recieved ", i1, "from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := <-c3:
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
	var grade = ""
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
	var a = 10

	if a < 20 {
		fmt.Printf("a is less than 20\n")
	}
	fmt.Printf("value of a is : %d\n", a)
}

func IfElseExample() {
	var a = 100
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
	var x = 20.0

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
