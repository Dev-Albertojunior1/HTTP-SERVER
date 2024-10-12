package handlers

import (
	"http-server/models"
	"net/http"
	"net/http/httputil"
)

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	items := models.GetAllItems()
	httputil.WriteJSON(w, http.StatusOK, items)
}
