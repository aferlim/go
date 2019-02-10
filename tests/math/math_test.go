package math

import "testing"

const defaultErr = "Valor esperado %v, mas o resultado foi %v"

func TestMedia(t *testing.T) {
	t.Parallel()
	expectValue := 7.28

	value := Media(7.2, 9.9, 6.1, 5.9)

	if value != expectValue {
		t.Errorf(defaultErr, expectValue, value)
	}
}
