package zincsearch

import (
	"api/models"
	"api/models/zincsearch_models"
	"fmt"
	"net/http"
	"time"
)

func BulkV2(indexName string, data []*models.Email) {
	requestStartTime := time.Now()

	bulkData := zincsearch_models.BulkV2{
		IndexName: indexName,
		Records:   data,
	}

	// Send the POST with the Bulk data for indexing
	resp, err := Request(http.MethodPost, "/api/_bulkv2", bulkData)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	response, err := HandleResponse(resp)
	if err != nil {
		fmt.Println("Error processing response:", err)
		return
	}

	// Prints how long take the request
	defer fmt.Printf("BulkV2 Processing took: %v\n", time.Since(requestStartTime))
	// Prints response
	fmt.Println("Response:", response)
}
