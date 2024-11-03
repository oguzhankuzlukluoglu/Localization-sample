package main

import (
	"encoding/json"
	"log"
	"my-localization-project/controllers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	//yeni bir i18n dil paketi oluştur
	bundle := i18n.NewBundle(language.Turkish) //sayfa default olarak türkçe yüklensin
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	_, err := bundle.LoadMessageFile("locales/active.en.json")
	if err != nil {
		log.Fatal(err)
	}
	_, err = bundle.LoadMessageFile("locales/active.tr.json")
	if err != nil {
		log.Fatal(err)
	}
	r := chi.NewRouter()
	r.Get("/localize", controllers.HandleLocalization(bundle))
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
