// ポインタ
// アドレスを指し示す変数
// 演算はできない
// メモリの節約ができるなど、効率的なプログラムにつながる

package main

import "fmt"

func main() {
	a := 5
	var pa *int // int 型を格納する領域のアドレスを格納する準備
	pa = &a     // &a = aのアドレス
	// paの領域にあるデータの値 = *pa
	fmt.Println(pa)
	fmt.Println(*pa) // ポインタのアドレスから値を出力する方法

	/* 結果
	0xc0000160b0
	5
	*/
}
