package main

import (
	"fmt"
)

/*
1. Verificar Número Par ou Ímpar
Escreva uma função chamada ParOuImpar(n int) string que receba um número
inteiro como parâmetro e retorne a string "Par" se o número for par e "Ímpar" caso
contrário.
Depois, no main(), solicite um número ao usuário, chame a função e exiba o
resultado.

2. Soma de Números de 1 a N
Crie uma função SomaAte(n int) int que recebe um número inteiro n e retorna a
soma de todos os números inteiros de 1 até n.
Use um loop for para calcular a soma.

3. Classificação de Notas (Switch)
Crie uma função ClassificarNota(nota int) string que recebe uma nota de 0 a 10 e
retorna uma classificação:
● 9-10: "Excelente"
● 7-8: "Bom"
● 5-6: "Regular"
● 3-4: "Ruim"
● 0-2: "Muito Ruim" Use switch para essa classificação.
*/
func somaAte(n int) int {
	var i int
	num := 0
	for i = 1; i <= n; i++ {
		num += i
	}
	return num
}
func main() {
	var numero int
	fmt.Print("Digite um numero: ")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Println("Seu numero é par")
	} else if numero%2 == 1 {
		fmt.Println("Seu numero é impar")
	}

	fmt.Printf("Numero contado: %d\n", somaAte(numero))

	fmt.Print("Digite sua nota: ")
	var nota int
	fmt.Scan(&nota)

	switch nota {
	case 9, 10:
		fmt.Print("Excelente")
	case 7, 8:
		fmt.Print("Bom")
	case 5, 6:
		fmt.Print("Regular")
	case 3, 4:
		fmt.Print("Ruim")
	case 0, 1, 2:
		fmt.Print("Muito ruim")
	default:
		fmt.Print("Nota irreconhecível")
	}
}
