package main

import "fmt"

func main() {
	var v [6]int
	v[0] = 1
	v[1] = 2
	v[2] = 4
	v[3] = 7
	v[4] = 7
	v[5] = 5

	fmt.Print(v[0])
	fmt.Print(v[1])
	fmt.Print(v[2])
	fmt.Print(v[3])
	fmt.Print(v[4])
	fmt.Print(v[5])

	fmt.Println()

	for i := 0; i < len(v); i++ {
		fmt.Print(v[i])
	}

	fmt.Println()

	for i, val := range v {
		fmt.Printf("Indice: %d,Valor de %d\n", i, val)
	}
	var matriz [3][2]int // Matriz de 3 linhas e 2 colunas
	var nomes [3]string
	var indice int
	var ra int
	var nome string
	fmt.Println()
	for i := 0; i < 3; i++ {
		fmt.Printf("Digite o indice do aluno %d: ", i+1)
		fmt.Scan(&indice)
		matriz[i][0] = indice
		fmt.Printf("Digite o RA do aluno %d: ", i+1)
		fmt.Scan(&ra)
		matriz[i][1] = ra
		fmt.Printf("Digite o nome do aluno %d: ", i+1)
		fmt.Scan(&nome)
		nomes[i] = nome
	}
	fmt.Println()
	fmt.Println("Indice RA Nome")
	for i := 0; i < len(matriz); i++ {
		fmt.Printf("%d %d %s\n", matriz[i][0], matriz[i][1], nomes[i])
	}

	type Produto struct {
		Nome       string
		Preco      float64
		Quantidade int
	}
	p := Produto{
		Nome:       "Maça",
		Preco:      10.00,
		Quantidade: 10,
	}

	fmt.Printf("O Nome do produto e %s\nO valor: %.2f\nA quantidade é %d", p.Nome, p.Preco, p.Quantidade)

}
