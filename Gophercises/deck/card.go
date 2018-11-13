//go:generate stringer -type=Suit,Rank
package deck

import (
  "math/rand"
  "time"
  "fmt"
)


type Suit int
const (
    Spade Suit = iota
    Diamonds
    Clubs
    Hearts
)

var suits = [] Suit{Spade,Diamonds,Clubs,Hearts}

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

func Deck(n int) func([]Card) []Card {
      return func(cards []Card) []Card {
           for i:=0;i<n;i++{
             for _,cd := range(cards){
             cards = append(cards,cd)
           }
         }
          return(cards)
      }
}

func Shuffle() func([]Card) []Card{
   return func(cards []Card) []Card {
     var shuffledCards = make([]Card,len(cards))
     r := rand.New(rand.NewSource(time.Now().Unix()))
     //shuffledDeck := make([]deck.Card,len(CardDeck))
     perm := r.Perm(len(cards))
     for i,randomIdx := range(perm){
       shuffledCards[i] = cards[randomIdx]
     }
     return(shuffledCards)
   }
}

func New(opts ...func([]Card) []Card) []Card {
  var cards []Card
  for _,suit := range suits{
      for rank:=1;rank<=13;rank++{
        cards = append(cards,Card{Rank:Rank(rank),Suit:suit})
      }
    }
  fmt.Println("options",opts)
  for _,opt := range(opts){
    cards = opt(cards)
  }
  return(cards)
}
