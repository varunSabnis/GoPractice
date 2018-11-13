package main

import (
       "fmt"
       "github.com/varunSabnis/GoPractice/Gophercises/deck"
)

func main(){
  // Passing multiple deck and shuffling options
  cards := deck.New(deck.Deck(3),deck.Shuffle())
  fmt.Println("Created a deck using constructor overloading like syntax by passing options and with help of New function")
  fmt.Println(cards)
}
