package main

import (
	"context"
	"fmt"
	//"lab1/game/etapas"
	"log"
	"time"
	//"flag"

	//"lab1/game/etapas"
	pb "lab1/game/helloworld"

	"google.golang.org/grpc"
)

// Comunicación Sincronica -> gRPC

// Comunicación Asincronica -> RabbitMQ

/*
 *	Enviar  la  petición  para  unirse  al  juego  del  calamar.
 *	Ser capaces de jugar cualquiera de las 3 etapas, en otras palabras, tener el código implementado para jugar esas rondas
 	(cada etapa es un subjuego diferente que se explicar´a en las pr´oximas secciones de este documento).
 * 	Enviar  sus  jugadas  en  cada  ronda  de  cada  etapa. En  caso  de  ser  eliminado,  debe  ﬁnalizar  el  proceso.
*/

func Jugador(client pb.GameClient, num_jug int32) {

	return
}

const (
	address     = "localhost:50051"
	defaultName = "world"
)

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
	message := "HOLA deseo unirme a the game"
	// Contact the server and print out its response.
	name := "Jugador 1"
	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	r, err := client.JoinGame(ctx, &pb.JoinRequest{Message: message, Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Jugador: %s", r.GetMessage())

	comienzaeljuego := false

	for !comienzaeljuego {
		time.Sleep(5 * time.Second)
		message := "Ya comienza el juego??"
		r, err := client.BeginGame(ctx, &pb.BeginRequest{Message: message})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		comienzaeljuego = r.GetMessage()
		Jugadoresonline := r.GetOnline()
		fmt.Println(">> Jugadores online: %i", Jugadoresonline)
	}

	// SOLICITUD A LA PRIMERA ETAPA
	stage := "Va a comenzar la etapa??"
	comienzalaetapa := false

	for !comienzalaetapa {

		time.Sleep(5 * time.Second)
		r, err := client.BeginStageE1(ctx, &pb.BeginStageRequest{Stage: stage})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		comienzalaetapa = r.GetStage()
	}

	fmt.Println("Comienza Squid Game")
	//COMIENZA LA PRIMERA ETAPA
	var suma_actual int32 = 0
	dead := false
	// Loop de rondas

	for !dead {
		var numero int
		fmt.Println("Ingrese su jugada:")
		_, err =  fmt.Scanf("%d",&numero)
		jugada := int32(numero)
		suma_actual = int32(suma_actual) + jugada
		r, err := client.SendJugadaE1(ctx, &pb.JugadaE1{Jugada: jugada, SumaActual: suma_actual})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		dead = r.GetDead()
		suma_actual = r.GetSumaTotal()
		ronda := r.GetRonda()
		fmt.Println("Suma total %d ronda: %d",suma_actual, ronda)
		
	}

	fmt.Println("Comienza Squid Game.")
	//go Jugador(16)

	return

}
