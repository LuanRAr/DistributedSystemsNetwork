package main

import (
	"encoding/json"
	"fmt"
	"net"
)

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
	Doors string
}

func main(){
	server, err := net.Listen("tcp", ":8983")
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	


	for{
		conn, err2 := server.Accept()
		if err2 != nil{
			fmt.Println("Erro: ", err2)
			return
		}

		defer conn.Close()

		decoder := json.NewDecoder(conn)
		
		var sensor Object

		decoder.Decode(&sensor)
		fmt.Println(sensor)


	}
}