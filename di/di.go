package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	Greet(os.Stdout, "Asad")
	http.ListenAndServe(":3000", http.HandlerFunc(GreetHandler))
}

func GreetHandler(w http.ResponseWriter, _ *http.Request) {
	Greet(w, "G")
}

// Greet will accept any concrete type which implements the the io.Writer interface
// i.e. http.ResponseWriter, os.Stdout or out bytes.Buffer from unit tests
// We are actually using the principle of DI to decouple the things and test our code easily.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}
