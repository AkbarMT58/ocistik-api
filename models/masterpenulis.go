package models

import "database/sql"

type Master_penulis struct {
	Id           int            `json:"id"`
	Nama_penulis string         `json:"nama_penulis`
	Created_date sql.NullString `json:"created_date"`
	Updated_date sql.NullString `json:"updated_date"`
	Deleted_date sql.NullString `json:"Deleted_date"`
}

type Responses_penulis struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Master_penulis
}
