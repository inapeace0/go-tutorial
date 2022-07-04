package main

import "fmt"

import "rsc.io/quote"

func main() {
	//	Print the text
	fmt.Println("Hello, World!")
	//	Call the module
	fmt.Println(quote.Go())

	//	Declare and initialize the variable
	var x float64 = 20
	y := 42

	var a, b, c = 3, 031, 0x23
	fmt.Println(x, y)
	fmt.Printf("x is of type %T %T\n", x, y)
	fmt.Println(a,b,c)

	//	Floating-point literal
	var f1 = 3.1234
	var f2 = 3234E-4

	fmt.Println(f1, f2)

	//	String literals
	fmt.Printf("\n%T\t%T\n", f1, f2)

	fmt.Printf("hello, de, 'Color', \"string\"\n");

	//	Constant
	const WIDTH int = 10
	const HEIGHT int = 5

	var area int
	area = WIDTH * HEIGHT

	fmt.Printf("value of area: %d\n", area);

	//	Bitwise operator

	area = area << 1

	//	Pointer

	fmt.Printf("%d %x\n", area, &area)

	var ptr *int

	ptr = &area

	fmt.Printf("%x %d\n", ptr, *ptr)

}