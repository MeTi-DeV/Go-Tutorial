package main

import "fmt"

func main() {
	//comment 2 : to define array in Go language should do like below 
	//it's Slice like Array but It growable but array has fixed length of values
	//always should define type of values in Slice like here that is String
	cards := []string{"first Card", newCard()}
	//comment 3 :for add new value into Slice use append() method
	//instade appended value to previous Slice
	//first argument of append is Slice name variable , second argument is new value that will be added to Slice
	cards = append(cards, "Appended Card")
	//comment 4 : for in Go language will define by first: index second:range of Slice
	for i, card := range cards {
		//comment 5 : print index of each elements and values
		fmt.Println(i, card)
	}
	
}
//comment 1 : create new func to return string value
func newCard() string {
	return "New Card"
}
