package handlers

import (
	"api/config"
	v1 "api/indexer/profiling/v1"
	"api/zincsearch"
	"fmt"
	"net/http"
	"sync"
)

func indexingBackgroundTask(wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()
	fmt.Println("Indexing background task started...")
	filesPath := config.GetEnv("DATA_FILES_PATH", "./indexer/test_files")
	v1.IndexData(filesPath)

	// Signal that the background task has completed
	done <- true
}

func IndexData() ResponseStructure {
	zincsearch.CreateIndices()

	// Due to amount of files if we sent this http request and wait until a response it will return us an HTTP Error
	// 408 Request Timeout, that's why it executes the indexing request at workers, so
	// I create a wait group and a channel for signaling the background task to finish
	var wg sync.WaitGroup
	done := make(chan bool)

	// Start the background task in a goroutine
	wg.Add(1)
	go indexingBackgroundTask(&wg, done)

	response := ResponseStructure{
		Status:  "ok",
		Message: "Indexing has been started with workers",
	}

	go func() {
		// Wait for the background task to complete
		wg.Wait()

		// Close the done channel to signal that the background task has completed
		close(done)
	}()

	return response
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	response := IndexData()
	WriteJsonResponse(w, http.StatusOK, response)
}
