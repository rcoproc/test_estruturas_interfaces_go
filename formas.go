package main

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.largura + retangulo.altura)
}

func Area(retangulo Retangulo) float64 {
	return retangulo.largura * retangulo.altura
}
