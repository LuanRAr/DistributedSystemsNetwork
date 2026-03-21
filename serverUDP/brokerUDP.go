package main

import (
	"fmt"
	"net"
)


func main(){
	//iniciar o server
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil{
		fmt.Println("Erro ao iniciar servidor: ", err)
		return
	}

	//se ouvir, fazer a conexão com o server
	conn, err2 := net.ListenUDP("udp", addr)
	if err2 != nil{
		fmt.Println("Erro ao fazer conexão com o server: ", err2)
		return
	}

	defer conn.Close()

	//leitura
	for{
		//mensagem que usuario passou em pacote
		buffer := make([]byte, 1024)

		

		//ler os dados
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil{
			fmt.Println("Erro ao ler dados de cliente: ", err)
		}

		//copiar os dados para usar no handleconnection
		dataCopy := make([]byte, 1024)
		copy(dataCopy, buffer[:n])

		go handleConnection(remoteAddr, conn, dataCopy)
	}

}

func handleConnection( remoteAddr *net.UDPAddr, conn *net.UDPConn, data []byte){
	
	fmt.Printf("%v: %s", remoteAddr, string(data))
	//mensagem respondendo cliente
	message := "Mensagem recebida!"

	_, err2 := conn.WriteToUDP([]byte(message), remoteAddr)
	if err2 != nil{
		fmt.Println("A mensagem não foi enviada de volta: ", err2)
	}
}