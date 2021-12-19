package gRPC

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"servidor/gRPC/client"
	pb "servidor/gRPC/management"
	"servidor/structs"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _client *mongo.Client
var _context context.Context

func Connection() (*mongo.Client, context.Context) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	/*
		List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	return client, ctx
}

func InsertMongo(client *mongo.Client, ctx context.Context, data structs.Persona) bool {
	trigger := false
	client, ctx = Connection()

	collection := client.Database("SopesProyecto2").Collection("datos")

	/*
	  Insert documents
	*/
	var docs []interface{}
	docs = append(docs, bson.D{{"name", data.Name}, {"location", data.Location}, {"age", data.Age}, {"vaccine_type", data.Vaccine_type}, {"n_dose", data.N_dose}})

	res, insertErr := collection.InsertMany(ctx, docs)
	if insertErr != nil {
		trigger = true
		log.Fatal("insert error", insertErr)
	}
	fmt.Println(res)
	/*
		Iterate a cursor and print it
	*/
	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		trigger = true
		log.Fatal("find error", currErr)
	}
	defer cur.Close(ctx)

	var posts []structs.Prueba
	if err := cur.All(ctx, &posts); err != nil {
		trigger = true
		log.Fatal("cur all", err)
	}

	defer client.Disconnect(_context)
	return trigger
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
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/close", closeClient).Methods("GET")
	router.HandleFunc("/test", echoEndPoint).Methods("POST")
	router.HandleFunc("/insert", insertData).Methods("POST")
	fmt.Println("server in port " + port)
	http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(router))
}

func closeClient(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Cliente desconectado"))
}

func welcome(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("hello from go server"))
}

func insertData(response http.ResponseWriter, request *http.Request) {
	data, errRead := ioutil.ReadAll(request.Body)
	if errRead != nil {
		response.Write([]byte("{\"error\":\"error en la entrada de datos\"}"))
		panic(errRead)
	}
	persona := structs.Persona{}
	err := json.Unmarshal(data, &persona)
	if err != nil {
		response.Write([]byte("{\"error\":\"error en la decodificacion de datos\"}"))
		panic(err)
	}
	personaInsert := pb.User{Name: persona.Name, Age: int32(persona.Age), Location: persona.Location, VaccineType: persona.Vaccine_type, NDose: int32(persona.N_dose)}
	trigger := client.Export(personaInsert)
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
