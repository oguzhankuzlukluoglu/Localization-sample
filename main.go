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
	files := []string{
		"locales/active.en.json",
		"locales/active.tr.json",
		"locales/active.es.json",
		"locales/active.fr.json",
		"locales/active.de.json",
		"locales/active.it.json",
		"locales/active.pt.json",
		"locales/active.ru.json",
		"locales/active.zh.json",
	}
	for _, file := range files {
		if _, err := bundle.LoadMessageFile(file); err != nil {
			log.Fatalf("Failed to load %s: %v", file, err)
		}
	}

	r := chi.NewRouter()
	r.Get("/localize", controllers.HandleLocalization(bundle))
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
