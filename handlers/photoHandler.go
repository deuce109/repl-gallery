package handlers

import (
	"encoding/json"
	"fmt"
	mux "github.com/gorilla/mux"
	models "main/models"
	repositories "main/repositories"
	"net/http"
)

type PhotoHandler struct {
	repository *repositories.PhotoRepository
	Subrouter  *mux.Router
}

// {base_url}/

func (p *PhotoHandler) getPhotos(w http.ResponseWriter, req *http.Request) {

	var query map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&query)

	if err != nil {
		w.WriteHeader(500)
		// log error
		return
	}

	photos, err := p.repository.GetPhotos(query)

	if err != nil {
		w.WriteHeader(500)
		// log error
		return
	} else if photos == nil {
		w.WriteHeader(404)
		return
	} else if len(photos) == 0 {
		w.WriteHeader(204)
		return
	}

	data, err := json.Marshal(photos)

	if err != nil {
		w.WriteHeader(500)
		// log error
		return
	}
	w.Write(data)

}

func (p *PhotoHandler) insertPhotos(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		w.WriteHeader(500)
		// log error
		return
	}

	photo := &models.Photo{
		Description: req.FormValue("description"),
		Name:        req.FormValue("name"),
		Data:        []byte(req.FormValue("file")),
	}

	fmt.Printf("%s", photo)

	w.WriteHeader(200)

}

// {base_url}/:id

func (p *PhotoHandler) getPhotoById(w http.ResponseWriter, req *http.Request) {
	photo, err := p.repository.GetPhoto(mux.Vars(req)["id"])
	if err != nil {
		w.WriteHeader(500)
		// log error
		return
	} else if photo == nil {
		w.WriteHeader(404)
		return
	}

	jsonString, err := photo.ToJson()
	if err != nil {
		w.WriteHeader(500)
		// log error
		return
	}
	w.Write(jsonString)
}

func (p *PhotoHandler) updatePhoto(w http.ResponseWriter, req *http.Request) {

}

func (p *PhotoHandler) deletePhoto(w http.ResponseWriter, req *http.Request) {

}

// {base_url}/:id/data

func (p *PhotoHandler) getPhotoDataById(w http.ResponseWriter, req *http.Request) {

}

func (p *PhotoHandler) Init() {
	p.repository = &repositories.PhotoRepository{}

	baseRoute := p.Subrouter.Path("/")

	baseRoute.Methods("GET").HandlerFunc(p.getPhotos)
	baseRoute.Methods("PUT").HandlerFunc(p.insertPhotos)

	idRoute := p.Subrouter.Path("/:id")
	idRoute.Methods("GET").HandlerFunc(p.getPhotoById)
	idRoute.Methods("DELETE").HandlerFunc(p.deletePhoto)
	idRoute.Methods("PATCH").HandlerFunc(p.updatePhoto)
	idRoute.Path("/data").Methods("GET").HandlerFunc(p.getPhotoDataById)
}
