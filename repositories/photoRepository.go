package repositories

import (
	"fmt"
	"main/models"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type PhotoRepository struct {
	collection     *mgo.Collection
	collectionName string
}

func NewPhotoRepository(db *mgo.Database) *PhotoRepository {
	c := db.C("photos")

	fmt.Println("Connected to collection 'photos'")

	return &PhotoRepository{
		c,
		"photos",
	}
}

func (p *PhotoRepository) GetPhoto(id string) (photo *models.Photo, err error) {
	err = p.collection.FindId(id).One(&photo)
	if err != nil {
		// log error
		return nil, err
	}
	return photo, nil
}

func (p *PhotoRepository) GetPhotos(query interface{}) (photos []*models.Photo, err error) {
	err = p.collection.Find(query).All(photos)

	if err != nil {
		// log error
		return nil, err
	}
	return photos, nil
}

func (p *PhotoRepository) InsertPhotos(photos []*models.Photo) (int, error) {
	inserted := len(photos)
	err := p.collection.Insert(photos)
	if err != nil {
		// log error
		return 0, err
	}
	return inserted, nil
}

func (p *PhotoRepository) DeletePhoto(ids []string) int {
	deleted := 0
	for id := range ids {
		err := p.collection.Remove(id)
		if err == nil {
			deleted += 1
		} else {
			//log error
		}
	}

	return deleted
}
