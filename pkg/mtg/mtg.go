package mtg

type Color rune

const (
	ColorWhite = 'W'
	ColorBlue  = 'U'
	ColorBlack = 'B'
	ColorRed   = 'R'
	ColorGreen = 'G'
	// TODO(busser): support colorless
)

type ManaCostItem rune

const (
	ManaCostWhite = ColorWhite
	ManaCostBlue  = ColorBlue
	ManaCostBlack = ColorBlack
	ManaCostRed   = ColorRed
	ManaCostGreen = ColorGreen
	ManaCost1     = '1'
	ManaCost2     = '2'
	ManaCost3     = '3'
	ManaCost4     = '4'
	ManaCost5     = '5'
	ManaCost6     = '6'
	ManaCost7     = '7'
	ManaCost8     = '8'
	ManaCost9     = '9'
	ManaCostX     = 'X'
	// TODO(busser): support 0, 10+, colorless, and hybrid
)

type Card struct {
	Name     string
	ManaCost []ManaCostItem
}

type Decklist struct {
	Name   string
	Colors []Color
	Cards  []DecklistItem
	Art    []byte
}

type DecklistItem struct {
	Count int
	Card  Card
}
