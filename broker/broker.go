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
	server, err := net.Listen("tcp", ":4041")
	if err != nil{
		fmt.Println("Erro: ", err)
		return
	}

	for {
		conn, err := server.Accept()
		if err != nil{
			fmt.Println("Erro: ", err)
			return
		}

		defer conn.Close()

		menu := Menu{
			Texto: "1. Listar Objetos\n2. Trancar Objeto",
		}

		//enviar menu pro cliente
		encoder := json.NewEncoder(conn)
		encoder.Encode(menu)

		//receberr input do cliente
		decoder := json.NewDecoder(conn)
		var input UserInput
		err2 := decoder.Decode(&input)
		if err2 != nil {
			fmt.Println("Erro ao ler resposta:", err2)
			return
		}

		fmt.Printf("Usuário escolheu: %d\n", input.Option)

	}


}