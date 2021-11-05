package main

import (
	"context"
	//"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	//"lab1/game/etapas"
	pb "lab1/game/helloworld"

	"google.golang.org/grpc"
)
var jugadoresOnline []int32
var contador int32 = 0
var count int32 = 1
var etapa = 1
var jugadores_etapa []int32
var jugadores_ronda []int32
var jugadores_turno []int32
var begin bool = false
var shuffled bool = false
var maxplayers int32 = 2
var players_alive int32 = 0
var players_round int32= 0
var optLiderE1 int32
var beginR bool = true
var contador_muertos int32 = 0
var equipos [][]int32


// Variables Globales etapa 1

var roundE1 int32 = 0 

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
//Utility
func contains(s []int32, elem int32) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}
//Int32 to String
func ItoS(n int32) string {
    buf := [11]byte{}
    pos := len(buf)
    i := int64(n)
    signed := i < 0
    if signed {
        i = -i
    }
    for {
        pos--
        buf[pos], i = '0'+byte(i%10), i/10
        if i == 0 {
            if signed {
                pos--
                buf[pos] = '-'
            }
            return string(buf[pos:])
        }
    }
}

func (s *server) JoinGame(ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	log.Printf("Received: %v", in.GetMessage())	

	var id int32
	var message string

	if(	in.GetId() == 0 ){ // Jugador no seteado
		id = count
		count++
		players_alive = players_alive + 1 
		jugadoresOnline = append(jugadoresOnline, id)
		message =  "Te has inscrito al juego del calamar, Bienvenido Jugador" + ItoS(id)
	}else {
		id = in.GetId()
		message = "Esperando al resto de jugadores..."
	}

	log.Printf("Jugador añadido ")
	fmt.Printf("%v", jugadoresOnline)
	return &pb.JoinReply{Id: id, Message: message}, nil
}
// BeginGame implements game.GameServer
func (s *server) BeginGame(ctx context.Context, in *pb.BeginRequest) (*pb.BeginReply, error) {
	log.Printf("Recibido de " +ItoS(in.GetId())+": %v", in.GetMessage())
	begin = false
	if players_alive == maxplayers {
		begin = true
		log.Printf("Comienza el juego")
	}	
	return &pb.BeginReply{Message: begin, Online: int32(players_alive)}, nil
}

//BeginStage
func (s *server) BeginStage(ctx context.Context, in *pb.BeginStageRequest) (*pb.BeginStageReply, error) {
	
	begin = false
	id_jugador := in.GetId()
	
	log.Printf("Received from "+ ItoS(id_jugador) +": %v", in.GetStage())
	contiene := contains(jugadores_etapa, id_jugador)
	
	if !contiene {
		jugadores_etapa = append(jugadores_etapa, id_jugador)
	}
	fmt.Println(players_alive, jugadores_etapa)
	if int32(len(jugadores_etapa)) == players_alive {
		begin = true
		
	}
	if etapa == 2{
		if !shuffled{
			//shuffle
			for i := range jugadores_etapa{
				j := rand.Intn(i + 1)
				jugadores_etapa[i], jugadores_etapa[j] = jugadores_etapa[j], jugadores_etapa[i]
			}
			shuffled = true
			for i := 0; i < len(jugadores_etapa)/2; i++{
				equipos[i] = jugadores_etapa[2*i: 2*i+1]
			}
			fmt.Println(equipos)
		}
		if len(jugadores_etapa) % 2 == 1 && id_jugador == jugadores_etapa[len(jugadores_etapa)-1]{
			return &pb.BeginStageReply{Stage: begin, Dead: true}, nil
		}
		
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
	
	var curr_sum int32 = 0

	if beginR {
		optLiderE1 = int32(rand.Intn(4) + 6)
		beginR = false
		jugadores_etapa = nil
		jugadores_ronda = nil
		jugadores_turno = nil
		roundE1++
	}
	curr_sum = in.GetJugada() + in.GetSumaActual()

	if  in.GetJugada() >= optLiderE1 {
		dead = true
	}else {
		dead = false
	}

	if roundE1 == 4 && curr_sum < 21 {
		dead = true
	}else {
		dead = false
	}
	
	id_jugador := in.GetId()
	contiene := contains(jugadores_ronda, id_jugador)
	
	if !contiene {
		jugadores_ronda = append(jugadores_ronda, id_jugador)
	}
	 
	if int32(len(jugadores_ronda)) == players_alive {
		beginR = true
		players_alive = players_alive - contador_muertos
		contador_muertos = 0
		etapa = 2
		fmt.Println("Comienza la siguiente ronda")
	}

	if dead {
		contador_muertos++
		fmt.Println("Ha muerto el jugador " + ItoS(in.GetId()))
	}

	return &pb.PlayerStatusE1{SumaTotal: curr_sum, Dead: dead, Ronda: roundE1 }, nil
	
}
//func (s *server) SendJugadaE2(ctx context.Context, in *pb.JugadaE2)  (*pb.PingMsg, error){}
func (s *server) BeginRound(ctx context.Context, in *pb.PingMsg) (*pb.BeginRoundReply, error) {
	begin = false
	id_jugador := in.GetId()
	log.Printf("Received from : "+ ItoS(id_jugador))
	contiene := contains(jugadores_turno, id_jugador)
	
	if !contiene {
		jugadores_turno = append(jugadores_turno, id_jugador)
	}

	if int32(len(jugadores_turno)) == players_alive {
		begin = true
		
	}
	return &pb.BeginRoundReply{Round: begin}, nil
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
