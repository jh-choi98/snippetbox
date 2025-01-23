package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}
// HTTP handler function which is responsible for processing incoming
// HTTP requests and generating appropriate responses
// w http.ResponseWriter
// - is used to construct the HTTP response
// - writes data back to the client

// r *http.Request
// - represents the incoming HTTP request



func main() {
	// Use the http.NewServeMux() function to initialize a new
	// servemux, then register the home function as the handler for
	// the "/" URL pattern
	mux := http.NewServeMux() // creates a new ServeMux instance
	mux.HandleFunc("/", home) 
	// registers the home function to handle requests to the root
	// URL("/")

	// Use the http.ListenAndServe() function to start a new web
	// server. We pass in two parameters: the TCP network address to
	// listen on (in this case ":4000") and the servemux we just
	// created. If http.ListenAndServe() returns an error we use the
	// log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always
	// non-nil.
	log.Print("Starting server on :4000") 
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// A servemux (serve multiplexer) in Go is an HTTP request router that
// matches incoming HTTP requests to the appropriate handler functions
// based on the request's URL pattern

// TCP (Transmission Control Protocol)
// : A set of rules (a protocol) that computers use to communicate
// reliably over the internet or a network. It ensures that data sent
// from one computer to another arrives correctly and in order

// How TCP uses network addresses
// : when you send or receive data over a network, TCP uses 2 key
// pieces of addressing information
// 1. IP Address (where to send the data)
//  - a unique number assigned to each device on a network
//	ex. 192.168.1.10
// 2. Port Number (which service to talk to)
//  - a unique identifier for a specific application or service on a device

