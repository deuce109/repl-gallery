package repositories

import (
	"fmt"

	"labix.org/v2/mgo"
)

type RequestRepository struct {
	collection     *mgo.Collection
	collectionName string
}

func NewRequestRepository(db *mgo.Database) *RequestRepository {
	c := db.C("requests")

	fmt.Println("Connected to collection 'requests'")

	return &RequestRepository{
		c,
		"requests",
	}
}
