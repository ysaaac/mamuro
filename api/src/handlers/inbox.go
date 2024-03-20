package handlers

import "net/http"

func InboxHandler(w http.ResponseWriter, r *http.Request) {
	page := GetPage(w, r)
	objId := r.URL.Query().Get("id")
	search := r.URL.Query().Get("search")
	target := "inbox"

	var response SearchResponse

	if objId != "" {
		response = EmailByTerm(target, "_id", objId)
	} else if search != "" {
		response = EmailSearchBy(target, search)
	} else {
		response = EmailList(page, target)
	}

	WriteJsonResponse(w, http.StatusOK, response)
}
