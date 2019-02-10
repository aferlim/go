package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {

	t.Parallel()
	testes := []struct {
		text     string
		part     string
		expected int
	}{
		{"It`s awesome", "It", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Vitor Andre Lima", "t", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa %v", teste)

		atual := strings.Index(teste.text, teste.part)

		if atual != teste.expected {
			t.Errorf(msgIndex, teste.text, teste.part, teste.expected, atual)
		}
	}
}
