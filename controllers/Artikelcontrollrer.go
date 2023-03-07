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

// READ ALL ARTIKEL
func AllArtikel(w http.ResponseWriter, r *http.Request) {
	var artikels models.Artikel
	var responses models.Responses
	var arrArtikel []models.Artikel

	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM artikel")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&artikels.Id, &artikels.Title, &artikels.Description, &artikels.Slug, &artikels.Headline_id, &artikels.Picture, &artikels.Created_date, &artikels.Updated_date, &artikels.Deleted_date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrArtikel = append(arrArtikel, artikels)
		}
	}

	responses.Status = 200
	responses.Message = "SUCCESS"
	responses.Data = arrArtikel

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(responses)
}

// READ ALL ARTIKEL BY SLUG
func Artikelbyslug(w http.ResponseWriter, r *http.Request) {
	var artikels models.Artikel
	var responses models.Responses
	var arrArtikel []models.Artikel

	db := config.ConnectDB()
	defer db.Close()

	params := r.URL

	// slug_ := params.Query().Get("slug")

	vars := mux.Vars(r)
	slug_ := vars["slug"]

	fmt.Println(`url := `, params)

	fmt.Println(`slug := `, slug_)

	rows, err := db.Query("SELECT * FROM artikel where slug=?", slug_)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&artikels.Id, &artikels.Title, &artikels.Description, &artikels.Slug, &artikels.Headline_id, &artikels.Picture, &artikels.Created_date, &artikels.Updated_date, &artikels.Deleted_date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrArtikel = append(arrArtikel, artikels)
		}
	}

	responses.Status = 200
	responses.Message = "SUCCESS"
	responses.Data = arrArtikel

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(responses)
}

// INSERT ARTIKEL
func InsertArtikel(w http.ResponseWriter, r *http.Request) {

	var responses models.Responses

	loc, _ := time.LoadLocation("Asia/Jakarta")

	//set timezone,
	now := time.Now().In(loc)

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	title := r.FormValue("title")
	description := r.FormValue("description")
	slug := r.FormValue("slug")
	headline_id := r.FormValue("headline_id")
	picture := r.FormValue("picture")
	created_date := now

	_, err = db.Exec("INSERT INTO artikel(title,description,slug,headline_id,picture,created_date) VALUES(?, ?, ?,?,?,?)", title, description, slug, headline_id, picture, created_date)
	if err != nil {
		log.Fatal(err.Error())
	}

	responses.Status = 200
	responses.Message = "INSERT DATA SUCCESFULLY"

	fmt.Print("INSERT DATA TO DATABASE")

	fmt.Println("MASUK")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(responses)
}

// UPDATE ARTIKEL
func UpdateArtikel(w http.ResponseWriter, r *http.Request) {
	var responses models.Responses

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	title := r.FormValue("title")
	description := r.FormValue("description")
	slug := r.FormValue("slug")
	headline_id := r.FormValue("headline_id")
	picture := r.FormValue("picture")
	created_date := r.FormValue("created_date")

	_, err = db.Exec("UPDATE artikel SET title=?, description=?,picture=? ,slug=?, headline_id=?, created_date=? WHERE id=?", title, description, picture, slug, headline_id, created_date, id)

	if err != nil {
		log.Print(err)
	}

	responses.Status = 200
	responses.Message = "UPDATE  DATA SUCCESFULLY"
	fmt.Print("UPDATE DATA SUCCESFULLY")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

// DELETE ARTIKEL
func DeleteArtikel(w http.ResponseWriter, r *http.Request) {
	var response models.Responses

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	_, err = db.Exec("DELETE FROM artikel WHERE id=?", id)

	if err != nil {
		log.Print(err)
		return
	}

	response.Status = 200
	response.Message = "DELETE DATA SUCCESSFULLY"
	fmt.Print("DELETE DATA SUCCESSFULLY")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
