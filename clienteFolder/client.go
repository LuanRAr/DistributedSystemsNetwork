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

	//Enviar resposta
	encoder := json.NewEncoder(conn)
	//Receber resposta
	decoder := json.NewDecoder(conn)

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

	//Caso apenas resolva mexer no menu
	switch escolha {
		case 1:
			err4 := encoder.Encode(resposta)
			if err4 != nil {
				fmt.Println("Erro: ", err4)
				return
			}

			//receber a lista de sensores 
			var respostaBroker Menu
			err5 := decoder.Decode(&respostaBroker)
			if err5 != nil {
				fmt.Println("Erro ao receber resposta do Broker:", err5)
				return
			}

			fmt.Println(respostaBroker.Texto)

			//receber a pergunta: "Digite o número do sensor para ver detalhes:"
			var perguntaSensor Menu
			errPergunta := decoder.Decode(&perguntaSensor)
			if errPergunta == nil {
				fmt.Print(perguntaSensor.Texto)
			}

			//cliente escolhe o sensor para ver
			var chooseSensor int
			_, err6 := fmt.Scan(&chooseSensor)
			if err6 != nil {
				fmt.Println("Erro no Scan: ", err6)
			}

			resposta2 := UserInput{
				Option: chooseSensor,
			}
			
			//enviar a escolha do sensor específico para o servidor
			errEnvia := encoder.Encode(resposta2)
			if errEnvia != nil {
				fmt.Println("Erro ao enviar escolha:", errEnvia)
				return
			}

			//recebe os detalhes do sensor escolhido
			var detalhesSensor Menu
			errFinal := decoder.Decode(&detalhesSensor)
			if errFinal != nil {
				fmt.Println("Erro ao receber detalhes:", errFinal)
				return
			}

			//exibe os detalhes na tela
			fmt.Println(detalhesSensor.Texto)

		case 2:
			err4 := encoder.Encode(resposta)
			if err4 != nil {
				fmt.Println("Erro ao enviar escolha inicial:", err4)
				return
			}

			//broker envia a lista
			var choose Menu
			decoder.Decode(&choose)
			fmt.Println(choose.Texto)

			// broker envia a pergunta
			var choose2 Menu
			decoder.Decode(&choose2)
			fmt.Println(choose2.Texto)

			//escolhe qual porta abrir
			var chooseDoor int
			_, errchooseDoor := fmt.Scan(&chooseDoor)
			if errchooseDoor != nil {
				fmt.Println("Erro: ", errchooseDoor)
			}

			respostaAtuador := UserInput{
				Option: chooseDoor,
			}

			//resposta final para o broker
			encoder.Encode(respostaAtuador)

			}

	
}
