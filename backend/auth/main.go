package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const mongodbUri = "mongodb://localhost:27017"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	port := 4000
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))
	if err != nil {
		panic(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	/*
		POST /users
		Crea un usuario en la base de datos
	*/
	handler1 := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			createUser(client, ctx, res, req)
		}
	})

	handler1WithCors := c.Handler(handler1)
	/*
		POST /logIn
		Verifica credenciales de usuario y retorna token si es
	*/
	handler2 := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			res.Header().Set("Access-Control-Allow-Origin", "*")
			logIn(client, ctx, res, req)
		}
	})
	handler2WithCors := c.Handler(handler2)

	http.Handle("/users", handler1WithCors)
	http.Handle("/logIn", handler2WithCors)

	fmt.Printf("Starting server on %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func createUser(client *mongo.Client, ctx context.Context, res http.ResponseWriter, req *http.Request) {
	fmt.Println("Starting createUser")
	coll := client.Database("authorization").Collection("users")

	// Guardar información de usuario en Objeto
	var user User
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&user)
	if err != nil {
		http.Error(res, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	// Verificación de que el usuario ingresado no exista
	var existingUser User
	filter := bson.D{{Key: "email", Value: user.Email}}
	errFinder := coll.FindOne(ctx, filter).Decode(&existingUser)
	if errFinder == nil {
		http.Error(res, "El usuario ya existe", http.StatusConflict)
		return
	}

	// Encriptación de contraseña
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(bytes)

	// Guardado de usuario en Base de datos
	insertResult, err := coll.InsertOne(ctx, user)

	if err != nil {
		log.Fatal(err)
	}

	res.WriteHeader(http.StatusCreated)

	json.NewEncoder(res).Encode(map[string]interface{}{
		"result": "Usuario creado",
	})

	fmt.Println("Documento insertado con ID:", insertResult.InsertedID)
}

func logIn(client *mongo.Client, ctx context.Context, res http.ResponseWriter, req *http.Request) {
	fmt.Println("Starting logIn")
	coll := client.Database("authorization").Collection("users")

	// Guardado de datos de usuario en Objeto
	var user User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(res, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	// Busqueda de usuario en BD
	var existingUser User
	filter := bson.D{{Key: "email", Value: user.Email}}
	errFinder := coll.FindOne(ctx, filter).Decode(&existingUser)
	if errFinder != nil {
		http.Error(res, "Error al verificar el usuario", http.StatusInternalServerError)
		return
	}

	// Verificación de contraseña encriptada con contraseña recibida
	errPass := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if errPass != nil {
		http.Error(res, "Contraseña incorrecta", http.StatusInternalServerError)
		return
	}

	// Creación de token jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": existingUser.Email,
		"name":  existingUser.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtString, err := token.SignedString([]byte("secret"))
	if errPass != nil {
		panic(err)
	}
	json.NewEncoder(res).Encode(map[string]interface{}{
		"token": jwtString,
	})

	fmt.Println("Usuario logeado con token", jwtString)

}
