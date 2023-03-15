package models

import "database/sql"

type Master_editor struct {
	Id           int            `json:"id"`
	Nama_editor  string         `json:"nama_editor`
	Created_date sql.NullString `json:"created_date"`
	Updated_date sql.NullString `json:"updated_date"`
	Deleted_date sql.NullString `json:"deleted_date"`
}

type Responses_editor struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Master_editor
}
