package main

import (
	"OCISTIK-API/controllers"
	"fmt"
	"log"
	"net/http"

	// "github.com/gin-contrib/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // change this later
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(204)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/getArtikel", controllers.AllArtikel).Methods(http.MethodGet)

	api.HandleFunc("/getArtikelbyslug/{slug}", controllers.Artikelbyslug).Methods(http.MethodGet)
	api.HandleFunc("/insertArtikel", controllers.InsertArtikel).Methods(http.MethodPost)
	api.HandleFunc("/updateArtikel", controllers.UpdateArtikel).Methods(http.MethodPut)
	api.HandleFunc("/deleteArtikel", controllers.DeleteArtikel).Methods(http.MethodDelete)

	api.HandleFunc("/getKontak", controllers.AllKontak).Methods(http.MethodGet)
	api.HandleFunc("/insertKontak", controllers.InsertKontak).Methods(http.MethodPost)
	api.HandleFunc("/updateKontak", controllers.UpdateKontak).Methods(http.MethodPut)
	api.HandleFunc("/deleteKontak", controllers.DeleteKontak).Methods(http.MethodDelete)

	api.HandleFunc("/getPenulis", controllers.AllPenulis).Methods(http.MethodGet)
	api.HandleFunc("/insertPenulis", controllers.InsertPenulis).Methods(http.MethodPost)
	api.HandleFunc("/updatePenulis", controllers.UpdatePenulis).Methods(http.MethodPut)
	api.HandleFunc("/deletePenulis", controllers.DeletePenulis).Methods(http.MethodDelete)

	api.HandleFunc("/getEditor", controllers.AllEditor).Methods(http.MethodGet)
	api.HandleFunc("/insertEditor", controllers.InsertEditor).Methods(http.MethodPost)
	api.HandleFunc("/updateEditor", controllers.UpdateEditor).Methods(http.MethodPut)
	api.HandleFunc("/deleteEditor", controllers.DeleteEditor).Methods(http.MethodDelete)

	c := cors.New(cors.Options{

		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Type", "Cache-Control", "Content-Type", "Xid"},
		Debug:            false,
	})
	handler := c.Handler(router)
	fmt.Print("CONNECTED TO PORT 9000")
	log.Fatal(http.ListenAndServe(":9000", handler))

}
