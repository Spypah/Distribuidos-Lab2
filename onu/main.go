package main

import (
	"bufio"
	"context"
	pb "distribuidos/proto"
	"log"
	"os"
	"time"

	grpc "google.golang.org/grpc"
)

const (
	omsIP = "dist001.inf.santiago.usm.cl:8080"
	// omsIP = "192.168.10.231:8080"
)

func connectToServer(serverAddress string) pb.GreeterClient {
	// log.println("Conectando a " + serverAddress)
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())

	for err != nil {
		log.Println("Failed to connect to server, retrying in 1 seconds")
		time.Sleep(1 * time.Second)
		conn, err = grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	}

	// log.println("Conexión establecida a" + serverAddress)

	client := pb.NewGreeterClient(conn)

	return client
}

func obtenerLista(serverAddress string, message *pb.Estado) *pb.ListaPersonas {
	client := connectToServer(serverAddress)
	response, err := client.ObtenerLista(context.Background(), message)
	// log.Println("Mensaje enviado a " + serverAddress)

	failOnError(err, "Failed to send message")

	// log.Println("Respuesta desde " + serverAddress + ": " + fmt.Sprint(response.GetPersonas()))

	return response
}

type server struct {
	pb.UnimplementedGreeterServer
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	log.Println("Iniciando ONU")

	for {
		log.Print("Busqueda de personas por estado. Ingrese el estado de la persona (muerto/infectado): ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		//log.Println(input.Text())
		log.Printf("Lista de personas con estado '%s':  \n", input.Text())

		response := obtenerLista(omsIP, &pb.Estado{Estado: input.Text()})

		for _, persona := range response.GetPersonas() {
			log.Println("  Persona: " + persona.GetNombre() + " " + persona.GetApellido())
		}
	}

	log.Println("Terminando ONU")
}
