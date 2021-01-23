package main

/*
struct type { // c代码与go关键字冲突并不影响cgo调用
	int ID;
	int func;
	double _float;
};
*/
import "C"
import (
	"encoding/json"
	"fmt"
)

func main() {
	var cStruct C.struct_type
	cStruct.ID = 1
	cStruct._func = 1
	cStruct._float = 3.14159
	// cgo对应的c结构体，只有在C里大写的才能会被认为go里也是大写，也才会在marshal结果里，所以func字段不会存在
	if b, err := json.Marshal(cStruct); err == nil {
		fmt.Println(string(b)) // stdout: {"ID":1}
	}
	fmt.Println("ID is: ", cStruct.ID)
	fmt.Println("func is: ", cStruct._func)
	fmt.Println("_float is: ", cStruct._float)
}
