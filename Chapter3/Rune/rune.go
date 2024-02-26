package main

import (
	"fmt"
	"unsafe"
)

func main() {
	char1 := '😂'
	char2 := '😒'
	char3 := 128512
	fmt.Printf("char1 %c %T %d \n", char1, char1, char2)
	fmt.Printf("char2 %c %T %d \n", char2, char2, char2)
	fmt.Printf("char3 %c %T %d \n", char3, char3, char3)
	myStr := "Hello World 😉 😎"
	fmt.Printf("myStr: %s %T , len: %d , size: %d \n", myStr, myStr, len(myStr), unsafe.Sizeof(myStr))
	for i := 0; i < len(myStr); i++ {
		fmt.Printf("myStr[%d]: %c %T \n", i, myStr[i], myStr[i])
	}
	// rune is type that save as int32 runeSize Bytes > ByteSize
	myRune := []rune(myStr)
	fmt.Printf("myRune: %v %T , len: %d , size:%d \n", myRune, myRune, len(myRune), unsafe.Sizeof(myRune))
	for i := 0; i < len(myRune); i++ {
		fmt.Printf("myStr[%d]: %c %T \n", i, myRune[i], myRune[i])
	}
	myFarsiStr := "سلام دنیا"
	fmt.Printf("myFarsiStr: %s %T , len: %d , size: %d \n ", myFarsiStr, myFarsiStr, len(myFarsiStr), unsafe.Sizeof(myFarsiStr))
	for i := 0; i < len(myFarsiStr); i++ {
		fmt.Printf("myFarsiStr[%d]: %c %T \n", i, myFarsiStr[i], myFarsiStr[i])
	}
	myFarsiRune := []rune("سلام دنیا")
	fmt.Printf("myFarsiRune: %v %T , len: %d , size:%d \n", myFarsiRune, myFarsiRune, len(myFarsiRune), unsafe.Sizeof(myFarsiRune))
	for i := 0; i < len(myFarsiRune); i++ {
		fmt.Printf("myStr[%d]: %c %T \n", i, myFarsiRune[i], myFarsiRune[i])
	}

}
