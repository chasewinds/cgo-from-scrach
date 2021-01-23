package main

//#include <hello.h>
import "C"
import "fmt"

// export SayHello
func SayHello(s *C.char) {
	fmt.Print(C.GoString(s))
}
