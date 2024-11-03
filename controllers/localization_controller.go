package controllers

import (
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func HandleLocalization(bundle *i18n.Bundle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userLang:=r.Header.Get("Accept-Language")
		localizer:=i18n.NewLocalizer(bundle,userLang)

		result,err:= localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "WelcomeMessage", //uygulama hangi mesajı göndersin onu belirliyoruz
		})
		if err!=nil{
			result="Çeviri bulunamadı"
		}
		w.Write([]byte(result))
	}

}
