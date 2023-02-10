package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	a1 := deck[0]
	b1 := deck[1]
	c1 := deck[2]
	fmt.Printf("Player 1: %v, %v, %v\n", a1, b1, c1)
	a2 := deck[3]
	b2 := deck[4]
	c2 := deck[5]
	fmt.Printf("Player 2: %v, %v, %v\n", a2, b2, c2)
	a3 := deck[6]
	b3 := deck[7]
	c3 := deck[8]
	fmt.Printf("Player 3: %v, %v, %v\n", a3, b3, c3)
	a4 := deck[9]
	b4 := deck[10]
	c4 := deck[11]
	fmt.Printf("Player 4: %v, %v, %v\n", a4, b4, c4)
}
