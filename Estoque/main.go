package main

import "fmt"

func main() {
	for {
		var op = menu()
		switch op {
		case 1:
			var produto Produto
			var codigo int
			var nome string
			var preco float32
			fmt.Print("Digite o nome do produto: ")
			fmt.Scan(&nome)
			fmt.Print("Digite o código do produto: ")
			fmt.Scan(&codigo)
			fmt.Print("Digite o preco do produto: ")
			fmt.Scan(&preco)
			produto.apresentar(codigo, nome, preco, 10)
			adicionar(produto)
		case 2:
			var codigo, quantidade int
			var preco float32
			fmt.Print("Digite o código do produto: ")
			fmt.Scan(&codigo)
			fmt.Print("Digite a quantidade de produtos: ")
			fmt.Scan(&quantidade)
			fmt.Print("Digite o preço do produto: ")
			fmt.Scan(&preco)
			comprar(codigo, quantidade, preco)
		case 3:
			var codigo, quantidade int
			fmt.Print("Digite o código do produto: ")
			fmt.Scan(&codigo)
			fmt.Print("Digite a quantidade de produtos: ")
			fmt.Scan(&quantidade)
			fmt.Print("Resultado da venda: ", vender(codigo, quantidade))
		case 4:
			listar()
		case 5:
			listarMin()
		case 0:
			break
		}
	}
}
