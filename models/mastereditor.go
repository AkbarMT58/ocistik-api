package models

type Master_editor struct {
	Id           int    `json:"id"`
	Nama_editor  string `json:"nama_editor`
	Created_date string `json:"created_date"`
	Updated_date string `json:"updated_date"`
	Deleted_date string `json:"Deleted_date"`
}

type Responses_editor struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Master_editor
}
