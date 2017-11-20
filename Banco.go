package main 

import "fmt"

type Banco struct{
	Nome string
	Saldo float32
}
func (b *Banco) Sacar(valor float32) float32{
	b.Saldo = b.Saldo - valor
	return b.Saldo
}
func (b *Banco) Depositar(valor float32) float32{
	b.Saldo = b.Saldo + valor
	return b.Saldo
}
func (b *Banco) TransferePara(valor float32, destino *Banco)(float32, float32){
	b.Saldo = b.Saldo - valor
	destino.Saldo = destino.Saldo + valor
	return b.Saldo, destino.Saldo
}
func main(){
	j := Banco{Nome: "Seu nome", Saldo: 2000}
	a := Banco{Nome: "Seu nome", Saldo: 3000}
	
	fmt.Println("Saldo de "+j.Nome+": ",j.Saldo)
	fmt.Println("Saldo de "+a.Nome+": ",a.Saldo)

	j.TransferePara(4000, &a)
	fmt.Println("---------- Transferencia ----------")
	fmt.Println("Saldo de "+j.Nome+": ",j.Saldo)
	fmt.Println("Saldo de "+a.Nome+": ",a.Saldo)

	//m.Sacar(300)
	//m.Depositar(500)
	//fmt.Println(m.Saldo)
}