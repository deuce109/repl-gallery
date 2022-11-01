package models

import (
	"encoding/json"

	bson "labix.org/v2/mgo/bson"
)

type artType string

const (
	Digital = "digital"
	Oil     = "oil"
	Acrylic = "acrylic"
)

type contactInfo struct {
	PhoneNumber   string `bson:"phoneNumber" json:"phoneNumber" faker:"phone_number"`
	Email         string `bson:"email" json:"email" faker:"email"`
	StreetAddress string `bson:"address" json:"address"`
}

type address struct {
	City    string `bson:"city" json:"city"`
	State   string `bson:"state" json:"state"`
	Country string `bson:"country" json:"country"`
	ZipCode int    `bson:"zip" json:"zip"`
}

type size struct {
	Height int    `bson:"height" json:"height"`
	Width  int    `bson:"width" json:"width"`
	Units  string `bson:"unit" json:"unit"`
}

type Request struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	ContactInfo contactInfo   `bson:"contact" json:"contact"`
	Address     address       `bson:"address" json:"address"`
	Size        size          `bson:"size" json:"size"`
	References  [][]byte      `bson:"references" json:"references"`
	Description string        `bson:"description" json:"description"`
	Type        artType       `bson:"type" json:"type"`
}

func (request *Request) ToJson() ([]byte, error) {
	return json.Marshal(&request)
}
