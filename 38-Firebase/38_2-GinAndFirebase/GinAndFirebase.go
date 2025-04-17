package main

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/", AAA)
	r.Run()
}

func AAA(ctx *gin.Context) {
	sa := option.WithCredentialsFile("../../../testgolang-98d3d-firebase-adminsdk-kqdra-7fa1a712f5.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		panic(fmt.Errorf("error getting Firestore client: %v", err))
	}
	defer client.Close()

	data := map[string]interface{}{
		"id":   223,
		"any":  "dddddd",
		"name": "dddddadadsad",
	}
	docRef, _ := AddAutoIDDocument(client, data)
	re, _ := GetOneDocument(client, docRef.ID)

	if re != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": re})
	}
}

func AddSetIDDocument(client *firestore.Client, docName string, data map[string]interface{}) error {
	_, err := client.Collection("").Doc(docName).Set(context.Background(), data)
	if err != nil {
		panic(fmt.Errorf("error adding document: %v", err))
	} else {
		fmt.Println("Set ID Document added successfully!")
	}
	return err
}

// Get Data
func GetOneDocument(client *firestore.Client, ID string) (map[string]interface{}, error) {
	doc, err := client.Collection("go-test-data").Doc(ID).Get(context.Background())
	if err != nil {
		panic(fmt.Errorf("error getting document: %v", err))
	} else {
		fmt.Println("Retrieved document data:", doc.Data())
	}
	return doc.Data(), err
}

func AddAutoIDDocument(client *firestore.Client, data map[string]interface{}) (*firestore.DocumentRef, error) {
	docRef, _, err := client.Collection("go-test-data").Add(context.Background(), data)
	if err != nil {
		panic(fmt.Errorf("error adding document: %v", err))
	} else {
		fmt.Println("Auto ID Document added successfully!")
	}
	return docRef, err
}
