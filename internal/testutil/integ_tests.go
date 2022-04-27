package testutil

import (
	"os"
	"testing"
)

const integTestEnvVar = "INTEG_TESTS"

func IntegTest(t *testing.T) {
	t.Helper()
	if os.Getenv(integTestEnvVar) == "" {
		t.Logf("%s is not set", integTestEnvVar)
		t.SkipNow()
	}
}
