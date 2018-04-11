package controller

import (
	"net/http"
	"sampleWebApi/models"
	"encoding/json"
	"log"
)

func AllContacts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint AllContacts")
	contacts, err := models.AllContacts()
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(contacts)
	log.Println("Sent all contacts")
}

func InsertContact(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint InsertContact")
	var c models.Contact
	if r.Body == nil {
		http.Error(w, "Please send any body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = models.InsertContact(&c)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Inserted new contact")
}