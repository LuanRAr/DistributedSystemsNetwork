package main

import (
	"fmt"
	"net"
)

type Client struct{
	Name string
	Req int
}

func main(){
	conn, err := net.Dial("tcp", "4041")
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}

	defer conn.Close()
	fmt.Println("Conectando ao server")



}


