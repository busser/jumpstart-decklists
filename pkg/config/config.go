package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/busser/jumpstart-decklists/pkg/mtg"
)

type config struct {
	Decks []deck `yaml:"decks"`
}

type deck struct {
	Name   string   `yaml:"name"`
	Colors string   `yaml:"colors"`
	Art    string   `yaml:"art"`
	Cards  []string `yaml:"cards"`
}

func ReadDecklists(r io.Reader) ([]mtg.Decklist, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("couldn't read decklists: %w", err)
	}

	var config config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	decklists := make([]mtg.Decklist, len(config.Decks))
	for i, deckConfig := range config.Decks {
		decklist, err := makeDecklist(deckConfig)
		if err != nil {
			return nil, err
		}

		decklists[i] = *decklist
	}

	return decklists, nil
}

func makeDecklist(config deck) (*mtg.Decklist, error) {
	colors := make([]mtg.Color, len(config.Colors))
	for i, colorSymbol := range config.Colors {
		color, err := makeColor(colorSymbol)
		if err != nil {
			return nil, err
		}

		colors[i] = color
	}

	cards := make([]mtg.DecklistItem, len(config.Cards))
	for i, cardConfig := range config.Cards {
		card, err := makeDecklistItem(cardConfig)
		if err != nil {
			return nil, fmt.Errorf("error parsing card %q: %w", cardConfig, err)
		}

		cards[i] = *card
	}

	var (
		art []byte
		err error
	)
	if config.Art != "" {
		art, err = os.ReadFile(config.Art)
		if err != nil {
			return nil, fmt.Errorf("error reading art %q: %w", config.Art, err)
		}
	}

	return &mtg.Decklist{
		Name:   config.Name,
		Colors: colors,
		Cards:  cards,
		Art:    art,
	}, nil
}

func makeColor(symbol rune) (mtg.Color, error) {
	switch symbol {
	case 'W':
		return mtg.ColorWhite, nil
	case 'U':
		return mtg.ColorBlue, nil
	case 'B':
		return mtg.ColorBlack, nil
	case 'R':
		return mtg.ColorRed, nil
	case 'G':
		return mtg.ColorGreen, nil
	default:
		return 0, fmt.Errorf("unknown color symbol %q", symbol)
	}
}

func makeDecklistItem(cardConfig string) (*mtg.DecklistItem, error) {
	parts := strings.Fields(cardConfig)

	if len(parts) < 3 {
		return nil, errors.New("need at least three parts (count, name, mana cost)")
	}

	count, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	name := strings.Join(parts[1:len(parts)-1], " ")

	rawManaCost := parts[len(parts)-1]
	manaCost := makeManaCost(rawManaCost)

	return &mtg.DecklistItem{
		Count: count,
		Card: mtg.Card{
			Name:     name,
			ManaCost: manaCost,
		},
	}, nil

}

const noManaCost = "-"

func makeManaCost(rawManaCost string) []mtg.ManaCostItem {
	if rawManaCost == noManaCost {
		return nil
	}

	manaCost := make([]mtg.ManaCostItem, len(rawManaCost))
	for i, symbol := range rawManaCost {
		item, err := makeManaCostItem(symbol)
		if err != nil {
			panic(err)
		}

		manaCost[i] = item
	}

	return manaCost
}

func makeManaCostItem(symbol rune) (mtg.ManaCostItem, error) {
	switch symbol {
	case 'W':
		return mtg.ManaCostWhite, nil
	case 'U':
		return mtg.ManaCostBlue, nil
	case 'B':
		return mtg.ManaCostBlack, nil
	case 'R':
		return mtg.ManaCostRed, nil
	case 'G':
		return mtg.ManaCostGreen, nil
	case 'C':
		return mtg.ManaCostColorless, nil
	case 'S':
		return mtg.ManaCostSnow, nil
	case '0':
		return mtg.ManaCost0, nil
	case '1':
		return mtg.ManaCost1, nil
	case '2':
		return mtg.ManaCost2, nil
	case '3':
		return mtg.ManaCost3, nil
	case '4':
		return mtg.ManaCost4, nil
	case '5':
		return mtg.ManaCost5, nil
	case '6':
		return mtg.ManaCost6, nil
	case '7':
		return mtg.ManaCost7, nil
	case '8':
		return mtg.ManaCost8, nil
	case '9':
		return mtg.ManaCost9, nil
	case 'X':
		return mtg.ManaCostX, nil
	default:
		return 0, fmt.Errorf("unknown mana cost symbol %q", symbol)
	}
}
