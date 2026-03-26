package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Menu struct {
	Texto string `json:"menu"`
}

type UserInput struct {
	Option int `json:"option"`
}

func main(){
	conn, err := net.Dial("tcp", ":4041")
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}

	defer conn.Close()
	fmt.Println("Conectando ao server")

	var msg Menu

	err2:= json.NewDecoder(conn).Decode(&msg)
	if err2 != nil{
		fmt.Println("Erro: ", err2)
	}

	fmt.Println(msg.Texto)

	//resposta do cliente
	var escolha int
	_, err3 := fmt.Scan(&escolha)
	if err3 != nil{
		println("Erro: ", err3)
	}

	resposta := UserInput{
		Option: escolha,
	}
	
	err4 := json.NewEncoder(conn).Encode(resposta)
	if err4 != nil{
		println("Erro: ", err4)
		return
	} 

}


