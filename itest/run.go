package itest

import "testing"

func RunITest(t *testing.T, test func(t *testing.T)) {
	if !isIntegratedTest() {
		t.Skip("This test should be run by make test")
	}
	setTestEnvVars(t)
	test(t)
}
