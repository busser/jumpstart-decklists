package render

import (
	"bytes"
	"embed"
	"encoding/base64"
	"fmt"
	"text/template"

	"github.com/busser/jumpstart-decklists/pkg/mtg"
)

var (
	//go:embed templates
	templatesFS embed.FS
	templates   = template.Must(template.ParseFS(templatesFS, "templates/*"))
)

func DecklistsAsWebPage(decklists []mtg.Decklist) ([]byte, error) {
	data := struct {
		Decklists []decklistData
	}{}

	data.Decklists = make([]decklistData, len(decklists))
	for i, decklist := range decklists {
		data.Decklists[i] = makeDecklistData(decklist)
	}

	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, "page.html", data); err != nil {
		return nil, fmt.Errorf("failed to render decklist: %w", err)
	}

	return buf.Bytes(), nil
}

func DecklistAsHTML(decklist mtg.Decklist) ([]byte, error) {
	data := makeDecklistData(decklist)

	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, "decklist.html", data); err != nil {
		return nil, fmt.Errorf("failed to render decklist: %w", err)
	}

	return buf.Bytes(), nil
}

type decklistData struct {
	Name      string
	Colors    []string
	Cards     []decklistItemData
	ArtBase64 string
}

type decklistItemData struct {
	Count int
	Card  cardData
}

type cardData struct {
	Name     string
	ManaCost []string
}

func makeDecklistData(decklist mtg.Decklist) decklistData {
	data := decklistData{
		Name:      decklist.Name,
		Colors:    make([]string, len(decklist.Colors)),
		Cards:     make([]decklistItemData, len(decklist.Cards)),
		ArtBase64: base64.StdEncoding.EncodeToString(decklist.Art),
	}

	for i, color := range decklist.Colors {
		data.Colors[i] = mtgColorToString(color)
	}

	for i, item := range decklist.Cards {
		data.Cards[i] = decklistItemData{
			Count: item.Count,
			Card: cardData{
				Name:     item.Card.Name,
				ManaCost: make([]string, len(item.Card.ManaCost)),
			},
		}

		for j, manaCostItem := range item.Card.ManaCost {
			data.Cards[i].Card.ManaCost[j] = mtgManaCostItemToString(manaCostItem)
		}
	}

	return data
}

func mtgColorToString(color mtg.Color) string {
	switch color {
	case mtg.ColorWhite:
		return "w"
	case mtg.ColorBlue:
		return "u"
	case mtg.ColorBlack:
		return "b"
	case mtg.ColorRed:
		return "r"
	case mtg.ColorGreen:
		return "g"
	default:
		return ""
	}
}

func mtgManaCostItemToString(item mtg.ManaCostItem) string {
	switch item {
	case mtg.ManaCostWhite:
		return "w"
	case mtg.ManaCostBlue:
		return "u"
	case mtg.ManaCostBlack:
		return "b"
	case mtg.ManaCostRed:
		return "r"
	case mtg.ManaCostGreen:
		return "g"
	case mtg.ManaCost1:
		return "1"
	case mtg.ManaCost2:
		return "2"
	case mtg.ManaCost3:
		return "3"
	case mtg.ManaCost4:
		return "4"
	case mtg.ManaCost5:
		return "5"
	case mtg.ManaCost6:
		return "6"
	case mtg.ManaCost7:
		return "7"
	case mtg.ManaCost8:
		return "8"
	case mtg.ManaCost9:
		return "9"
	case mtg.ManaCostX:
		return "x"
	default:
		return ""
	}
}
