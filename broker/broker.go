package main

import (
	"fmt"
	"net"
	"encoding/json"
)

type UserInput struct {
	Option int `json:"option"`
}

type Menu struct {
	Texto string `json:"menu"`
}

func main(){
	serverTCP()

}

//UDP
type Coords struct{
	Latitude float32  
	Longitude float32 
}

type Object struct{
	Name string
	Coordinates []Coords
	Doors string
}

//Servidor TCP----------------------------------------------------------------------------
func serverTCP(){
	server, err := net.Listen("tcp", ":4041")
	if err != nil{
		fmt.Println("Erro: ", err)
		return
	}

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil{
			fmt.Println("Erro: ", err)
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn){
	defer conn.Close()

	fmt.Println("Nova conexão de ", conn.RemoteAddr())

	menu := Menu{
		Texto: "1. Listar Objetos\n2. Trancar Objeto",
	}

	//enviar menu pro cliente
	encoder := json.NewEncoder(conn)
	encoder.Encode(menu)

	//receber input do cliente
	decoder := json.NewDecoder(conn)

	var input UserInput

	err2 := decoder.Decode(&input)
	if err2 != nil {
		fmt.Println("Erro ao ler resposta:", err2)
		return
	}

	switch input.Option {
		case 1:
			fmt.Println("Usuário escolheu 1")
		case 2: 
			fmt.Println("Usuário escolheu 2")
		default:
			fmt.Println("Default")

	}

}

//Servidor UDP----------------------------------------------------------------------------
func serverUDP(){
	addr, err := net.ResolveUDPAddr("udp", ":8081")
	if err != nil {
		fmt.Println("Erro", err)
		return
	}

	conn, err2 := net.ListenUDP("udp", addr)
	if err != nil{
		fmt.Println("Erro: ", err2)
	}

	defer conn.Close()

	for {
		//mensagem que usuario passou em pacote
		buffer := make([]byte, 1024)

		//ler os dados
		n, _, err3 := conn.ReadFromUDP(buffer)
		if err3 != nil{
			fmt.Println("Erro: ", err3)
			continue
		}
		
		var sensor Object
		err4 := json.Unmarshal(buffer[:n], &sensor)
		if err4 != nil {
			fmt.Println("Erro: ", err4)
			continue
		}

	}
}