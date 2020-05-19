//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String() )
}

func NewDeck(opt ...CardOptions) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{
				Suit: suit,
				Rank: rank,
			})
		}
	}
	for _, option := range opt {
		cards = option(cards)
	}
	return cards
}

//Filter
func Filter(f func(card Card) bool) func ([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

//Options



//Shuffle
type Permer interface {
	Perm(n int) []int
}

func Shuffle(p Permer) func ([]Card) []Card {
	return func(cards []Card) []Card {
		ret := make([]Card, len(cards))
		perm := p.Perm(len(cards))
		for i, j := range perm {
			ret[i] = cards[j]
		}
		return ret
	}
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

func Shuffle2(cards []Card) []Card {
	ret := make([]Card, len(cards))
	perm := shuffleRand.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

func Sort(less func(cards []Card) func(i, j int) bool) CardOptions {
	return func (cards []Card) []Card {
		sort.SliceStable(cards, less(cards))
		return cards
	}
}

func DefaultSort(cards []Card) []Card {
	sort.SliceStable(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

//Jokers
func Jokers(n int) func ([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Joker,
				Rank(i),
			})
		}
		return cards
	}
}

//Decks
func Decks(n int) func ([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}

//Helpers
func absRank(c Card) int {
	return int(c.Suit) * int(maxRank) + int(c.Rank)
}
