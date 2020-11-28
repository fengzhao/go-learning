package http

import (
	"log"
	"net/http"
)

func main() {
	// 启动一个HTTP服务器，监听
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	http.HandleFunc("/", sayHelloWorld)

}
