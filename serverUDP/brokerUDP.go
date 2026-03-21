package main

import (
	"fmt"
	"net"
	"encoding/json"
)

type Coords struct {
	Latitude  float32
	Longitude float32
}

type Mensagem struct {
	Name        string
	Coordinates []Coords
	Doors       string
}


func main(){
	//iniciar o server
	addr, err := net.ResolveUDPAddr("udp", ":4041")
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
		n, _, err1 := conn.ReadFromUDP(buffer)
		if err1 != nil{
			fmt.Println("Erro ao ler dados de cliente: ", err1)
		}

		//copiar os dados para usar no handleconnection
		var msg Mensagem
		err := json.Unmarshal(buffer[:n], &msg)
		if err != nil {
			fmt.Println("Erro no Unmarshal:", err)
			continue
		}
		fmt.Printf("Recebido: %+v\n", msg)
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