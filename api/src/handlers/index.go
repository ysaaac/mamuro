package handlers

import (
	"api/zincsearch"
	"net/http"
)

func IndexData() ResponseStructure {
	zincsearch.CreateIndices()
	// Due to amount of files if we sent this http request and wait until a response it will return us an HTTP Error
	// 408 Request Timeout, that's why it executes the indexing request at workers
	//v1.IndexData()
	return ResponseStructure{
		Status:  "ok",
		Message: "Indexing has been started with workers",
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	response := IndexData()
	WriteJsonResponse(w, http.StatusOK, response)
}
