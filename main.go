package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Client struct {
	ID   string `json: "id"`
	Name string `json: "name"`
	Age  int    `json: "age"`
}

var Clients = []Client{}

func main() {
	r := gin.Default()

	clientRoutes := r.Group("/clients")

	{
		clientRoutes.GET("/", getAllClients)
		clientRoutes.POST("/newClient", createClient)
		clientRoutes.PUT("/editClient/:id", editCLient)
		clientRoutes.DELETE("/delete/:id", deleteClient)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}

func getAllClients(ctx *gin.Context) {

	if len(Clients) == 0 {
		ctx.JSON(401, gin.H{
			"message": "Não há usuarios cadastrados",
		})
	}
	ctx.JSON(200, Clients)
}

func createClient(ctx *gin.Context) {
	var reqBody Client

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
	}

	reqBody.ID = uuid.New().String()

	Clients = append(Clients, reqBody)

	ctx.JSON(200, gin.H{
		"error": false,
	})
}

func editCLient(ctx *gin.Context) {
	id := ctx.Param("id")

	var body Client

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
	}

	for index, client := range Clients {
		if client.ID == id {
			Clients[index].Name = body.Name
			Clients[index].Age = body.Age

			ctx.JSON(200, gin.H{
				"error": false,
				"att":   Clients[index],
			})
			return
		}
	}

	ctx.JSON(404, gin.H{
		"err":     true,
		"message": "invalid user id",
	})
}

func deleteClient(ctx *gin.Context) {
	id := ctx.Param("id")

	for index, client := range Clients {
		if client.ID == id {
			Clients = append(Clients[:index], Clients[index+1:]...)

			ctx.JSON(200, gin.H{
				"error": false,
				"att":   Clients,
			})
			return
		}
	}

	ctx.JSON(404, gin.H{
		"err":     true,
		"message": "invalid user id",
	})
}
