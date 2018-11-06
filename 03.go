// 変数
// 変数名: 1文字目に注意
// なぜなら1文字目が小文字ならそのパッケージの中だけで見える変数、大文字の場合は他のパッケージからも見えてしまう変数になるから

package main

import "fmt"

func main() {
	// 宣言してから値を代入
	var msg1 string
	msg1 = "hello world"
	fmt.Println(msg1)

	// 宣言と代入を同時に行う
	// 型の指定はいらない
	var msg2 = "hello world"
	fmt.Println(msg2)

	// さらに短くすると varも省略できる
	msg3 := "hello world"
	fmt.Println(msg3)

	// 複数の変数を同時に型宣言
	var a, b int
	a, b = 10, 15
	fmt.Println(a, b)

	// 複数の変数を同時に型宣言しつつ値も入力
	a1, b1 := 10, 15
	fmt.Println(a1, b1)

	// 型の異なる変数をまとめて宣言することもできる
	var (
		c int
		d string
	)
	c = 20
	d = "hoge"
	fmt.Println(c, d)
}
