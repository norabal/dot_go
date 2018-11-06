// webサーバー

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path[1:]では"http://localhost:8080"以降が入る
	fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

/*
http://localhost:8080/でアクセスするとhandler()で指定したレスポンスが返る
*/
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
