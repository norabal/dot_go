// 関数

package main

import "fmt"

/*
返り値msgの型指定をしている。pythonのタイプヒンティングとほぼ類似
*/
func hi(name string) string {
	// fmt.Println("hi!" + name)
	msg := "hi!" + name
	return msg
}

/*
自分にとってこれは新しい概念。
(msg string)とすることで、関数内で宣言されたすべての変数を返り値にできる
*/
func hi2(name string) (msg string) {
	// fmt.Println("hi!" + name)
	msg = "hi!" + name
	return
}

func main() {
	fmt.Println(hi("norabal"))
	fmt.Println(hi2("works"))
}
