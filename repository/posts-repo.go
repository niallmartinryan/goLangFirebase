package repository

import (
	"context"
	"log"
	"google.golang.org/api/iterator"
	"golang-firestore/entity"
	firebase "firebase.google.com/go"
  	"google.golang.org/api/option"
)

type PostRepository interface {
	AddPersonRepo(*entity.Person) (*entity.Person, error)
	EditPersonRepo(*entity.Person) (*entity.Person, error)
	FindAll() ([]entity.Person, error)

}

type repo struct{}

const (
	projectId      string = "accela-nr"
	collectionName string = "Person"
)

//NewPostRepository

func NewPostRepository() PostRepository {
	return &repo{}
}

func (*repo) AddPersonRepo(person *entity.Person) (*entity.Person, error) {

	ctx := context.Background()
	sa := option.WithCredentialsFile("./accela-nr-firebase-adminsdk-7nkyg-cd9d43ee9f.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
	  log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
	  log.Fatalln(err)
	}
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    person.ID,
		"firstName": person.FirstName,
		"lastName":  person.LastName,
		"postalCode" : "",
		"state" : "",
		"street" : "",
		"city" : "",
	})
	if err != nil {
		log.Fatalf("Failed to adding a new person: %v", err)
		return nil, err
	}
	
	return person, nil
}

func (*repo) EditPersonRepo(person *entity.Person) (*entity.Person, error) {

	ctx := context.Background()
	sa := option.WithCredentialsFile("./accela-nr-firebase-adminsdk-7nkyg-cd9d43ee9f.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
	  log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
	  log.Fatalln(err)
	}
	defer client.Close()

	// _, _, err = client.Collection(collectionName).Where("ID", "==" , person.ID).Documents()

	
	if err != nil {
		log.Fatalf("Failed to update a person: %v", err)
		return nil, err
	}

	return person, nil
}


func (*repo) FindAll() ([]entity.Person, error) {

	ctx := context.Background()
	sa := option.WithCredentialsFile("./accela-nr-firebase-adminsdk-7nkyg-cd9d43ee9f.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
	  log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
	  log.Fatalln(err)
	}
	defer client.Close()

	var people []entity.Person

	itr := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of people: %v", err)
			return nil, err
		}

		person := entity.Person{
			ID:    doc.Data()["ID"].(int64),
			FirstName: 	doc.Data()["firstName"].(string),
			LastName:  	doc.Data()["lastName"].(string),
			Street:  	doc.Data()["street"].(string),
			City:  		doc.Data()["city"].(string),
			State:  	doc.Data()["state"].(string),
			PostalCode: doc.Data()["postalCode"].(string),
		}
		people = append(people, person)
	}
	return people, nil
}




