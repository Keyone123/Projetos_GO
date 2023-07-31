package main

import "fmt"

var produtos []Produto

func menu() int {
	var numero int
	fmt.Println("\n---------------------------- Menu Inicial ----------------------------")
	fmt.Println("1.Adicionar Produto")
	fmt.Println("2.Comprar Produto")
	fmt.Println("3.Vender Produto")
	fmt.Println("4.Listar Produtos")
	fmt.Println("5.Listar Produtos Abaixo do Mínimo")
	fmt.Print("0.Sair")
	fmt.Print("\n------------------------------------------------------------------------")
	fmt.Print("\nDigite a opção que deseja selecionar: ")
	fmt.Scan(&numero)
	return numero
}

func adicionar(p Produto) {
	if p.getCodigo() <= 0 || len(p.getNome()) == 0 || p.getPreco() <= 0.000 {
		fmt.Println("Os dados do produto estão incorretos")
		return
	}
	produtos = append(produtos, p)
}

func comprar(cod int, quant int, preco float32) {
	var i int
	if cod <= 0 || quant <= 0 {
		fmt.Println("Os dados do produto estão incorretos")
		return
	}
	for i = 0; i < len(produtos); i++ {
		if cod == produtos[i].getCodigo() {
			produtos[i].comprar(preco, cod)
			fmt.Print("\nProduto comprado com sucesso")
			return
		}
	}
	fmt.Print("Produto não foi encontrado")
}

func vender(cod int, quant int) float32 {
	var i int
	if cod <= 0 || quant <= 0 {
		fmt.Println("Os dados do produto estão incorretos")
		return -1
	}
	for i = 0; i < len(produtos); i++ {
		if cod == produtos[i].getCodigo() {
			fmt.Println("\nProduto vendido com sucesso")
			return produtos[i].vender(quant)
		}
	}
	fmt.Println("\nProduto não foi encontrado")
	return -1
}

func listar() {
	var i int
	for i = 0; i < len(produtos); i++ {
		fmt.Println("Produto ", i+1, ": ", produtos[i].getNome())
	}
}

func listarMin() {
	var i int
	fmt.Println("\nProdutos com estoque abaixo do minimo: ")
	for i = 0; i < len(produtos); i++ {
		if produtos[i].estoqueMin() {
			fmt.Println(produtos[i].getNome())
		}
	}
}
