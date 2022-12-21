package itest

import (
	"os"
	"testing"
)

func setTestEnvVars(t *testing.T) {
	t.Setenv("REC_POSTGRES_SERVICE_HOST", "localhost")
	t.Setenv("REC_POSTGRES_SERVICE_PORT", "5432")
	t.Setenv("REC_POSTGRES_SERVICE_USER", "postgres")
	t.Setenv("REC_POSTGRES_SERVICE_PASSWORD", "")
	t.Setenv("REC_POSTGRES_SERVICE_DATABASE", "local")
}

func isIntegratedTest() bool {
	return os.Getenv("IT_TEST") == "TRUE"
}
