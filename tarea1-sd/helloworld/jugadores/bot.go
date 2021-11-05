package main

import (
	"log"
	"context"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func Bot(ctx ,client pb.GameClient, num_jug int) {

	
	client.

	return
}

func main() {

	// Crear un gRPC canal para comunicarse con el servidor
	// 	-> Esto se crea pasando server address y port number a grpc.Dial()

	// Configurar conexión al server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	// Client Stub to perform RPCs
	client := pb.NewGameClient(conn)

	ctx := 

	for i := 1; i < 16; i++ {
		go Bot(client, i)
	}

	return

}
