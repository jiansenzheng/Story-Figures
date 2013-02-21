package main

import (
	"fmt"
	// "log"
	"iflytek.com/mongotxt"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Heelo yfyang!")
}

// func main() {
// 	http.HandleFunc("/", sayHelloName)
// 	err := http.ListenAndServe(":9090", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe:", err)
// 	}
// }
func main() {
	mongotxt.ScanImportMongo("/Users/yfyang/Documents/小说/", "127.0.0.1:5565")
}
