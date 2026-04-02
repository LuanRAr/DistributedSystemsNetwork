package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

//TCP--------------------------------
type UserInput struct {
	Option int 
}

type Menu struct {
	Texto string `json:"menu"`
}

//UDP-------------------------------
type Coords struct{
	Latitude float32  
	Longitude float32 
}

type Object struct{
	Name string
	Coordinates []Coords
	Doors string
}

type MemoriaSensor struct {
	sync.Mutex
	verDados []Object
}

//globais------------------------------
var currentStatus MemoriaSensor

//SERVER-------------------------------------------------------------------------------------------------------------
func main(){
	go serverUDP()
	serverTCP()
	
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

		go handleConnectionTCP(conn)
	}

}

func handleConnectionTCP(conn net.Conn){
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
			showMenu(encoder)
			
		case 2: 
			showMenu(encoder)
		default:
			break
	}

}

func showMenu(encoder *json.Encoder){
	currentStatus.Lock()

	var resposta Menu

	if len(currentStatus.verDados) == 0 {
		resposta.Texto = "Sem sensores ativos no momento"
	} else {
		texto := "----Sensores ativos----\n"	

		for i, v := range currentStatus.verDados {
			texto += fmt.Sprintf("%d: %s\n", i+1, v.Name)
		}

		resposta.Texto = texto
	}

	encoder.Encode(resposta)

	currentStatus.Unlock()

}

//Servidor UDP----------------------------------------------------------------------------
func serverUDP(){
	addr, err := net.ResolveUDPAddr("udp", ":4042")
	if err != nil {
		fmt.Println("Erro", err)
		return
	}

	conn, err2 := net.ListenUDP("udp", addr)
	if err2 != nil{
		fmt.Println("Erro: ", err2)
	}

	defer conn.Close()

	//mensagem que usuario passou em pacote
	buffer := make([]byte, 1024)

	for {
		//ler os dados
		n, remoteAddr, err3 := conn.ReadFromUDP(buffer)
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

		//adicionar objeto à lista
		currentStatus.Lock()

		find := false
		for i, item := range currentStatus.verDados{
			if item.Name == sensor.Name {
				currentStatus.verDados[i] = sensor
				find = true
				break 
			}
		}

		if !find {
			currentStatus.verDados = append(currentStatus.verDados, sensor)
		}
	
		currentStatus.Unlock()

		handleConnectionUDP(remoteAddr, conn, buffer[:n])
	}
	
}

func handleConnectionUDP(remoteAddr *net.UDPAddr, conn *net.UDPConn, data []byte){
	
	//fmt.Printf("%v: %s\n", remoteAddr, string(data))

	//mensagem respondendo sensor
	message := "Mensagem recebida!"

	_, err2 := conn.WriteToUDP([]byte(message), remoteAddr)
	if err2 != nil{
		fmt.Println("A mensagem não foi enviada de volta: ", err2)
	}
}