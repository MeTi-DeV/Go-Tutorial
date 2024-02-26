package main

import "fmt"

func main() {
	i, j := 20, 40
	var ipointer *int = &i
	var jpointer *int = &j
	fmt.Println("Address of i:", &i)
	fmt.Println("Address of i:", ipointer)
	fmt.Println("Address of i:", &j)
	fmt.Println("Address of i:", jpointer)
	i2 := i
	fmt.Println("value of i:", i)
	i2 += 2
	fmt.Println("value of i2:", i2)
	fmt.Println("Address of i:", &i2)
	var ipointer2 *int = &i
	*ipointer2 += 2
	fmt.Println("value of i:", *ipointer2)
	fmt.Println("value of i:", i)
	// changeNumberByValue(i)
	num := 20
	fmt.Println("Address of num:", &num)
	changeNumberByValue(num)
	fmt.Println("value of num after Change by value:", num)
	changeNumberByRef(&num)
	fmt.Println("value of num after Change by ref:", num)

}
func changeNumberByValue(num int) {
	fmt.Println("ChangeNumberByValue Address:", &num)
	num += 2
}
func changeNumberByRef(num *int) {
	fmt.Print("ChangeNumberByRef Address:", num)
	*num += 2
}
