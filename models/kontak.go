package models

import "database/sql"

type Kontak struct {
	Id           int    `json:"id"`
	Nama         string `json:"nama"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
	Pesan        string `json:"pesan"`

	Created_date sql.NullString `json:"created_date"`
	Updated_date sql.NullString `json:"updated_date"`
	Deleted_date sql.NullString `json:"Deleted_date"`
}

type Responses_kontak struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Kontak
}
