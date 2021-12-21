package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	pb "client/management"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

const (
	address = "172.19.0.4:50051"
)

func main() {
	LevantarServidor()
}

func LevantarServidor() {
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})
	port := os.Getenv("PORT")
	if port == "" {
		port = "4300"
	}
	router.HandleFunc("/insert", insertData).Methods("POST")
	router.HandleFunc("/test", echoEndPoint).Methods("POST")
	fmt.Println("server in port " + port)
	http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(router))
}

func insertData(response http.ResponseWriter, request *http.Request) {
	data, errRead := ioutil.ReadAll(request.Body)
	if errRead != nil {
		response.Write([]byte("{\"error\":\"error en la entrada de datos\"}"))
		panic(errRead)
	}

	persona := Persona{}
	err := json.Unmarshal(data, &persona)
	if err != nil {
		response.Write([]byte("{\"error\":\"error en la decodificacion de datos\"}"))
		panic(err)
	}
	personaInsert := pb.User{Name: persona.Name, Age: int32(persona.Age), Location: persona.Location, VaccineType: persona.Vaccine_type, NDose: int32(persona.N_dose)}
	trigger := Export(personaInsert)
	if trigger {
		response.Write([]byte("{\"error\":\"error al insertar datos\"}"))
		panic("error al insertar datos")
	}
	response.Write([]byte("{\"respuesta\":\"insertado correctamente\"}"))
}

func echoEndPoint(response http.ResponseWriter, request *http.Request) {
	data, errRead := ioutil.ReadAll(request.Body)
	if errRead != nil {
		response.Write([]byte("{\"error\":\"error en la entrada de datos\"}"))
	}
	fmt.Println(string(data))
	response.Write(data)
}

func Export(user pb.User) bool {
	trigger := false
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		trigger = true
	}
	defer conn.Close()
	c := pb.NewUserManagmentClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CreateNewUser(ctx, &user)
	if err != nil {
		trigger = true
	}
	log.Printf(`User Details:
	NAME: %s
	AGE: %d`, r.GetName(), r.GetAge())
	return trigger
}

type Persona struct {
	Name         string `bson:"name,omitempty"`
	Location     string `bson:"location,omitempty"`
	Age          int    `bson:"age,omitempty"`
	Vaccine_type string `bson:"vaccine_type,omitempty"`
	N_dose       int    `bson:"n_dose,omitempty"`
}
