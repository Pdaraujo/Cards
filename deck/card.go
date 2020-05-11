//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

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

//Options
func absRank(c Card) int {
	return int(c.Suit) * int(maxRank) + int(c.Rank)
}
