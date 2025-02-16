package render_test

import (
	"path/filepath"
	"testing"

	"github.com/busser/jumpstart-decklists/pkg/mtg"
	"github.com/busser/jumpstart-decklists/pkg/render"
	"github.com/busser/jumpstart-decklists/pkg/testutils"
	"github.com/stretchr/testify/require"
)

func TestDecklistsAsWebPage(t *testing.T) {
	testutils.Unit(t)
	t.Parallel()

	tests := []struct {
		name  string
		decks []mtg.Decklist
	}{
		{
			name:  "base_case",
			decks: mtg.SampleDecks,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			rendered, err := render.DecklistsAsWebPage(tt.decks)
			require.NoError(t, err)

			pathToGolden := filepath.Join("testdata", t.Name(), "decks.html")
			testutils.CompareGoldenFile(t, pathToGolden, rendered)
		})
	}
}
