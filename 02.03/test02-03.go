package main

import ( f "fmt"

)

func main () {
	f.Print()

	var test byte
	//test = ([]byte)("test")
	test =200
	f.Print(test,"\n")

	//関数
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	f.Printf("%d + %d = %d\n", x, y, xPLUSy)
	f.Printf("%d * %d = %d\n", x, y, xTIMESy)

	//now 可変長引数
	for _, n := range arg  {
		f.Printf("And the number is: %d\n", n)
	}

}

//A+B と A*B を返します
func SumAndProduct(A, B int) (int, int) {
	return A+B, A*B
}

func myfunc(arg ...int) {}