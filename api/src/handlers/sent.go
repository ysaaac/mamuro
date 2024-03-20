package handlers

import "net/http"

func SentHandler(w http.ResponseWriter, r *http.Request) {
	page := GetPage(w, r)
	objId := r.URL.Query().Get("id")
	target := "sent_items"

	var response SearchResponse

	if objId != "" {
		response = EmailByTerm(target, "_id", objId)
	} else {
		response = EmailList(page, target)
	}
	WriteJsonResponse(w, http.StatusOK, response)
}
