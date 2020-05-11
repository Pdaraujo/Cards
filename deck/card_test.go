package deck

import (
	"fmt"
	"testing"
)

func ExampleCard_String() {
	fmt.Println(Card{
		Suit: Heart,
		Rank: Ace,
	})
	fmt.Println(Card{
		Suit: Club,
		Rank: Three,
	})
	fmt.Println(Card{
		Suit: Diamond,
		Rank: Seven,
	})
	fmt.Println(Card{
		Suit: Spade,
		Rank: King,
	})
	fmt.Println(Card{
		Suit: Joker,
	})

	//Output:
	//Ace of Hearts
	//Three of Clubs
	//Seven of Diamonds
	//King of Spades
	//Joker
}

func TestNewDeck(t *testing.T) {
	cards := NewDeck()
	if len(cards) != 52 {
		t.Errorf("Wrong number of cards in new deck")
	}
}
