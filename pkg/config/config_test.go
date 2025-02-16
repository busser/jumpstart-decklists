package config_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/busser/jumpstart-decklists/pkg/config"
	"github.com/busser/jumpstart-decklists/pkg/testutils"
	"github.com/stretchr/testify/require"
)

func TestReadDecklists(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "base_case",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pathToInput := filepath.Join("testdata", tt.name, "input.yml")

			input, err := os.Open(pathToInput)
			require.NoError(t, err)
			defer input.Close()

			decklists, err := config.ReadDecklists(input)
			require.NoError(t, err)

			encoded, err := json.MarshalIndent(decklists, "", "  ")
			require.NoError(t, err)

			pathToGolden := filepath.Join("testdata", tt.name, "output.json")
			testutils.CompareGoldenFile(t, pathToGolden, encoded)
		})
	}
}
