package controllers

import (
	"OCISTIK-API/config"
	"OCISTIK-API/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

type Thelastid struct {
	Last_id string `json:"last_id"`
}

// READ ALL ARTIKEL
func AllArtikel(w http.ResponseWriter, r *http.Request) {
	var artikels models.Artikel
	var responses models.Responses
	var arrArtikel []models.Artikel

	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM artikel order by created_date desc")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&artikels.Id, &artikels.Title, &artikels.Description, &artikels.Konten, &artikels.Slug, &artikels.Slug_id, &artikels.Headline_id, &artikels.Picture, &artikels.Penulis, &artikels.Tag_name, &artikels.Meta_title, &artikels.Meta_description, &artikels.Meta_keyword, &artikels.Created_date, &artikels.Updated_date, &artikels.Deleted_date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrArtikel = append(arrArtikel, artikels)
		}
	}

	responses.Status = 200
	responses.Message = "SUCCESS"
	responses.Data = arrArtikel

	Get_Last_ID()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(responses)
}

func LoadImageFromURL(w http.ResponseWriter, r *http.Request) {
	//Get the response bytes from the url

	vars := mux.Vars(r)
	filename := vars["url_path"]

	filepath := "upload/" + filename

	f, e := os.Open(filepath)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	io.Copy(w, f)

}

// READ ALL ARTIKEL BY ID
func ArtikelbyId(w http.ResponseWriter, r *http.Request) {
	var artikels models.Artikel
	var responses models.Responses
	var arrArtikel []models.Artikel

	db := config.ConnectDB()
	defer db.Close()

	params := r.URL

	vars := mux.Vars(r)
	id_ := vars["id"]

	fmt.Println(`url := `, params)

	// fmt.Println(`slug := `, slug_)

	rows, err := db.Query("SELECT * FROM artikel where id=?", id_)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(
			&artikels.Id,
			&artikels.Title,
			&artikels.Description,
			&artikels.Konten,
			&artikels.Slug,
			&artikels.Slug_id,
			&artikels.Headline_id,
			&artikels.Picture,
			&artikels.Penulis,
			&artikels.Tag_name,
			&artikels.Meta_title,
			&artikels.Meta_description,
			&artikels.Meta_keyword,
			&artikels.Created_date,
			&artikels.Updated_date,
			&artikels.Deleted_date)
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
func ArtikelbySlug(w http.ResponseWriter, r *http.Request) {
	var artikels models.Artikel
	var responses models.Responses
	var arrArtikel []models.Artikel

	db := config.ConnectDB()
	defer db.Close()

	params := r.URL

	vars := mux.Vars(r)
	slug_ := vars["slug"]

	fmt.Println(`url := `, params)

	// fmt.Println(`slug := `, slug_)

	rows, err := db.Query("SELECT * FROM artikel where slug=?", slug_)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&artikels.Id, &artikels.Title, &artikels.Description, &artikels.Konten, &artikels.Slug, &artikels.Slug_id, &artikels.Headline_id, &artikels.Picture, &artikels.Penulis, &artikels.Tag_name, &artikels.Meta_title, &artikels.Meta_description, &artikels.Meta_keyword, &artikels.Created_date, &artikels.Updated_date, &artikels.Deleted_date)
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
	// var artikels models.Artikel
	// var arrArtikel []models.Artikel
	var slug_id int
	var last_id int

	loc, _ := time.LoadLocation("Asia/Jakarta")

	//set timezone,
	now := time.Now().In(loc)

	db := config.ConnectDB()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	//last id increment

	title := r.FormValue("title")
	description := r.FormValue("description")
	konten := r.FormValue("konten")
	headline_id := r.FormValue("headline_id")
	picture, handler, err := r.FormFile("picture")
	penulis := r.FormValue("penulis")
	tag_name := r.FormValue("tag_name")
	meta_title := r.FormValue("meta_title")
	meta_description := r.FormValue("meta_description")
	meta_keyword := r.FormValue("meta_keyword")

	created_date := now

	slugify := slug.Make(title)

	defer picture.Close()

	alias := ""

	defer picture.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename
	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	fileLocation := filepath.Join(dir, "upload", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, picture); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// lastID, err := db.Query("SELECT max(id) as last_id from artikel")

	// if err != nil {
	// 	log.Print(err)
	// }

	// for lastID.Next() {
	// 	err = lastID.Scan(&artikels.Last_id)
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	} else {
	// 		arrArtikel = append(arrArtikel, artikels)
	// 	}
	// }

	slug_id = 0

	_, err = db.Exec("INSERT INTO artikel(title,description,konten,slug,slug_id,headline_id,picture,penulis,tag_name,meta_title,meta_description,meta_keyword,created_date) VALUES(?, ?,?,?,?,?,?,?,?,?,?,?,?)", title, description, konten, slugify, slug_id, headline_id, filename, penulis, tag_name, meta_title, meta_description, meta_keyword, created_date)

	if err != nil {
		log.Fatal(err.Error())
	}

	responses.Status = 200
	responses.Message = "INSERT DATA SUCCESFULLY"

	fmt.Print("INSERT DATA TO DATABASE")

	fmt.Println("MASUK")
	fmt.Println("LAST ID:", last_id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)

}

