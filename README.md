# prueba-tcp
Prueba basica del protocolo TCP con Golang.

Se desarrollara un servidor TCP basico el cual podremos conectarnos usando un cliente TCP... por el momento utilizaremos netcat hasta desarrollar el cliente en Go.

## Run Server

`cd server/`
`go build && ./server localhost 8080`

o

`cd server/`
`go run main.go localhost 8080`

## Run Client

`cd client/`
`go build && ./client localhost 8080`

o

`cd client/`
`go run main.go localhost 8080`
