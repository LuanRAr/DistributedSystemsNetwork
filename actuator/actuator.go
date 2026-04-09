package main

import (
	"encoding/json"
	"fmt"
	"net"
)

//-----------------------------------
type Coords struct {
	Latitude  string
	Longitude string
}

type Object struct {
	Id          string
	Name        string
	Coordinates []Coords
	Door        string
}

//-----------------------------------
func main() {
	server, err := net.Listen("tcp", ":8983")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	defer server.Close()

	fmt.Println("[ATUADOR] Aguardando conexões na porta 8983...")

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		go handleActuator(conn) //concorrência
	}
}

func handleActuator(conn net.Conn) {
	defer conn.Close()

	fmt.Println("[ATUADOR] Nova conexão de:", conn.RemoteAddr())

	decoder := json.NewDecoder(conn)

	var sensor Object

	err := decoder.Decode(&sensor)
	if err != nil {
		fmt.Println("Erro ao decodificar:", err)
		return
	}

	// Simula ação do atuador
	msg := fmt.Sprintf("🔒 Atuador acionado trancando porta do objeto [%s] - %s", sensor.Id, sensor.Name)
	fmt.Println(msg)

	//Envia confirmação de volta para o Broker
	feedback := struct {
		Status string `json:"status"`
	}{
		Status: msg,
	}
	encoder := json.NewEncoder(conn)
	encoder.Encode(feedback)
}