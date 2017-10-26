package main

import(
	"fmt"
	"errors"
)


func main () {
	s := "hello"
	c := []byte(s)  // 文字列 s を []byte 型にキャスト
	c[0] = 'c'
	s2 := string(c)  // もう一度 string 型にキャストし直す
	fmt.Printf("%s\n", s2)

	s = "hello"
	s = "c" + s[1:] // 文字列を変更することはできませんが、スライスは行えます。
	fmt.Printf("%s\n", s)

	fmt.Printf("raw型\n");
	m := `hello
    world`
	fmt.Printf("%s\n", m)

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}

	//iota列挙型
	fmt.Printf("\niota列挙型\n")
	const(
		x = iota  // x == 0
		y = iota  // y == 1
		z = iota  // z == 2
		w  // 定数の宣言で値を省略した場合、デフォルト値は前の値と同じになります。
		// ここではw = iotaと宣言していることと同じになりますので、
		// w == 3となります。実は上のyとzでもこの"= iota"は省略することができます。
	)

	const v = iota // constキーワードが出現する度に、iotaは置き直されます。ここではv == 0です。

	const (
		e, f, g = iota, iota, iota //e=0,f=0,g=0 iotaの同一行は同じです
	)
	fmt.Printf("%d %d %d",x,y,z)

	//next スライス


}