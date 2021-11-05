package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	//"lab1/game/etapas"
	pb "lab1/game/helloworld"

	"google.golang.org/grpc"
)
var jugadoresOnline []string
var contador = 0

var begin bool
var maxplayers = 3
var players_alive = 0

var ronda int

// Variables Globales etapa 1

var roundE1 int 

/*
- Operaciones :
	-> Jugadores
		* Enviar señal de inicio de etapa
		* Actualizar sobre el estado en el juego
	-> Pozo
		* Enviar jugadores eliminados
		* Pedir monto total acumulado
	-> NameNode
		* Enviar jugadas de los jugadores

- Recepción:
	->	Jugadores
		* Participar Juego
		* Enviar jugadas
	-> Pozo
		* Devolver monto total del pozo
	-> NameNode
		*Devolver registro del jugador
*/

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGameServer
}

// Etapas
/*func Etapa1() {

	// Mensaje comienza etapa 1 a jugadores
	total_jug_vivos := 0
	total_jug_resp := 0

	for total_jug_resp < total_jug_vivos {
		

	}

	for rondas < 3 {

		total_jug_resp = 0

		// Recibir jugadas	
		for total_jug_resp < total_jug_vivos {

		}

		// Respuesta del lider
		numLider := etapas.RandomNumberLider()

		for j := 0;  j < total_jug_vivos; j++ {

		}

		if numLider >    {


		}	


		
	}

	return 1
}


func Etapa2() {

	// Dividir jugadores en equipo

	for round :=  ; round <   


	numLider :=	etapa.RandomNumberStageTwo()


}*/

// BeginGame implements game.GameServer
func (s *server) BeginGame(ctx context.Context, in *pb.BeginRequest) (*pb.BeginReply, error) {
	log.Printf("Recibido: %v", in.GetMessage())
	begin = false
	if players_alive == maxplayers {
		begin = true
	}	
	return &pb.BeginReply{Message: begin, Online: int32(players_alive)}, nil
}

func (s *server) JoinGame(ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	log.Printf("Received: %v", in.GetMessage())
	jugadoresOnline = append(jugadoresOnline, in.GetName())
	players_alive = players_alive + 1
	log.Printf("Jugador añadido ")
	fmt.Printf("%v", jugadoresOnline)
	return &pb.JoinReply{Message: "Te has inscrito al juego del calamar " + in.GetMessage()}, nil
}
//BeginStage
func (s *server) BeginStageE1(ctx context.Context, in *pb.BeginStageRequest) (*pb.BeginStageReply, error) {
	log.Printf("Received: %v", in.GetStage())
	begin = false
	contador = contador +1
	if contador == players_alive {
		begin = true
		ronda = 1
	}
	return &pb.BeginStageReply{Stage: begin}, nil
}

// EndGame implements game.GameServer
func (s *server) EndGame(ctx context.Context, in *pb.EndRequest) (*pb.EndReply, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.EndReply{Message: "Hello " + in.GetMessage()}, nil
}

/* func (s *server) BeginStage(ctx context.Context, in *pb.BeginStageRequest) (*pb.BeginStageReply, error) {
	log.Printf("Received: %v", in.GetStage())
	return &pb.BeginStageReply{Stage: "1" + in.GetStage()}, nil
 } */
 
func (s *server) SendJugadaE1(ctx context.Context, in *pb.JugadaE1)  (*pb.PlayerStatusE1, error){

	var dead bool
	var optLiderE1 = rand.RandIntn(5) +5
	if  in.GetJugada() >= int32(optLiderE1) {
		dead = true
	}else {
		dead = false
	}

	if roundE1 == 4 && in.GetSumaActual() > 21 {
		dead = false
	}else {
		dead = false
	}
	
	return &pb.PlayerStatusE1{SumaTotal: in.GetJugada()+in.GetSumaActual(), Dead: dead, Ronda: int32(1) }, nil
	
}

const (
	port = ":50051"
)

func main() {

	end := false

	// Puerto para escucha de los clientes
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	// Crea una isntancia de server gRPC
	s := grpc.NewServer()

	// Registrar nuestra implementación de server con gRPC server
	pb.RegisterGameServer(s, &server{})

	log.Printf("Server escuchando en %v", lis.Addr())

	// Llamar Server() con los detalles de puerto para realizar un bloquero
	// de espera hasta que el proceso sea killed o Stop() es llamado
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	for !end {

		// Menu de inicio
		fmt.Println("Bienvenido a Squid Game.")
		fmt.Println("Selecciona una opción:")
		fmt.Println(" 1. Iniciar Juego")
		fmt.Println(" 2. Salir del Juego")

		var uselect string

		fmt.Scanln(&uselect)

		if uselect == "1" { // Inicia Juego

			for !begin {			
				fmt.Println("Esperando Jugadores... [%d]", players_alive)
			}

			fmt.Println("Comienzo de la Etapa 1...")
			time.Sleep(2 * time.Second)
			//etapaUno()
			fmt.Println("Comienzo de la Etapa 2...")
			time.Sleep(2 * time.Second)
			//etapaDos()
			fmt.Println("Comienzo de la Etapa 3...")
			time.Sleep(2 * time.Second)
			//etapaTres()

		} else if uselect == "2" { // Salir del Juego
			end = true
			s.Stop()
		} else {

			fmt.Println("Ingresaste una opción no válida, intentalo denuevo.")

		}

	}

	return

}
