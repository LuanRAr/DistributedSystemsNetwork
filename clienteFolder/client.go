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

	//caso apenas resolva mexer no menu
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


			if respostaBroker.Texto != "Sem sensores ativos no momento" {
			//recebe a pergunta
			var perguntaSensor Menu
			decoder.Decode(&perguntaSensor)
			fmt.Print(perguntaSensor.Texto)

			//escolhe o sensor
			var chooseSensor int
			fmt.Scan(&chooseSensor)
			encoder.Encode(UserInput{Option: chooseSensor})

			//recebe detalhes
			var detalhesSensor Menu
			decoder.Decode(&detalhesSensor)
			fmt.Println(detalhesSensor.Texto)
		}

		
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

			//só continua se NÃO for a mensagem de "vazio"
			if choose.Texto != "Sem sensores ativos no momento" {
				//recebe a pergunta
				var choose2 Menu
				decoder.Decode(&choose2)
				fmt.Println(choose2.Texto)

				//escolhe qual trancar
				var chooseDoor int
				fmt.Scan(&chooseDoor)
				encoder.Encode(UserInput{Option: chooseDoor})

				//recebe confirmação
				var confirmacao Menu
				if err := decoder.Decode(&confirmacao); err == nil {
					fmt.Println("\n>>> " + confirmacao.Texto)
				}
			}
		}
}