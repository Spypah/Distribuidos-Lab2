package main

import (
	"bufio"
	"bytes"
	"context"
	pb "distribuidos/proto"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	grpc "google.golang.org/grpc"
)

const (
	omsIP = "dist001.inf.santiago.usm.cl:8080"
	//omsIP = "192.168.10.231:8080"
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

func notificarPersona(serverAddress string, message *pb.EstadoPersona) bool {
	client := connectToServer(serverAddress)
	response, err := client.NotificarPersona(context.Background(), message)
	// log.Println("Mensaje enviado a " + serverAddress)

	failOnError(err, "Failed to send message")

	// log.Println("Respuesta desde " + serverAddress + ": " + fmt.Sprint(response.GetOk()))

	return response.GetOk()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func lineCounter(r bufio.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func getPerson() (string, string) {
	file, err := os.Open("continente/names.txt")
	failOnError(err, "Failed to open file")

	defer file.Close()

	reader := bufio.NewReader(file)
	lines, err := lineCounter(*reader)
	lines += 1
	failOnError(err, "Failed to count lines")

	rand.Seed(time.Now().UnixNano())
	randomPerson := rand.Intn(lines)

	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	for i := 0; i <= randomPerson; i++ {
		scanner.Scan()
	}
	personName := scanner.Text()

	content, err := os.ReadFile("continente/names.txt")
	failOnError(err, "Failed to read file")

	linesContent := strings.Split(string(content), "\n")
	linesContent = append(linesContent[:randomPerson], linesContent[randomPerson+1:]...)
	output := strings.Join(linesContent, "\n")
	err = os.WriteFile("continente/names.txt", []byte(output), 0644)
	failOnError(err, "Failed to write to file")
	fmt.Printf("Lineas contadas %d, Eliminando a %s, de la lista. Su posición en la lista fue: %d\n", lines, scanner.Text(), randomPerson)

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1

	var state string

	if randomNumber <= 55 {
		state = "infectado"
	} else if randomNumber <= 100 {
		state = "muerto"
	}

	return personName, state
}

func main() {
	log.Println("Iniciando continente")

	file, err := os.Open("continente/names.txt")
	failOnError(err, "Failed to open file")

	reader := bufio.NewReader(file)
	len, err := lineCounter(*reader)
	failOnError(err, "Failed to count lines")

	maxLines := len

	for len >= 0 {
		if len <= maxLines-5 {
			log.Println("Esperando 3 segundo para enviar más personas")
			time.Sleep(3 * time.Second)
		}

		name, state := getPerson()

		nameParts := strings.Split(name, " ")
		firstName := nameParts[0]
		lastName := nameParts[1]

		notificarPersona(omsIP, &pb.EstadoPersona{
			Nombre:   firstName,
			Apellido: lastName,
			Estado:   state,
		})

		log.Printf("Persona enviada: %s, con estado: %s\n", name, state)

		len -= 1
	}
}
