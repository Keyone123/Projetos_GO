package main

import "fmt"

type Produto struct {
	codigo     int
	nome       string
	preco      float32
	quantidade int
	quantMin   int
}

func (p Produto) apresentar(codigo int, nome string, preco float32, quantMin int) {
	p.codigo = codigo
	p.nome = nome
	p.preco = preco
	p.quantMin = quantMin
	p.quantidade = 0
}

func (p Produto) comprar(preco float32, quantidade int) {
	p.preco = (preco * 0.3) + p.preco
	p.quantidade += quantidade
}

func (p Produto) vender(quantidade int) float32 {
	if p.quantidade > quantidade {
		p.quantidade -= quantidade
		return p.preco * float32(quantidade)
	} else if p.quantidade == quantidade {
		p.quantidade = 0
		return p.preco * float32(quantidade)
	} else {
		fmt.Println("Ocorreu um erro na venda do produto")
		return -1
	}
}

func (p Produto) estoqueMin() bool {
	return p.quantidade < p.quantMin
}

func (p Produto) getCodigo() int {
	return p.codigo
}

func (p Produto) getNome() string {
	return p.nome
}

func (p Produto) getPreco() float32 {
	return p.preco
}

func (p Produto) getQuantidade() int {
	return p.quantidade
}

func (p Produto) setPreco(preco float32) {
	p.preco = preco
}
