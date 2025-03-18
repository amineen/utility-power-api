package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID    string             `bson:"id,omitempty"`
	Name          string             `bson:"name,omitempty"`
	PhoneNumber   string             `bson:"phone_number,omitempty"`
	Code          string             `bson:"code,omitempty"`
	ServiceAreaID string             `bson:"service_area_id,omitempty"`
	Meters        []string           `bson:"meters,omitempty"`
	Balances      Balance            `bson:"balances,omitempty"`
	Active        bool               `bson:"active,omitempty"`
	LastHeartbeat *time.Time         `bson:"last_heartbeat,omitempty"`
	CreatedAt     time.Time          `bson:"created_at,omitempty"`
	UpdatedAt     time.Time          `bson:"updated_at,omitempty"`
}

type Balance struct {
	Credit MonitoryValue `bson:"credit,omitempty"`
}

type MonitoryValue struct {
	Value    string `bson:"value,omitempty"`
	Currency string `bson:"currency,omitempty"`
}
