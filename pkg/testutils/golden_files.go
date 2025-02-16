package testutils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

// CompareGoldenFile compares the given data against the golden file at the
// given path. If the contents differ, the test will fail and continue. If the
// golden file doesn't exist, the test will fail. If the update flag is set,
// the golden file will be created or updated with the given data.
func CompareGoldenFile(t *testing.T, path string, data []byte) {
	t.Helper()

	if *updateGoldenFiles {
		t.Logf("updating golden file: %s", path)
		writeGoldenFile(t, path, data)
	}

	expected := readGoldenFile(t, path)

	if diff := cmp.Diff(expected, data); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func readGoldenFile(t *testing.T, path string) []byte {
	t.Helper()

	data, err := os.ReadFile(path)
	require.NoError(t, err)

	return data
}

func writeGoldenFile(t *testing.T, path string, data []byte) {
	t.Helper()

	const dirPermissions = 0o755
	const filePermissions = 0o644

	err := os.MkdirAll(filepath.Dir(path), dirPermissions)
	require.NoError(t, err)

	err = os.WriteFile(path, data, filePermissions)
	require.NoError(t, err)
}
