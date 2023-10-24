package main

import (
	"bufio"
	"context"
	pb "distribuidos/proto"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	grpc "google.golang.org/grpc"
)

var lastId int64 = 0

const (
	//"dist002.inf.santiago.usm.cl:8081"
	dataNode1IP = "dist002.inf.santiago.usm.cl:8081"
	//"dist003.inf.santiago.usm.cl:8081"
	dataNode2IP = "dist003.inf.santiago.usm.cl:8081"
)

func ListenGrpcServer() {
	socket, err := net.Listen("tcp", ":8080")

	failOnError(err, "Failed to connect to open socket")

	grpcServer := grpc.NewServer()

	pb.RegisterGreeterServer(grpcServer, &server{})

	err = grpcServer.Serve(socket)

	failOnError(err, "Failed to connect to server")
}

func connectToServer(serverAddress string) pb.GreeterClient {
	log.Println("Conectando a " + serverAddress)

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())

	for err != nil {
		log.Println("Failed to connect to server, retrying in 1 seconds")
		time.Sleep(1 * time.Second)
		conn, err = grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	}

	log.Println("Conexión establecida a" + serverAddress)

	client := pb.NewGreeterClient(conn)

	return client
}

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) NotificarPersona(ctx context.Context, message *pb.EstadoPersona) (*pb.Ok, error) {
	log.Println("Persona recibida: " + fmt.Sprint(message.GetNombre()))

	lastId += 1

	id := lastId

	persona := &pb.Persona{
		Id:       id,
		Nombre:   message.GetNombre(),
		Apellido: message.GetApellido(),
	}

	lastnameByte := byte(persona.Apellido[0])

	dataNodeId := 0

	if (lastnameByte < 78 && lastnameByte > 64) || (lastnameByte < 110 && lastnameByte > 96) {
		dataNodeId = 1

		guardarNombre(dataNode1IP, persona)
	} else {
		dataNodeId = 2

		guardarNombre(dataNode2IP, persona)
	}

	line := fmt.Sprintf("%s;%s;%s",
		message.GetEstado(),
		strconv.Itoa(int(id)),
		strconv.Itoa(int(dataNodeId)),
	)

	file, err := os.OpenFile("oms/DATA.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	failOnError(err, "Failed to open file")
	defer file.Close()

	fmt.Fprintf(file, "%s\n", line)

	return &pb.Ok{
		Ok: true,
	}, nil
}

func (s *server) ObtenerLista(ctx context.Context, message *pb.Estado) (*pb.ListaPersonas, error) {
	log.Println("Query recibida" + fmt.Sprint(message.GetEstado()))

	fd, err := os.Open("oms/DATA.txt")
	failOnError(err, "Failed to open file")
	defer fd.Close()

	reader := bufio.NewReader(fd)

	peopleList := make([]*pb.NombrePersona, 0)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		fields := strings.Split(line, ";")
		estado := fields[0]

		failOnError(err, "Failed to convert id to int")

		if estado == message.GetEstado() {
			id, err := strconv.Atoi(fields[1])
			failOnError(err, "Failed to convert id to int")

			datanode, err := strconv.Atoi(strings.ReplaceAll(fields[2], "\n", ""))
			failOnError(err, "Failed to convert id to int")

			ip := ""

			if datanode == 1 {
				ip = dataNode1IP
			} else {
				ip = dataNode2IP
			}

			persona := obtenerNombre(ip, &pb.IdPersona{
				Id: int64(id),
			})

			peopleList = append(peopleList, persona)
		}
	}

	return &pb.ListaPersonas{
		Personas: peopleList,
	}, nil
}

func guardarNombre(serverAddress string, message *pb.Persona) bool {
	client := connectToServer(serverAddress)
	response, err := client.GuardarNombre(context.Background(), message)
	log.Println("Mensaje enviado a " + serverAddress)

	failOnError(err, "Failed to send message")

	log.Println("Respuesta desde " + serverAddress + ": " + fmt.Sprint(response.GetOk()))

	return response.GetOk()
}

func obtenerNombre(serverAddress string, message *pb.IdPersona) *pb.NombrePersona {
	client := connectToServer(serverAddress)
	response, err := client.ObtenerNombre(context.Background(), message)
	log.Println("Mensaje enviado a " + serverAddress)

	failOnError(err, "Failed to send message")

	log.Println("Respuesta desde " + serverAddress + ": " + fmt.Sprint(response.GetNombre()))

	return response
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	log.Println("Iniciando oms")

	ListenGrpcServer()

	log.Print("Cerrando oms")
}
