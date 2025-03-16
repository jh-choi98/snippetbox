package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Create a file server which serves files out of the
	// "./ui/static" directory. Note that the path given to the
	// http.Dir function is relative to the project directory root
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as
	// the handler for all URL paths that start with "/static". For
	// mathcing paths, we strip the "/static" prefix before the
	// request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Register the other application routes as normal
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}


/*
브라우저 URL 요청 → 서버 요청 처리 과정:
(1) http://localhost:4000/static/css/main.css
        ↓
(2) mux.Handle("/static/", ...) 를 통해 "/static" 경로 확인.
        ↓
(3) http.StripPrefix("/static", ...) 가 "/static" 제거 후, "/css/main.css"로 전달.
        ↓
(4) fileServer 핸들러가 "./ui/static/css/main.css" 경로에서 실제 파일을 찾아 응답.
*/
