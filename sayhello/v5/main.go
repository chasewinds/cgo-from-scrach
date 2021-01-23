package main

import "C"

/*
#include <hello.h>
*/
import "C"

func main() {
	C.SayHello(C.CString("hello for go"))
}
