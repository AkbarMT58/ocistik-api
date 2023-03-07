package controllers

import (
	"OCISTIK-API/config"
	"OCISTIK-API/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// READ ALL PENULIS
func AllPenulis(w http.ResponseWriter, r *http.Request) {
	var Penulis models.Master_penulis
	var Responses_penulis models.Responses_penulis
	var arrPenulis []models.Master_penulis

	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM master_penulis")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&Penulis.Nama_penulis, &Penulis.Created_date, &Penulis.Updated_date, &Penulis.Deleted_date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrPenulis = append(arrPenulis, Penulis)
		}
	}

	Responses_penulis.Status = 200
	Responses_penulis.Message = "SUCCESS"
	Responses_penulis.Data = arrPenulis

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(Responses_penulis)
}

// INSERT PENULIS
func InsertPenulis(w http.ResponseWriter, r *http.Request) {
	var responses models.Responses_penulis

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	nama_penulis := r.FormValue("nama_penulis")
	created_date := r.FormValue("created_date")

	_, err = db.Exec("INSERT INTO master_penulis(nama_penulis,created_date) VALUES(?,?)", nama_penulis, created_date)
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

// UPDATE PENULIS
func UpdatePenulis(w http.ResponseWriter, r *http.Request) {
	var responses models.Responses_kontak

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
	nama_penulis := r.FormValue("nama_penulis")
	updated_date := now

	_, err = db.Exec("UPDATE master_penulis SET nama=? updated_date=?  WHERE id=?", nama_penulis, updated_date, id)
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
func DeletePenulis(w http.ResponseWriter, r *http.Request) {
	var responses models.Responses_penulis

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	_, err = db.Exec("DELETE FROM master_penulis WHERE id=?", id)

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
