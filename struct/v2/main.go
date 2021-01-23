package main

/*
enum chess {
	Queen,
	King,
	Knight,
	Pawn,
};
*/
import "C"
import "fmt"

func main() {
	var queen C.enum_chess = C.Queen
	var king C.enum_chess = C.gKing
	var pawn C.enum_chess = C.Pawn
	var knight C.enum_chess = C.Knight
	fmt.Println(queen)
	fmt.Println(king)
	fmt.Println(pawn)
	fmt.Println(knight)
}
