package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	// 當為一個 DataBase
	CollectionName = "go-test-data"
)

func main() {
	sa := option.WithCredentialsFile("../../testgolang-98d3d-firebase-adminsdk-kqdra-f289ad5970.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		panic(fmt.Errorf("error getting Firestore client: %v", err))
	}
	defer client.Close()

	data := map[string]interface{}{
		"id":   1111123333,
		"any":  "dddddd",
		"name": "dddddadadsad",
	}
	// 
	// ---------------------Add document of auto id---------------------
	docRef, _ := AddAutoIDDocument(client, data)
	// AddAutoIDDocument(client, data)

	// --------------------- Add document of set id---------------------
	// AddSetIDDocument(client, "docName1", data)

	// ---------------------Get One Data---------------------
	re, _ := GetOneDocument(client, docRef.ID)
	fmt.Println(re)

	// ---------------------Get all Document---------------------
	// allre := GetAllDocument(client)
	// fmt.Println(allre)

	// ---------------------Update Data---------------------
	// updates := []firestore.Update{
	// 	{Path: "any", Value: "updatedValue"},
	// 	{Path: "newField", Value: "newValue"}, // Add a new field
	// }
	// UpdateDocument(client, "docName1", updates)

	// ---------------------Delete Data---------------------
	// DeleteDocument(client, "docName1")

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

func AddSetIDDocument(client *firestore.Client, docName string, data map[string]interface{}) error {
	_, err := client.Collection("go-test-data").Doc(docName).Set(context.Background(), data)
	if err != nil {
		panic(fmt.Errorf("error adding document: %v", err))
	} else {
		fmt.Println("Set ID Document added successfully!")
	}
	return err
}

// Get all documents from the collection
func GetAllDocument(client *firestore.Client) map[string]interface{} {
	iter := client.Collection("go-test-data").Documents(context.Background())
	data := make(map[string]interface{})

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(fmt.Errorf("error iterating over documents: %v", err))
		}
		data[doc.Ref.ID] = doc.Data()
	}
	return data
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

// Update Data
func UpdateDocument(client *firestore.Client, ID string, updates []firestore.Update) error {
	_, err := client.Collection("go-test-data").Doc(ID).Update(context.Background(), updates)
	if err != nil {
		panic(fmt.Errorf("error updating document: %v", err))
	} else {
		fmt.Println("Document updated successfully!")
	}
	return err
}

// Delete Data
func DeleteDocument(client *firestore.Client, ID string) error {
	_, err := client.Collection("go-test-data").Doc(ID).Delete(context.Background())
	if err != nil {
		panic(fmt.Errorf("error deleting document: %v", err))
	} else {
		fmt.Println("Document deleted successfully!")
	}
	return err
}
