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

// READ ALL KONTAK
func AllKontak(w http.ResponseWriter, r *http.Request) {
	var Kontaks models.Kontak
	var responses_kontak models.Responses_kontak
	var arrKontak []models.Kontak

	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM kontak")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&Kontaks.Id, &Kontaks.Nama, &Kontaks.Phone_number, &Kontaks.Email, &Kontaks.Pesan, &Kontaks.Created_date, &Kontaks.Updated_date, &Kontaks.Deleted_date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrKontak = append(arrKontak, Kontaks)
		}
	}

	responses_kontak.Status = 200
	responses_kontak.Message = "SUCCESS"
	responses_kontak.Data = arrKontak

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(responses_kontak)
}

// INSERT KONTAK
func InsertKontak(w http.ResponseWriter, r *http.Request) {
	var responses models.Responses_kontak

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	nama := r.FormValue("nama")
	phone_number := r.FormValue("phone_number")
	email := r.FormValue("email")
	pesan := r.FormValue("pesan")

	// int_nik, _ := strconv.Atoi(nik)

	_, err = db.Exec("INSERT INTO kontak(nama,phone_number,email,pesan) VALUES(?,?, ?, ?)", nama, phone_number, email, pesan)
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

// UPDATE KONTAK
func UpdateKontak(w http.ResponseWriter, r *http.Request) {
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
	nama := r.FormValue("nama")
	phone_number := r.FormValue("phone_number")
	email := r.FormValue("email")
	pesan := r.FormValue("pesan")
	update_date := now

	_, err = db.Exec("UPDATE kontak SET nama=?, phone_number=?, email=?, pesan=?, updated_date=?  WHERE id=?", nama, phone_number, email, pesan, update_date, id)
	if err != nil {
		log.Print(err)
	}

	responses.Status = 200
	responses.Message = "UPDATE DATA SUCCESFULLY"
	fmt.Print("UPDATE DATA SUCCESFULLY")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

// DELETE KONTAK
func DeleteKontak(w http.ResponseWriter, r *http.Request) {
	var responses models.Responses_kontak

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	_, err = db.Exec("DELETE FROM kontak WHERE id=?", id)

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
