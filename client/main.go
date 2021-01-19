package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) >= 3 {
		serverAddress := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])

		c, err := net.Dial("tcp", serverAddress)
		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = c.Write([]byte("hola mundo, cliente en Go!"))
		if err != nil {
			log.Fatal(err)
			return
		}

	}
}
