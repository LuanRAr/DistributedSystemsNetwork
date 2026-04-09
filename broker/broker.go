package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
)

//TCP--------------------------------
type UserInput struct {
	Option int 
}

type Menu struct {
	Texto string `json:"menu"`
}

type Actuator struct {
	sensor Object
}

//UDP-------------------------------
type Coords struct{
	Latitude string
	Longitude string 
}

type Object struct{
	Id string
	Name string
	Coordinates []Coords
	Door string
	Time time.Time
}

type MemoriaSensor struct {
	sync.Mutex
	verDados []Object
}

//globais------------------------------
var currentStatus MemoriaSensor

//SERVER----------------------------------------------------------------------------------
func main(){
	go cleanupSensors()
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
			//envia a lista
			showSensors(encoder)
			
			//envia a pergunta como uma nova mensagem JSON
			watchSensor := Menu{Texto: "Digite o número do sensor para ver detalhes:\n"}
			encoder.Encode(watchSensor)

			//recebe a escolha do sensor específica deste menu
			var input2 UserInput
			if err := decoder.Decode(&input2); err != nil {
				fmt.Println("Erro ao ler escolha do usuário:", err)
				return
			}

			//processa a escolha com Lock para segurança
			currentStatus.Lock()
			index := input2.Option - 1
			
			var respostaDetalhe Menu
			if index >= 0 && index < len(currentStatus.verDados) {
				obj := currentStatus.verDados[index]
				respostaDetalhe.Texto = fmt.Sprintf (
					"📍 [%s] %s | 🌐 (%s, %s) | 🚪 %s\n",
					obj.Id,
					obj.Name,
					obj.Coordinates[0].Latitude,
					obj.Coordinates[0].Longitude,
					obj.Door,
				)

			} else {
				respostaDetalhe.Texto = "\nSensor inválido ou não encontrado."
			}
			currentStatus.Unlock()
			
			//envia o resultado final
			encoder.Encode(respostaDetalhe)

		case 2: 
			//Mostra os sensores ativos
			showSensors(encoder)
			
			//Escolhe qual porta vai fechar
			questionActuator := Menu {
				Texto: "Digite o número do objeto que você deseja acionar o fechamento da porta:",
			}
			encoder.Encode(questionActuator)

			//Resposta do cliente
			var clientResponse UserInput
			decoder.Decode(&clientResponse)

			currentStatus.Lock()
			var statusMsg string = "Objeto não encontrado."

			for i, item := range currentStatus.verDados {
        		if i+1 == clientResponse.Option {
            		statusMsg = sendActuator(item) // Pega o retorno do atuador
            		break
				}
			}

			currentStatus.Unlock()

			//enviar a confirmação final para o cliente
			finalMsg := Menu{ Texto: statusMsg }
			encoder.Encode(finalMsg)

		default:
			break
	}
}

func showSensors(encoder *json.Encoder){
	currentStatus.Lock()
	defer currentStatus.Unlock()

	var resposta Menu

	if len(currentStatus.verDados) == 0 {
		resposta.Texto = "Sem sensores ativos no momento"
	} else {
		texto := "----Sensores ativos----\n"	

		for i, v := range currentStatus.verDados {
			texto += fmt.Sprintf("%d: 📍[%s] %s\n", i+1, v.Id, v.Name)
		}

		resposta.Texto = texto
	}

	encoder.Encode(resposta)
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
		sensor.Time = time.Now()

		//verificação de segurança que aciona atuador pra fechar a porta
		if sensor.Coordinates[0].Latitude < "0" && sensor.Coordinates[0].Longitude < "0" && sensor.Door == "Aberta" {
        
        	//executa em uma goroutine para não travar o recebimento de outros sensores
			go func(obj Object) {
				status := sendActuator(obj)
				fmt.Println("Resultado do Alerta:", status)
			}(sensor)
		}

		find := false
		for i, item := range currentStatus.verDados{
			if item.Id == sensor.Id{
				currentStatus.verDados[i] = sensor
				find = true
				break 
			}
		}

		if !find {
			currentStatus.verDados = append(currentStatus.verDados, sensor)
		}
	
		currentStatus.Unlock()

		handleConnectionUDP(remoteAddr, conn)
	}
	
}

//funcao para limpar a lista com sensores inativos
func cleanupSensors() {
	for {
		time.Sleep(7 * time.Second)

		currentStatus.Lock()
		
		//Nova fatia vazia para armazenar quem ainda está ativo
		novosAtivos := []Object{}

		for _, item := range currentStatus.verDados {
			//verifica se o sensor enviou algo nos últimos 7 segundos
			if time.Since(item.Time) <= 7*time.Second {
				novosAtivos = append(novosAtivos, item)
			} else {
				fmt.Printf(">>> Removendo sensor inativo: [%s] %s\n", item.Id, item.Name)
			}
		}

		//Substitui a lista global pela nova lista (mesmo que vazia)
		currentStatus.verDados = novosAtivos
		
		currentStatus.Unlock()
	}
}

func handleConnectionUDP(remoteAddr *net.UDPAddr, conn *net.UDPConn){

	//mensagem respondendo sensor
	message := "Mensagem recebida!"

	_, err2 := conn.WriteToUDP([]byte(message), remoteAddr)
	if err2 != nil{
		fmt.Println("A mensagem não foi enviada de volta: ", err2)
	}
}

//Atuador
func sendActuator(item Object) string{
	conn, err := net.Dial("tcp", ":8983")
	if err != nil{
		fmt.Println("Erro:", err)
		return "\nErro ao conectar no atuador"
	}

	defer conn.Close()

	//Enviar resposta
	encoder := json.NewEncoder(conn)

	encoder.Encode(item)

	//Ler a resposta do atuador
	var feedback struct {
		Status string `json:"status"`
	}
	decoder := json.NewDecoder(conn)
	err = decoder.Decode(&feedback)
	if err != nil {
		return "\nErro ao receber confirmação do atuador"
	}

	return feedback.Status

}