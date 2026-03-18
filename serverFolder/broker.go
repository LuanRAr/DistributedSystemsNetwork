package main

import  (
	"fmt"
	"net"
	"os"
	"bufio"
)

func main(){
	ln, error1 := net.Listen("tcp", ":8080")
	if error1 != nil{
		fmt.Println(error1)
		os.Exit(3)
	}

	fmt.Println("Servidor aguardando conexão...")


	for{
		conexao, error2 := ln.Accept()
		if error2 != nil{
			fmt.Println(error2)
			os.Exit(3)
		}

		mensagem, error3 := bufio.NewReader(conexao).ReadString('\n')
		if error3 != nil{
			fmt.Println(error3)
			os.Exit(3)
		}
		fmt.Println(mensagem)

		//resposta
		resposta := "Você conseguiu!\n"
		fmt.Fprintf(conexao, resposta)
	}
}