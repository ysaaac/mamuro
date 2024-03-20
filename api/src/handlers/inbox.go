package handlers

import "net/http"

func InboxHandler(w http.ResponseWriter, r *http.Request) {
	page := GetPage(w, r)
	response := EmailList(page, "inbox")
	WriteJsonResponse(w, http.StatusOK, response)
}
