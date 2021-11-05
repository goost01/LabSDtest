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

func Jugador() {

	return
}

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {

	// Crear un gRPC canal para comunicarse con el servidor
	// 	-> Esto se crea pasando server address y port number a grpc.Dial()

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
	var id int32 = 0
	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	r, err := client.JoinGame(ctx, &pb.JoinRequest{Id: id, Message: message, Name: name})
	
	id = r.GetId()
	
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Jugador: %s", r.GetMessage())

	comienzaeljuego := false

	for !comienzaeljuego {
		time.Sleep(5 * time.Second)
		message := "Ya comienza el juego??"
		r, err := client.BeginGame(ctx, &pb.BeginRequest{Id: id, Message: message})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		comienzaeljuego = r.GetMessage()
		Jugadoresonline := r.GetOnline()
		fmt.Println(">> Jugadores online: %d", Jugadoresonline)
	}

	// SOLICITUD A LA PRIMERA ETAPA
	stage := "Va a comenzar la etapa??"
	comienzalaetapa := false

	for !comienzalaetapa {

		time.Sleep(10 * time.Second)
		r, err := client.BeginStage(ctx, &pb.BeginStageRequest{Id: id , Stage: stage})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		comienzalaetapa = r.GetStage()
	}

	fmt.Println("Comienza Squid Game")
	//COMIENZA LA PRIMERA ETAPA
	var suma_actual int32 = 0
	dead := false
	var newEtapa bool = true
	// Loop de rondas
	for newEtapa {
		
		var numero int
		fmt.Println("Ingrese su jugada:")
		_, err =  fmt.Scanf("%d",&numero)
		jugada := int32(numero)
		
		r, err := client.SendJugadaE1(ctx, &pb.JugadaE1{Id: id, Jugada: jugada, SumaActual: suma_actual})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		dead = r.GetDead()
		if dead{
			fmt.Println("Has muerto...")
			
		}
		suma_actual = r.GetSumaTotal()
		ronda := r.GetRonda()
		fmt.Println("Suma total %d ronda: %d",suma_actual, ronda)
		if ronda == 4 {
			newEtapa = false
		}
		var newRound bool = false	
		for !newRound{
			fmt.Println("Aun no comienza la ronda")
			time.Sleep(5*time.Second)
			response,err := client.BeginRound(ctx, &pb.PingMsg{Id: id})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			newRound = response.GetRound()
			
		}
		
	}
	var action string
	fmt.Println("Ingrese accion: play = comenzar nueva etapa; check = ver pozo")
	
	comienzalaetapa = false
	for !comienzalaetapa {
		fmt.Scanln(&action)
		if action == "play"{
			r, err := client.BeginStage(ctx, &pb.BeginStageRequest{Id: id , Stage: stage})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			comienzalaetapa = r.GetStage()
			
		}

	}
	fmt.Println("Comienza la ronda 2.")
	newEtapa = false
	for !newEtapa {
		
		var numero int
		fmt.Println("Ingrese su jugada:")
		_, err =  fmt.Scanf("%d",&numero)
		//jugada := int32(numero)
		
		/*r, err := client.SendJugadaE2(ctx, &pb.JugadaE1{Id: id, Jugada: jugada, SumaActual: suma_actual})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		dead = r.GetDead()
		if dead{
			fmt.Println("Has muerto...")
			
		}*/

	// Configurar conexión al server
	//go Jugador(16)
	}
	return
		
}
