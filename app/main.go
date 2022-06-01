// Source: https://gobyexample.com/http-servers

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html");
	fmt.Fprintf(w, "<!DOCTYPE html>\n");
	fmt.Fprintf(w, "<html>\n");
	fmt.Fprintf(w, "<head>\n");
	fmt.Fprintf(w, "<meta charset=\"utf-8\" />\n");
	fmt.Fprintf(w, "<title>Hello, World</title>\n");
	fmt.Fprintf(w, "</head>\n");
	fmt.Fprintf(w, "<body>\n");
	fmt.Fprintf(w, "Ohai\n")
	fmt.Fprintf(w, "</body>\n");
	fmt.Fprintf(w, "</html>");
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting on port :8090...")
	http.ListenAndServe(":8090", nil)
}
