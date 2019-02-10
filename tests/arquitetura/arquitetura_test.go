package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {

	t.Parallel()

	if runtime.GOARCH == "amd64" {
		t.Skip("Naão funciona na arquitetura amd64")
	}

	t.Fail()
}
