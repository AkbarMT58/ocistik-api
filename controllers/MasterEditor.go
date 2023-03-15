package controllers

import (
	"OCISTIK-API/config"
	"OCISTIK-API/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// READ ALL EDITOR
func AllEditor(w http.ResponseWriter, r *http.Request) {
	var Editor models.Master_editor
	var Responses_editor models.Responses_editor
	var arrEditor []models.Master_editor

	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM master_editor")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&Editor.Id, &Editor.Nama_editor, &Editor.Created_date, &Editor.Updated_date, &Editor.Deleted_date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrEditor = append(arrEditor, Editor)
		}
	}

	Responses_editor.Status = 200
	Responses_editor.Message = "SUCCESS"
	Responses_editor.Data = arrEditor

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(Responses_editor)
}

// READ ALL ARTIKEL BY ID
func EditorbyId(w http.ResponseWriter, r *http.Request) {
	var Editor models.Master_editor
	var Responses_editor models.Responses_editor
	var arrEditor []models.Master_editor

	db := config.ConnectDB()
	defer db.Close()

	params := r.URL

	vars := mux.Vars(r)
	id_ := vars["id"]

	fmt.Println(`url := `, params)

	rows, err := db.Query("SELECT * FROM master_editor where id=?", id_)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&Editor.Id, &Editor.Nama_editor, &Editor.Created_date, &Editor.Updated_date, &Editor.Deleted_date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrEditor = append(arrEditor, Editor)
		}
	}

	Responses_editor.Status = 200
	Responses_editor.Message = "SUCCESS"
	Responses_editor.Data = arrEditor

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(Responses_editor)
}

// INSERT EDITOR
func InsertEditor(w http.ResponseWriter, r *http.Request) {
	var responses models.Responses_editor

	loc, _ := time.LoadLocation("Asia/Jakarta")
	//set timezone,
	now := time.Now().In(loc)

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	nama_editor := r.FormValue("nama_editor")
	created_date := now

	_, err = db.Exec("INSERT INTO master_editor(nama_editor,created_date) VALUES(?,?)", nama_editor, created_date)
	if err != nil {
		log.Print(err)
		return
	}

	responses.Status = 200
	responses.Message = "INSERT DATA SUCCESFULLY"
	fmt.Print("INSERT DATA TO DATABASE")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(responses)
}

// UPDATE EDITOR
func UpdateEditor(w http.ResponseWriter, r *http.Request) {
	var responses models.Responses_editor

	loc, _ := time.LoadLocation("Asia/Jakarta")
	//set timezone,
	now := time.Now().In(loc)

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	nama_editor := r.FormValue("nama_editor")
	updated_date := now

	_, err = db.Exec("UPDATE master_editor SET nama_editor=? , updated_date=?  WHERE id=?", nama_editor, updated_date, id)
	if err != nil {
		log.Print(err)
	}

	responses.Status = 200
	responses.Message = "UPDATE DATA SUCCESFULLY"
	fmt.Print("UPDATE DATA SUCCESFULLY")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

// DELETE PENULIS
func DeleteEditor(w http.ResponseWriter, r *http.Request) {
	var responses models.Responses_penulis

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	_, err = db.Exec("DELETE FROM master_editor WHERE id=?", id)

	if err != nil {
		log.Print(err)
		return
	}
	responses.Status = 200
	responses.Message = "DELETE DATA SUCCESSFULLY"
	fmt.Print("DELETE DATA SUCCESSFULLY")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}