// UPDATE ARTIKEL
func UpdateArtikel(w http.ResponseWriter, r *http.Request) {
	var responses models.Responses
	var filename_ string
	var filename_fin string

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
	title := r.FormValue("title")
	description := r.FormValue("description")
	konten := r.FormValue("konten")
	headline_id := r.FormValue("headline_id")
	picture_form := r.FormValue("picture_form")
	// picture, handler, err := r.FormFile("picture")
	penulis := r.FormValue("penulis")
	tag_name := r.FormValue("tag_name")
	meta_title := r.FormValue("meta_title")
	meta_description := r.FormValue("meta_description")
	meta_keyword := r.FormValue("meta_keyword")
	updated_date := now

	slugify := slug.Make(title)

	picture, handler, err := r.FormFile("picture")

	if err != nil || handler.Size == 0 {
		// file was not sent

	} else {
		// process file

		alias := ""

		defer picture.Close()

		dir, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filename := handler.Filename
		if alias != "" {
			filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
		}

		fileLocation := filepath.Join(dir, "upload", filename)
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, picture); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if picture_form == "" {

			filename_fin = filename

		}
		if picture_form != "" {

			filename_fin = picture_form

		}

	}

	_, err = db.Exec("UPDATE artikel SET title=?, description=?,konten=? ,slug=?, headline_id=?,picture=?,penulis=?,tag_name=?,meta_title=?,meta_description=?,meta_keyword=?, updated_date=? WHERE id=?", title, description, konten, slugify, headline_id, filename_fin, penulis, tag_name, meta_title, meta_description, meta_keyword, updated_date, id)

	if err != nil {
		log.Print(err)
	}

	fmt.Print("UPDATE DATA SUCCESFULLY")
	fmt.Println("id:", id)
	fmt.Println("desc:", description)
	fmt.Println("head:", headline_id)
	fmt.Println("penulis:", penulis)
	fmt.Println("update date:", updated_date)
	fmt.Println("slug:", slugify)
	fmt.Println("filename:", filename_)
	fmt.Println("picture form:", picture_form)

	responses.Status = 200
	responses.Message = "UPDATE  DATA SUCCESFULLY"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)

}

// SEARCH ARTIKEL BY TEXT
func SearchArtikel(w http.ResponseWriter, r *http.Request) {
	var artikels models.Artikel
	var responses models.Responses
	var arrArtikel []models.Artikel

	db := config.ConnectDB()
	defer db.Close()

	params_search := r.FormValue("search")

	rows, err := db.Query("SELECT * FROM artikel WHERE title like '%" + params_search + "%'")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(
			&artikels.Id,
			&artikels.Title,
			&artikels.Description,
			&artikels.Konten,
			&artikels.Slug,
			&artikels.Slug_id,
			&artikels.Headline_id,
			&artikels.Picture,
			&artikels.Penulis,
			&artikels.Tag_name,
			&artikels.Meta_title,
			&artikels.Meta_description,
			&artikels.Meta_keyword,
			&artikels.Created_date,
			&artikels.Updated_date,
			&artikels.Deleted_date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrArtikel = append(arrArtikel, artikels)
		}
	}

	responses.Status = 200
	responses.Message = "SUCCESS FIND DATA"
	responses.Data = arrArtikel

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

func Get_Last_ID() int {

	var data_lastid models.Artikel
	// var arrArtikel []models.Artikel

	db := config.ConnectDB()
	defer db.Close()

	last_id, _ := db.Exec("SELECT  max(id) FROM artikel")

	// data_lastid = append(data_lastid, last_id)

	if last_id != nil {
		log.Print(last_id)

		fmt.Println("check data last id :", data_lastid.Id)

	}

	return data_lastid.Id

}
