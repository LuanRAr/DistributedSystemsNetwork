package main

import(
	"fmt"
	"math/rand"
	"net"
	"encoding/json"
	"time"
)

type Coords struct{
	Latitude float32  
	Longitude float32 
}

type StatusDoor struct{
	StatusDoor []string
}

type Object struct{
	Name string
	Coordinates []Coords
	Doors string
}

func main(){
	//prepara a porta 4041 para uso em conexão UDP
	addr, err := net.ResolveUDPAddr("udp", ":4041")
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}

	// loop
	for {
		conn, err2 := net.DialUDP("udp", nil, addr)
		if err2 != nil{
			fmt.Println("Erro: ", err2)
			return
		}

		//fechar conexão
		defer conn.Close()

		//mensagem para enviar pro servidor
		message, err3 := json.Marshal(objectData(-40, 40))
		if err3 != nil{
			fmt.Println("Err: ", err3)
			return
		}

		//enviar mensagem
		conn.Write(message)
		fmt.Println("Mensagem enviada")

		//delay de 5 segundos para enviar outro
		time.Sleep(time.Second * 3)
		
	}


}

func objectData(min float32, max float32) Object{

	//fórmula aplicada:
	randomNumber := min + rand.Float32()*(max-min)
	randomNumber2 := min + rand.Float32()*(max-min)

	//escolher porta aberta ou fechada
	status := StatusDoor {
		StatusDoor: []string{"ABERTA", "FECHADA"},
	}
	//sorteio do status da porta
	door:= (status.StatusDoor[rand.Intn(2)])


	object1 := Object {
		Name: "Objeto1",
		Coordinates: []Coords{
			{
			Latitude: randomNumber,
			Longitude: randomNumber2,
			},
		},
		Doors: door,
	}

	return object1
}


