package zincsearch

import (
	"api/models/zincsearch_models"
	"fmt"
	"log"
	"net/http"
)

func CreateIndices() {
	createIndexFor("inbox", zincsearch_models.InboxIndexTypeStructure())
	createIndexFor("sent_items", zincsearch_models.SentItemsIndexTypeStructure())
}

func createIndexFor(indexName string, indexConfig zincsearch_models.CreateIndex) {
	if exists, _ := indexExists(indexName); !exists {
		resp, err := Request(http.MethodPost, "/api/index/", indexConfig)
		if err != nil {
			log.Println("Error: ", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Println("Index created (Status 200) -> ", indexName)
		} else {
			log.Println("Index: ", indexName, " wasn't created.")
		}
	}
}

func indexExists(indexName string) (bool, error) {
	resp, err := Request(http.MethodHead, "/api/index/", nil)
	if err != nil {
		return false, fmt.Errorf("error making request %s ", err)
	}
	defer resp.Body.Close()

	// Check the response status code to valid if index exists or not
	// https://zincsearch-docs.zinc.dev/api/index/exists/#request

	if resp.StatusCode == http.StatusOK {
		// Index found
		//fmt.Println("Request succeeded (Status 200) -> ", indexName, " found")
		return true, nil
	} else if resp.StatusCode == http.StatusNotFound {
		// Index don't exist
		//fmt.Println("Resource -> ", indexName, " not found (Status 404)")
		return false, nil
	} else {
		return false, fmt.Errorf("Unexpected status code: %d\n", resp.StatusCode)
	}
}
