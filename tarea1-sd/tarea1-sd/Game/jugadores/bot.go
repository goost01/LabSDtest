package main

import (
	"context"
	"log"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func Bot(client pb.GameClient, num_jug int) {

	ctx := context.Background()

	client.BeginGame(ctx)

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

	for i := 1; i < 16; i++ {
		go Bot(client, i)
	}

	return

}
