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
	Doors       string
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

		go handleActuator(conn) // concorrência
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
	fmt.Printf("🔒 Atuador acionado no objeto [%s] - %s\n", sensor.Id, sensor.Name)
}