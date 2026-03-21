package main

import (
	"fmt"
	"net"
)

func main(){
	//criar resolver
	serverAddr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil{
		fmt.Println("Erro ao definir portas: ", err)
		return
	}
	

	conn, err2 := net.DialUDP("udp", nil, serverAddr)
	if err2 != nil {
		fmt.Println("Erro ao se conectar ao servidor: ", err2)
		return
	}

	defer conn.Close()

	//escrever para o server
	message := "Hello Server!"

	_, err3 := conn.Write([]byte(message))
	if err3 != nil{
		fmt.Println("Mensagem não enviada: ", err3)
		return
	}

	//ler a resposta do server
	buffer := make([]byte, 1024)
	n, remoteAddr, err4 := conn.ReadFromUDP(buffer)
	if err4 != nil {
		fmt.Println("Resposta do server não recebida: ", err4)
		return
	}

	fmt.Printf("%v: %s", remoteAddr, string(buffer[:n]))
}