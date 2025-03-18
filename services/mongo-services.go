package services

import (
	"context"

	"github.com/amineen/utility-api/db"
	"github.com/amineen/utility-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCustomers() ([]models.Customer, error) {
	collection := db.CustomersCollection()

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var customers []models.Customer
	if err := cursor.All(context.TODO(), &customers); err != nil {
		return nil, err
	}
	return customers, nil
}
