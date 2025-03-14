package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Card represents a playing card
type Card struct {
	Name  string
	Value int
}

// Game represents the state of a blackjack game
type Game struct {
	PlayerHand []Card
	DealerHand []Card
	Deck       []Card
	Reader     *bufio.Reader
}

// NewGame creates and initializes a new blackjack game
func NewGame() *Game {
	game := &Game{
		Reader: bufio.NewReader(os.Stdin),
	}
	game.initializeDeck()
	return game
}

// initializeDeck creates a standard deck of cards
func (g *Game) initializeDeck() {
	cardValues := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}

	g.Deck = []Card{}
	for name, value := range cardValues {
		g.Deck = append(g.Deck, Card{Name: name, Value: value})
	}
}

// drawCard draws a random card from the deck
func (g *Game) drawCard() Card {
	cardIndex := rand.Intn(len(g.Deck))
	card := g.Deck[cardIndex]
	return card
}

// calculateTotal calculates the total value of a hand, accounting for aces
func calculateTotal(hand []Card) int {
	total := 0
	aceCount := 0

	for _, card := range hand {
		total += card.Value
		if card.Name == "ace" {
			aceCount++
		}
	}

	// Adjust for aces if needed
	for aceCount > 0 && total > 21 {
		total -= 10 // Convert an ace from 11 to 1
		aceCount--
	}

	return total
}

// suggestMove suggests the best move based on player's hand and dealer's visible card
func suggestMove(playerHand []Card, dealerVisibleCard Card) string {
	// Check for pair of aces
	if len(playerHand) == 2 && playerHand[0].Name == "ace" && playerHand[1].Name == "ace" {
		return "P" // Always split aces
	}

	playerTotal := calculateTotal(playerHand)
	dealerCardValue := dealerVisibleCard.Value

	switch {
	case playerTotal == 21:
		if dealerCardValue == 11 || dealerCardValue == 10 {
			return "S" // Stand if dealer might have Blackjack
		}
		return "W" // Automatic win

	case playerTotal >= 17 && playerTotal <= 20:
		return "S" // Always stand on 17-20

	case playerTotal >= 12 && playerTotal <= 16:
		if dealerCardValue >= 7 {
			return "H" // Hit if dealer has a strong card
		}
		return "S" // Otherwise, stand

	default:
		return "H" // Always hit if total is 11 or lower
	}
}

// displayHand displays the cards in a hand and their total value
func displayHand(hand []Card, hideSecondCard bool) {
	fmt.Print("Cards: ")

	for i, card := range hand {
		if i > 0 && hideSecondCard {
			fmt.Print("[hidden]")
			break
		}
		fmt.Printf("%s ", card.Name)
	}

	if !hideSecondCard {
		fmt.Printf("(Total: %d)", calculateTotal(hand))
	}
	fmt.Println()
}

// playerTurn handles the player's turn
func (g *Game) playerTurn() (bool, error) {
	for {
		fmt.Printf("\nDo you want to (H)it, (S)tand, or (Q)uit? ")
		input, err := g.Reader.ReadString('\n')
		if err != nil {
			return false, fmt.Errorf("error reading input: %v", err)
		}

		input = strings.TrimSpace(strings.ToUpper(input))

		switch input {
		case "Q":
			fmt.Println("Thanks for playing! Exiting game...")
			return false, nil

		case "S":
			fmt.Println("You chose to Stand.")
			return true, nil

		case "H":
			newCard := g.drawCard()
			g.PlayerHand = append(g.PlayerHand, newCard)
			fmt.Printf("You drew: %s\n", newCard.Name)

			playerTotal := calculateTotal(g.PlayerHand)
			fmt.Printf("Your new total: %d\n", playerTotal)

			if playerTotal > 21 {
				fmt.Println("You busted! Dealer wins.")
				return false, nil
			}
		default:
			fmt.Println("Invalid input. Please enter H, S, or Q.")
		}
	}
}

// dealerTurn handles the dealer's turn
func (g *Game) dealerTurn() {
	fmt.Println("\nDealer's turn:")
	fmt.Println("Dealer reveals second card...")
	displayHand(g.DealerHand, false)

	dealerTotal := calculateTotal(g.DealerHand)
	for dealerTotal < 17 {
		fmt.Println("Dealer hits...")
		newCard := g.drawCard()
		g.DealerHand = append(g.DealerHand, newCard)
		fmt.Printf("Dealer draws: %s\n", newCard.Name)

		dealerTotal = calculateTotal(g.DealerHand)
		fmt.Printf("Dealer total: %d\n", dealerTotal)
	}
}

// determineWinner determines the winner of the game
func (g *Game) determineWinner() {
	playerTotal := calculateTotal(g.PlayerHand)
	dealerTotal := calculateTotal(g.DealerHand)

	fmt.Printf("\nFinal Scores - You: %d, Dealer: %d\n", playerTotal, dealerTotal)

	if playerTotal > 21 {
		fmt.Println("You busted! Dealer wins.")
	} else if dealerTotal > 21 {
		fmt.Println("Dealer busted! You win!")
	} else if playerTotal > dealerTotal {
		fmt.Println("You win!")
	} else if playerTotal < dealerTotal {
		fmt.Println("Dealer wins!")
	} else {
		fmt.Println("It's a tie!")
	}
}

// playOneRound plays one round of blackjack
func (g *Game) playOneRound() bool {
	// Reset hands
	g.PlayerHand = []Card{g.drawCard(), g.drawCard()}
	g.DealerHand = []Card{g.drawCard(), g.drawCard()}

	// Display initial hands
	fmt.Println("\n--- New Round ---")
	fmt.Print("Your ")
	displayHand(g.PlayerHand, false)
	fmt.Print("Dealer's ")
	displayHand(g.DealerHand, true)

	// Determine suggested move
	move := suggestMove(g.PlayerHand, g.DealerHand[0])
	fmt.Println("Suggested move:", move)

	// Player's turn
	continueToDealer, err := g.playerTurn()
	if err != nil {
		fmt.Println("Error during player's turn:", err)
		return false
	}

	if continueToDealer {
		// Dealer's turn
		g.dealerTurn()
		g.determineWinner()
	}

	// Ask to play again
	fmt.Print("\nDo you want to play again? (Y/N): ")
	input, err := g.Reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return false
	}

	input = strings.TrimSpace(strings.ToUpper(input))
	return input == "Y"
}

// play starts the game
func (g *Game) play() {
	fmt.Println("Welcome to CLI Blackjack!")

	for {
		if !g.playOneRound() {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create and start the game
	game := NewGame()
	game.play()
}
