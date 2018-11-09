// 関数

package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	//関数実行
	fmt.Println(swap(5, 2))

	//関数を変数に代入
	f := func(a, b int) (int, int) {
		return b, a
	}
	fmt.Println(f(2, 3))

	//即時関数 使い捨ての関数
	func(msg string) {
		fmt.Println(msg)
	}("norabal")

}
