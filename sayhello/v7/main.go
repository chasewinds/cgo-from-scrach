// +build go1.10

package main

// #cgo !windows LDFLAGS: -lm
//void SayHello(_GoString_ s); // _GoString_ is native go str type
import "C"
import "fmt"

//export SayHello
func SayHello(s string) {
	fmt.Println(s)
}

func main() {
	C.SayHello("hey from go by native string")
}
