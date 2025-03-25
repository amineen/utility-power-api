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
}

type Balance struct {
	Credit MonitoryValue `bson:"credit,omitempty"`
}

type MonitoryValue struct {
	Value    string `bson:"value,omitempty"`
	Currency string `bson:"currency,omitempty"`
}

type Payment struct {
	ID            string             `bson:"id,omitempty"`
	TimeStamp     time.Time          `bson:"timestamp,omitempty"`
	RecipientId   string             `bson:"recipient_id,omitempty"`
	CustomerId    string             `bson:"customer_id,omitempty"`
	Amount        Amount             `bson:"amount,omitempty"`
	Memo          string             `bson:"memo,omitempty"`
	ExternalId    string             `bson:"external_id,omitempty"`
	ServiceAreaID primitive.ObjectID `bson:"service_area_id,omitempty"`
	Vendor        string             `bson:"vendor,omitempty"`
	VendorId      string             `bson:"vendorId,omitempty"`
}

type Amount struct {
	Value  string  `bson:"value,omitempty"`
	Amount string  `bson:"amount,omitempty"`
	Kwh    float32 `bson:"kwh,omitempty"`
}
