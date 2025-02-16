package testutils

import (
	"flag"
	"testing"
)

var (
	runUnitTests        = flag.Bool("unit", true, "run unit tests")
	runIntegrationTests = flag.Bool("integration", false, "run integration tests")
	updateGoldenFiles   = flag.Bool("update", false, "update golden files")
)

func Unit(t *testing.T) {
	t.Helper()

	if !*runUnitTests {
		t.Skip("skipping unit test")
	}
}

func Integration(t *testing.T) {
	t.Helper()

	if !*runIntegrationTests {
		t.Skip("skipping integration test")
	}
}
