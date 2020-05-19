package deck

import (
	"fmt"
	"math/rand"
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

func TestDefaultSort(t *testing.T) {
	cards := NewDeck(DefaultSort)
	card := Card{
		Suit: Spade,
		Rank: Ace,
	}
	if cards[0] != card {
		t.Error("Expected Ace of Spades, Received: ", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := NewDeck(Sort(Less))
	card := Card{
		Suit: Spade,
		Rank: Ace,
	}
	if cards[0] != card {
		t.Error("Expected Ace of Spades, Received: ", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := NewDeck(Jokers(3))
	counter := 0
	for _, c := range cards {
		if c.Suit == Joker {
			counter++
		}
	}

	if counter != 3 {
		t.Error("Expected 3 Jokers, Received: ", counter)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := NewDeck(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("There should be no Two's or Three's")
		}
	}
}

func TestDecks(t *testing.T) {
	cards := NewDeck(Decks(3))
	if len(cards) != 13 * 4 * 3 {
		t.Errorf("Expected %d cards and received %d cards", 13 * 4 * 3, len(cards))
	}
}

func TestShuffle2(t *testing.T) {
	//Make shuffle rand deterministic
	//First call to shuffleRand.Perm(52 should be :
	// [40, 35 ...]
	shuffleRand = rand.New(rand.NewSource(0))
	originalDeck := NewDeck()
	first := originalDeck[40]
	second := originalDeck[35]
	cars := NewDeck(Shuffle2)
	if cars[0] != first {
		t.Errorf("Expected first card to be %s, received %s.", first, cars[0])
	}

	if cars[1] != second {
		t.Errorf("Expected first card to be %s, received %s.", second, cars[1])
	}

}
