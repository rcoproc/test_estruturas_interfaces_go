package main

import "testing"

func TestArea(t *testing.T) {
	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()
		resultado := forma.Area()

		if resultado != esperado {
			t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
		}
	}

	t.Run("Retângulos", func(t *testing.T) {
		retangulo := Retangulo{10.0, 6.0}
		verificaArea(t, retangulo, 60.0)
	})

	t.Run("Círculos", func(t *testing.T) {
		circulo := Circulo{10}
		verificaArea(t, circulo, 314.1592653589793)
	})
}
