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

	//Enviar resposta
	encoder := json.NewEncoder(conn)
	
	err4 := encoder.Encode(resposta)
	if err4 != nil{
		println("Erro: ", err4)
		return
	} 

	//Receber a resposta do Broker sobre a opção escolhida
	var respostaBroker Menu
	err5 := json.NewDecoder(conn).Decode(&respostaBroker)
	if err5 != nil {
		fmt.Println("Erro ao receber resposta do Broker:", err5)
		return
	}

	fmt.Print(respostaBroker)

	//Cliente escolhe sensor para ver
	var chooseSensor int
	_, err6 := fmt.Scan(&chooseSensor)
	if err6 != nil{
		println("Erro: ", err4)
	}

	resposta2 := UserInput{
		Option: chooseSensor,
	}

	encoder.Encode(resposta2)
    
    //Envia a escolha do sensor específico para o servidor
    errEnvia := json.NewEncoder(conn).Encode(resposta2)
    if errEnvia != nil {
        fmt.Println("Erro ao enviar escolha:", errEnvia)
        return
    }

    //Recebe os detalhes do sensor escolhido
    var detalhesSensor Menu
    errFinal := json.NewDecoder(conn).Decode(&detalhesSensor)
    if errFinal != nil {
        fmt.Println("Erro ao receber detalhes:", errFinal)
        return
    }

    // Exibe os detalhes na tela
    fmt.Println(detalhesSensor.Texto)




}


