// 基本的なデータ型
/*
string      "hello"
int         53
float64     10.2 ※ 少数点付きの数値の場合、float32もあるが通常はfloat64で良い
bool        false true
nil			NullやNoneのようなもの

[ 型宣言後の初期値それぞれ ]
var s string // ""
var a int // 0
var f bool // false
*/

package main

import "fmt"

func main() {
	a := 10
	b := 12.3
	c := "hoge"
	var d bool
	fmt.Printf("a: %d, b:%f, c:%s, d:%t\n", a, b, c, d)
}
