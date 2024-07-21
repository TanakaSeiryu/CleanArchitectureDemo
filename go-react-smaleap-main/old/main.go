package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	_ "github.com/lib/pq"
)

func Router(r *chi.Mux) {
	r.Get("/grades", getGradesHandler)
	r.Post("/grades", registerGradeHandler)
}
func main() {
	err := Connect()
	if err != nil {
		log.Fatal("DB接続に失敗", err)
	}

	r := chi.NewRouter()
	Router(r)

	port := ":8080"
	fmt.Printf("サーバーをポート%sで起動中...\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("サーバーの起動に失敗しました: %s", err)
	}
}
