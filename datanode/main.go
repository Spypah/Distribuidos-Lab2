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

	grpc "google.golang.org/grpc"
)

func ListenGrpcServer() {
	socket, err := net.Listen("tcp", ":8081")

	failOnError(err, "Failed to connect to open socket")

	grpcServer := grpc.NewServer()

	pb.RegisterGreeterServer(grpcServer, &server{})

	err = grpcServer.Serve(socket)

	failOnError(err, "Failed to connect to server")
}

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) GuardarNombre(ctx context.Context, message *pb.Persona) (*pb.Ok, error) {
	log.Println("Persona recibida: " + fmt.Sprint(message.GetNombre()))

	line := fmt.Sprintf("%s;%s;%s",
		strconv.Itoa(int(message.GetId())),
		message.GetNombre(),
		message.GetApellido(),
	)

	file, err := os.OpenFile("datanode/DATA.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	failOnError(err, "Failed to open file")
	defer file.Close()

	fmt.Fprintf(file, "%s\n", line)

	return &pb.Ok{
		Ok: true,
	}, nil
}

func (s *server) ObtenerNombre(ctx context.Context, message *pb.IdPersona) (*pb.NombrePersona, error) {
	log.Println("Id recibido " + fmt.Sprint(message.GetId()))

	fd, err := os.Open("datanode/DATA.txt")
	failOnError(err, "Failed to open file")
	defer fd.Close()

	reader := bufio.NewReader(fd)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		fields := strings.Split(line, ";")
		id, err := strconv.Atoi(fields[0])

		failOnError(err, "Failed to convert id to int")

		if id == int(message.GetId()) {
			return &pb.NombrePersona{
				Nombre:   fields[1],
				Apellido: strings.ReplaceAll(fields[2], "\n", ""),
			}, nil
		}
	}

	log.Printf("No se encontró el id %d", message.GetId())

	return &pb.NombrePersona{
		Nombre:   "",
		Apellido: "",
	}, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	log.Println("Iniciando DataNode")

	ListenGrpcServer()

	log.Println("Cerrando DataNode")
}
