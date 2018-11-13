package main

import (
       "fmt"
       "github.com/varunSabnis/GoPractice/Gophercises/deck"
       "math/rand"
       "time"
       "os"
)


//BlackJack rules -

//Deal Cards : Two cards are given to you face up. The dealer gets one card face up and the other is face down.
func DealCards(CardDeck []deck.Card, DeckPos *int)([]deck.Card,[]deck.Card){
   var user_cards []deck.Card
   var dealer_cards []deck.Card
   user_cards = append(user_cards,CardDeck[*DeckPos])
   user_cards = append(user_cards,CardDeck[*DeckPos + 1])
   dealer_cards = append(dealer_cards,CardDeck[*DeckPos + 2])
   dealer_cards = append(dealer_cards,CardDeck[*DeckPos + 3])
   *DeckPos = *DeckPos + 4
   return user_cards, dealer_cards
}

//Hit : If we hit, then we get to pick a card.
 //Now, if we bust(more than 21) , we loose. If we get 21 you win.
 func Hit(cardDeck []deck.Card,DeckPos *int) deck.Card{
      hitCard := cardDeck[*DeckPos]
      *DeckPos = *DeckPos + 1
      return(hitCard)
}

// Get score of card in hand
func getScore(handCards []deck.Card) int{
   score := getMinScore(handCards)
   if(score > 11){
     return(score)
   }
   for _,cd := range(handCards){
       if(cd.Rank == 1){
         score = score + 10
       }
   }
   //fmt.Println("score in function",score)
   return(score)
}

func min(a int, b int) int {
  if(a<b){
    return(a)
  }else{
    return(b)
  }
}

func getMinScore(handCards []deck.Card) int{
  var score = 0
  for _,cd := range(handCards){

    score = score + min(int(cd.Rank),10)
    //fmt.Println("score add",score)
  }
  //fmt.Println("Min score",score)
  return(score)
}

/*
Stand : If we play stand , the dealer will reveal his hand (hidden card) and will always hit if they have 16 or lower. They will
stop hitting if they have 17 or more.
  Soft 17 rule - if delaer has A + 6 , he will hit until he gets a hard 17
  Now, if dealer has more points than you, he will win. If dealer busts or has lesser points than you , you win.
*/

func Stand(CardDeck []deck.Card, DeckPos *int,dealer_cards []deck.Card) int {
    var score int
    var hitCard deck.Card
    for ((getScore(dealer_cards)<=16) || (getScore(dealer_cards)==17 && getMinScore(dealer_cards)!=17) ) {
      hitCard = Hit(CardDeck,DeckPos)
      dealer_cards = append(dealer_cards,hitCard)
      fmt.Println("dealer cards after hit\n")
      PrintCards(dealer_cards,1)
   }
  score = getScore(dealer_cards)
  return(score)
}
/*
Split : It involves spliting the bet when you get 2 cards of same value. so you get 2 cards, each, for one card already in hand
(so basically you have 2 hands to play or 2 chances to win/loose). Your bet also gets doubled.

Double : Your bet doubles and can get only one more card.

Surrender : You loose have your bet to the dealer
*/

func PrintCards(hand []deck.Card, show_all int){

if(show_all == 1){
for _,cd := range(hand){
  fmt.Println(cd.Suit, cd.Rank)
 }
}else{
  fmt.Println(hand[0].Suit, hand[0].Rank)
}
fmt.Println("\n")
}



func BlackJack(CardDeck []deck.Card){
fmt.Println("\n ************** Let's Play BlackJack *************************  \n")
var DeckPos int
var choice string
DeckPos = 0
var playerScore int
var DealerScore int
user_cards, dealer_cards := DealCards(CardDeck,&DeckPos)
fmt.Println("Cards Dealt to Player : ")
PrintCards(user_cards,1)
fmt.Println("Cards dealt to Dealer : ")
PrintCards(dealer_cards,0)
var hitCard deck.Card
if (getScore(user_cards) == 21){
    fmt.Printf("BlackJack ! You Win!!\n")
  }else{
    fmt.Println("Enter your choice of Action 1.s 2.h 3.e \n")
    fmt.Scanf("%s\n",&choice)
    fmt.Println("choice",choice)
  for(choice != "s"){
     if(choice == "h"){
       hitCard = Hit(CardDeck,&DeckPos)
       user_cards = append(user_cards,hitCard)
       PrintCards(user_cards,1)
       playerScore = getScore(user_cards)
       fmt.Println("Player Score",playerScore)
       if(playerScore > 21){
           fmt.Println("You are busted!..you loose\n")
           os.Exit(0)
       }
       if(playerScore == 21){
         fmt.Println("You win!!")
         os.Exit(0)
       }
     }
     fmt.Println("Enter your choice of Action 1.s 2.h 3.e \n")
     fmt.Scanf("%s\n",&choice)
     fmt.Println("choice",choice)
   }
   fmt.Println("Dealers turn :\n Dealers Cards\n")
   PrintCards(dealer_cards,1)
   DealerScore = Stand(CardDeck,&DeckPos,dealer_cards)
   fmt.Println("dealer score",DealerScore)
   playerScore = getScore(user_cards)
   fmt.Println("player score",playerScore)
  if(DealerScore > 21){
    fmt.Println("Dealer busted ...You win!!")
    os.Exit(0)
  }
  if(playerScore >= DealerScore){
    fmt.Println("You win!!")
  }
  if(playerScore < DealerScore){
    fmt.Println("You Loose!!")
  }
 }
}


type MyCard struct {
  deck.Card
}

type DeckCreation interface {
   NewDeck() []deck.Card
   ShuffleDeck([]deck.Card)
   MultipleDeck([]deck.Card,int) [][]deck.Card
   //SortDeck([]deck.Card)
}


func (c MyCard) MultipleDeck(CardDeck []deck.Card, count int) [][]deck.Card {
   var DecksOfDeck [][]deck.Card
   for i := 0;i<count;i++{
     DecksOfDeck = append(DecksOfDeck,CardDeck)
   }
   return(DecksOfDeck)
}


func (c MyCard) ShuffleDeck(cardDeck []deck.Card){
    r := rand.New(rand.NewSource(time.Now().Unix()))
    //shuffledDeck := make([]deck.Card,len(CardDeck))
    perm := r.Perm(len(cardDeck))
    for i,randomIdx := range(perm){
      cardDeck[i] = cardDeck[randomIdx]
    }
}


func (c MyCard) NewDeck() []deck.Card {
var cardDeck []deck.Card
//var rk string
card := deck.Card{}
for i := 0;i<4;i++ {
  card.Rank = deck.Card{}.Rank + 1
  for j := 0;j<13;j++ {
     //fmt.Println(card.Rank)
     cardDeck = append(cardDeck,card)
     card.Rank = card.Rank + 1
  }
  card.Suit = card.Suit + 1
}
return(cardDeck)
}


func main(){

  var i int
  card := deck.Card{}
  mycard := MyCard{card}
  cardDeck := mycard.NewDeck()
  //fmt.Println("Card Deck \n ",cardDeck)
  //shuffledDeck := mycard.ShuffleDeck(cardDeck)
  //fmt.Println("Shuffled Card Deck\n",shuffledDeck)
  DecksOfDeck := mycard.MultipleDeck(cardDeck,1)
  for i=0;i<len(DecksOfDeck);i++{
   mycard.ShuffleDeck(DecksOfDeck[i])
  }
  //fmt.Println("Shuffled Deck in DecksOfDeck",DecksOfDeck[0])
  BlackJack(DecksOfDeck[0])
}
