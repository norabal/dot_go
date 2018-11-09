// 配列

package main

import "fmt"

func main() {
	var a [5]int // a[0] - a[4]
	fmt.Println(a)

	//配列の値をインデックスで指定して変更
	a[2] = 3
	a[4] = 10
	fmt.Println(a)

	//3つの配列に整数型の数字を登録
	b := [3]int{1, 3, 5}
	fmt.Println(b)

	//配列の数を指定しない場合
	c := [...]int{1, 3, 5}
	fmt.Println(c)
}
