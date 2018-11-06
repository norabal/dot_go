// 演算
/*
数値 + - * / %
文字列 +
論理値 AND(&&) OR(||) NOT(!)
*/

package main

import "fmt"

func main() {
	var x int
	x = 10 % 3
	fmt.Println(x)

	x += 3 // x = x + 3と同じ
	fmt.Println(x)

	x++ // xが1増える
	fmt.Println(x)

	//　文字列の連結
	var s string
	s = "hello " + "world"
	fmt.Println(s)

	a := true
	b := false
	fmt.Println(a && b) // aとbが両方true?
	fmt.Println(a || b) // aまたはbがtrue?
	fmt.Println(!a)     // aはtrueではない？

	/* 結果
	1
	4
	5
	hello world
	false
	true
	false
	*/
}
