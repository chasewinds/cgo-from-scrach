package main

import "C"

/*
#include <stdio.h>
static void SayHello(const char* c) {
    puts(c);
}
*/
import "C"

func main() {
	C.SayHello(C.CString("hello, world"))
}
