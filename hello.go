package main 
import (
"fmt"
"io"
"os"
"log"
"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Receive request path: %s\n", r.URL.Path)
	if r.URL.Path[1:] == "" {
		_, _ = fmt.Fprintf(w, "<h1>Hello, world!</h1>")
	} else {
		_, _ = fmt.Fprintf(w, "<h1>Hello, %s!</h1>", r.URL.Path[1:])
	}
}

func main() {
	fmt.Printf("Server started\n")
	http.HandleFunc("/", helloWorldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
