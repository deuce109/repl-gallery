package models

import (
	"encoding/json"

	"labix.org/v2/mgo/bson"
)

type Photo struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Data        []byte        `bson:"data" json:"data"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
}

func (p *Photo) ToJson() ([]byte, error) {
	return json.Marshal(p)
}
