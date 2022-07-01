package configs

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var App *firebase.App
var HTTP_PORT int

func InitDb() {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("../blueprint/key.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

}

func InitPort() {
	HTTP_PORT = 80
}
