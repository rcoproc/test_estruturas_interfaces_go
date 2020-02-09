package main

import "testing"

//func TestArea(t *testing.T) {
//  verificaArea := func(t *testing.T, forma Forma, esperado float64) {
//    t.Helper()
//    resultado := forma.Area()

//    if resultado != esperado {
//      t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
//    }
//  }

//  t.Run("Retângulos", func(t *testing.T) {
//    retangulo := Retangulo{10.0, 6.0}
//    verificaArea(t, retangulo, 60.0)
//  })

//  t.Run("Círculos", func(t *testing.T) {
//    circulo := Circulo{10}
//    verificaArea(t, circulo, 314.1592653589793)
//  })
//}

// Table driven tests - Testes baseados em tabela
func TestArea(t *testing.T) {
	testesArea := []struct {
		nome    string
		forma   Forma
		temArea float64
	}{
		{nome: "Retângulo", forma: Retangulo{Largura: 12, Altura: 6}, temArea: 72.0},
		{nome: "Circulo", forma: Circulo{Raio: 10}, temArea: 314.1592653589793},
		{nome: "Triângulo", forma: Triangulo{Base: 12, Altura: 6}, temArea: 36.0},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.temArea {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.temArea)
			}
		})
	}
}
