//go:generate stringer -type=Suit,Rank
package deck



type Suit int
const (
    Spade Suit = iota
    Diamonds
    Clubs
    Hearts
)

type Rank int
const (
     _ Rank = iota
     A
     Two
     Three
     Four
     Five
     Six
     Seven
     Eight
     Nine
     Ten
     J 
     Q
     K
)

type Card struct {
  Suit
  Rank
}

//type DeckCreation interface {
//  NewDeck() []Card
//}
