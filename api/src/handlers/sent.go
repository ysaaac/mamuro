package handlers

import "net/http"

func SentHandler(w http.ResponseWriter, r *http.Request) {
	page := GetPage(w, r)
	response := EmailList(page, "sent_items")
	WriteJsonResponse(w, http.StatusOK, response)
}
