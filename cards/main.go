package main

func main() {
	cards := deck{"first Card", newCard()}
	cards = append(cards, "Appended Card")
	cards.print()
}
func newCard() string {
	return "New Card"
}
