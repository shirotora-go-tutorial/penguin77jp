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
	fmt.Printf("%d %d %d\n",x,y,z)

	// 10個の要素を宣言します。要素の型はbyteの配列です。
	var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// byteを含む２つのsliceを宣言します
	var a, b []byte

	// aポインタ配列の3つ目の要素から始まり、5つ目の要素まで
	a = ar[2:5]
	//現在aの持つ要素は：ar[2]、ar[3]とar[4]です。

	// bは配列arのもう一つのsliceです。
	b = ar[3:5]
	// bの要素は：ar[3]とar[4]です。
	fmt.Printf("%c %c\n",a,b)

	// keyを文字列で宣言します。値はintとなるディクショナリです。
	// この方法は使用される前にmakeで初期化される必要があります。
	/* var numbers map[string]int */
	// もうひとつのmapの宣言方法
	numbers := make(map[string]int)
	numbers["one"] = 1  //代入
	numbers["ten"] = 10 //代入
	numbers["three"] = 3

	fmt.Println("３つ目の数字は: ", numbers["three"]) // データの取得
	fmt.Println("\n")
	// "３つ目の数字は： 3"という風に出力されます。



	// ディクショナリを初期化します。
	rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
	// mapは２つの戻り値があります。２つ目の戻り値では、もしkeyが存在しなければ、okはfalseに、存在すればokはtrueになります。
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C")  // keyがCの要素を削除します。
}