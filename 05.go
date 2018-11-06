// 定数

package main

import "fmt"

func main() {
	const name = "norabal" // 定数宣言
	//name = "works" これをやろうとするとエラーになる。定数は変えられない。
	fmt.Println(name)

	const (
		// "iota"を使うと続く定数の値に連続した数値を登録できる
		sun = iota // 0
		mon        // 1
		tue        // 2
	)
	fmt.Println(sun, mon, tue)

	/**
	出力
	0 1 2
	*/
}
