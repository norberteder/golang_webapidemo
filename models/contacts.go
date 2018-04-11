package models

import "log"

type Contact struct {
	Id int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
}

func AllContacts() ([]*Contact, error) {
	rows, err := db.Query("SELECT id, firstName, lastName, email FROM contact")
	if err != nil {
		return nil, err
	}

	log.Println("Successfully loaded contacts from database")
	contacts := make([]*Contact, 0)
	for rows.Next() {
		contact := new(Contact)
		err := rows.Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.Email)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	return contacts, nil
}

func InsertContact(contact *Contact) (error){
	result, err := db.Exec("INSERT INTO contact (firstName, lastName, email) VALUES (?,?,?)", contact.FirstName, contact.LastName, contact.Email)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()

	contact.Id = int(id)

	return nil
}