package main

import (
	"fmt"
	"net"
	"bufio"
)

func maoin(){
	conectServer, error1 := net.Dial("tcp", "meu-server:8080")
	if error1 != nil{
		fmt.Println(error1)
		return
	}
	
	fmt.Println("Se conectando no server...")
	defer conectServer.Close()
	
	mensagem := "Hello Server!\n"

	fmt.Fprintf(conectServer, mensagem)

	fmt.Println("mensagem Enviada")

	//ouvir resposta
	res, error4 := bufio.NewReader(conectServer).ReadString('\n')
	if error4 != nil{
		fmt.Println(error4)
	}
	fmt.Println(res)
}