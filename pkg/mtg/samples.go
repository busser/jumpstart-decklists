package mtg

var SampleDecks = []Decklist{
	SampleDecklistStompStomp,
	SampleDecklistTasteTheRainbow,
}

var SampleDecklistStompStomp = Decklist{
	Name: "Stomp stomp",
	Colors: []Color{
		ColorGreen,
	},
	Art: []byte("not really an image to save space"),
	Cards: []DecklistItem{
		{
			Count: 1,
			Card: Card{
				Name: "Llanowar Elves",
				ManaCost: []ManaCostItem{
					ManaCostGreen,
				},
			},
		},
		{
			Count: 1,
			Card: Card{
				Name: "Grizzly Bears",
				ManaCost: []ManaCostItem{
					ManaCost1,
					ManaCostGreen,
				},
			},
		},
		{
			Count: 1,
			Card: Card{
				Name: "Goldvein Hydra",
				ManaCost: []ManaCostItem{
					ManaCostX,
					ManaCostGreen,
				},
			},
		},
		{
			Count: 17,
			Card: Card{
				Name: "Forest",
			},
		},
	},
}

var SampleDecklistTasteTheRainbow = Decklist{
	Name: "Taste the rainbow",
	Art:  []byte("not really an image to save space"),
	Colors: []Color{
		ColorWhite,
		ColorBlue,
		ColorBlack,
		ColorRed,
		ColorGreen,
	},
	Cards: []DecklistItem{
		{
			Count: 5,
			Card: Card{
				Name: "Progenitus",
				ManaCost: []ManaCostItem{
					ManaCostWhite,
					ManaCostWhite,
					ManaCostBlue,
					ManaCostBlue,
					ManaCostBlack,
					ManaCostBlack,
					ManaCostRed,
					ManaCostRed,
					ManaCostGreen,
					ManaCostGreen,
				},
			},
		},
		{
			Count: 3,
			Card: Card{
				Name: "Plains",
			},
		},
		{
			Count: 3,
			Card: Card{
				Name: "Island",
			},
		},
		{
			Count: 3,
			Card: Card{
				Name: "Swamp",
			},
		},
		{
			Count: 3,
			Card: Card{
				Name: "Mountain",
			},
		},
		{
			Count: 3,
			Card: Card{
				Name: "Forest",
			},
		},
	},
}
