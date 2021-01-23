package main

/*
#include<string.h>
char arrInC[10] = {'a', '7'};
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var bytesInGo []byte
	// 把一个go的byte数组转换成可以容纳c char数组的类型
	var bytesHeader = (*reflect.SliceHeader)(unsafe.Pointer(&bytesInGo))
	// 从c数组读取数据（因为c数组是数组头的指针，所以读[0]即可）
	bytesHeader.Data = uintptr(unsafe.Pointer(&C.arrInC[0]))
	bytesHeader.Cap = 10
	bytesHeader.Len = 10
	fmt.Println("c arr read in go by way 1: ", string(bytesInGo))

	// 另一种读取方式
	anotherBytesInGo := (*[31]byte)(unsafe.Pointer(&C.arrInC[0]))[:10:10]
	var x []byte
	for _, e := range anotherBytesInGo {
		x = append(x, e)
	}
	fmt.Println("c arr read in go by way 2: ", string(bytesInGo))
}
