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

	myfunc(1,1,1,1,1,1)

	//defer
	for i := 0; i < 5; i++ {
		defer f.Printf("%d ", i)
	}
}

//A+B と A*B を返します
func SumAndProduct(A, B int) (int, int) {
	return A+B, A*B
}

//可変長引数
func myfunc(arg ...int) {
	for _, n := range arg  {
		f.Printf("And the number is: %d\n", n)
	}
}