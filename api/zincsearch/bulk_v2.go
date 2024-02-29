package zincsearch

import (
	"fmt"
	"net/http"
	"time"
)

func BulkV2(data interface{}) {
	requestStartTime := time.Now()

	// Send the POST with the Bulk data for indexing
	resp, err := Request(http.MethodPost, "/api/_bulkv2", data)
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
