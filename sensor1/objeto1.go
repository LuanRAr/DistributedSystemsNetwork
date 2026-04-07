package main

import(
	"fmt"
	"math/rand"
	"net"
	"encoding/json"
	"time"
)

type Coords struct{
	Latitude string 
	Longitude string
}

type Object struct{
	Id string
	Name string
	Coordinates []Coords
	Doors string
}

func main(){
	//prepara a porta 4042 para uso em conexão UDP
	addr, err := net.ResolveUDPAddr("udp", ":4042")
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}

	conn, err2 := net.DialUDP("udp", nil, addr)
		if err2 != nil{
			fmt.Println("Erro: ", err2)
			return
		}
		//fechar conexão
		defer conn.Close()


	// loop
	for {

		//mensagem para enviar pro servidor
		message, err3 := json.Marshal(objectData(-40, 40))
		if err3 != nil{
			fmt.Println("Err: ", err3)
			return
		}

		//enviar mensagem
		conn.Write(message)
		fmt.Println("Mensagem enviada")

		//delay de 3 segundos para enviar outro
		time.Sleep(time.Second * 3)
	}
}

func objectData(min float32, max float32) Object{

	//fórmula para numeros aleatórios:
	randomNumber := min + rand.Float32()*(max-min) / 100
	randomNumber2 := min + rand.Float32()*(max-min) / 100

	latNumber := fmt.Sprintf("%.2f", randomNumber)
	longNumber := fmt.Sprintf("%.2f", randomNumber2)

	object1 := Object {
		Id: "26182",
		Name: "Objeto1",
		Coordinates: []Coords{
			{
			Latitude: latNumber,
			Longitude: longNumber,
			},
		},
		Doors: "Fechada",
	}

	return object1
}


