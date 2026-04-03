package main

import (
	"fmt"
	"net"
)

func actuator(){
	server, err := net.Listen("tcp", ":8983")
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	
	server.Close()

	for{
		conn, err2 := server.Accept()
		if err2 != nil{
			fmt.Println("Erro: ", err2)
			return
		}

		defer conn.Close()

	}
}