package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) >= 3 {
		serverAddress := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])

		l, err := net.Listen("tcp", serverAddress)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
		}
		defer l.Close()

		fmt.Println("Listening on " + serverAddress)
		fmt.Println()
		for {
			fmt.Println("esperando una conexion...")
			conn, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("conexion recibida/aceptada")

			fmt.Println("leyendo...")
			content, err := ioutil.ReadAll(conn)
			if err != nil {
				//log.Fatal(err)
			}
			fmt.Println("se a leido!")
			fmt.Println()

			fmt.Println("Message received:", string(content))

		}

	}
}
