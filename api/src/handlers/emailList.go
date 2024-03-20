package handlers

import (
	"api/zincsearch"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type Total struct {
	Value int `json:"value"`
}

type Source struct {
	Date      string `json:"date"`
	From      string `json:"from"`
	MessageId string `json:"message_id"`
	Subject   string `json:"subject"`
	To        string `json:"to,omitempty"`
	Content   string `json:"content,omitempty"`
}

type Hit struct {
	Id        string      `json:"_id"`
	Score     json.Number `json:"_score"`
	Timestamp time.Time   `json:"@timestamp"`
	Source    Source      `json:"_source"`
}

type Hits struct {
	Total Total `json:"total"`
	Hits  []Hit `json:"hits"`
}

type SearchResponse struct {
	Hits Hits `json:"hits"`
}

type QueryAllDocuments struct {
	SearchType string   `json:"search_type"`
	SortFields []string `json:"sort_fields"`
	From       int      `json:"from"`
	MaxResults int      `json:"max_results"`
	Source     []string `json:"_source"`
}

type TermFilter struct {
	Field string `json:"field"`
	Term  string `json:"term"`
}

type QueryByTerm struct {
	SearchType string     `json:"search_type"`
	Query      TermFilter `json:"query"`
	Source     []string   `json:"_source"`
}

func EmailList(from int, target string) SearchResponse {
	queryAllDocuments := QueryAllDocuments{
		SearchType: "match_all",
		SortFields: []string{"-date"},
		From:       from,
		MaxResults: 20,
		Source:     []string{"message_id", "from", "subject", "date"},
	}

	//fmt.Println("Request Body:\n", queryAllDocuments)

	queryResult := RequestEmail(target, queryAllDocuments)

	sort.Slice(queryResult.Hits.Hits, func(i, j int) bool {
		date1, _ := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700 (MST)", queryResult.Hits.Hits[i].Source.Date)
		date2, _ := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700 (MST)", queryResult.Hits.Hits[j].Source.Date)
		return date1.After(date2)
	})

	return queryResult
}

func EmailByTerm(target string, field string, term string) SearchResponse {
	queryByTerm := QueryByTerm{
		SearchType: "term",
		Query:      TermFilter{Field: field, Term: term},
		Source:     []string{"message_id", "from", "to", "subject", "content", "date"},
	}

	queryResult := RequestEmail(target, queryByTerm)
	return queryResult
}

func RequestEmail(target string, query interface{}) SearchResponse {
	inboxList, inboxListError := zincsearch.Request(http.MethodPost, "/api/"+target+"/_search", query)
	if inboxListError != nil {
		fmt.Printf("Error making request: %s\n", inboxListError)
		return SearchResponse{}
	}

	defer inboxList.Body.Close()

	body, readErr := ioutil.ReadAll(inboxList.Body)
	if readErr != nil {
		fmt.Printf("Error reading response body: %s\n", readErr)
		return SearchResponse{}
	}

	//fmt.Printf("JSON Response: %s\n", body)

	var result SearchResponse

	// Unmarshal the JSON data into the struct
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %s\n", err)
		return SearchResponse{}
	}
	return result
}

func GetPage(w http.ResponseWriter, r *http.Request) int {
	page := r.URL.Query().Get("page")
	var from int
	if page == "" {
		from = 0
	} else {
		// Converts string param to integer
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			http.Error(w, "Invalid parameter value", http.StatusBadRequest)
		}
		// 20 is because I'm putting 20 as max_result so for page 1 starts 20, page 2 -> 40 and so on :)
		from = 20 * pageInt
	}
	return from
}
