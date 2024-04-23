package handlers

import (
	"net/http"

	jh "github.com/dpurbosakti/go-native/pkg/jsonhelper"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	jh.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "pong",
	})
}
