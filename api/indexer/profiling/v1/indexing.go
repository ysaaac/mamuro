package v1

import (
	"api/models"
	"api/src"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func IndexData(folderPath string) {
	fmt.Printf("=================\n")
	// I'm defining a counter because I'm going to send a bulk every 1000 records due to bulk Zincsearch
	// specification, you can see more here -> https://zincsearch-docs.zinc.dev/api/document/bulk/#request-action
	//var inboxTypeItemCounter int
	//var sentItemsTypeItemCounter int

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// parent dir is the index type that can be inbox type or sent_items type
			parentDir := filepath.Base(filepath.Dir(path))

			if !strings.HasPrefix(info.Name(), ".") {
				fmt.Printf("File: %s, Index Type: %s\n", path, parentDir)

				email, parseEmailError := parseEmail(path)
				if parseEmailError != nil {
					fmt.Println("Error:", parseEmailError)
				}

				emailJSON, _ := json.MarshalIndent(email, "", "  ")
				//zincsearch.BulkV2()
				fmt.Println(string(emailJSON))
			}
		}
		return nil
	})

	if err != nil {
		return
	}
	fmt.Printf("=================\n")
	return
}

func parseHeader(email *models.Email, line string, currentLineIndex *int) {
	headerTag := models.HeadersList[*currentLineIndex]
	headerCompleteTag := headerTag + ":"

	if strings.Contains(line, headerCompleteTag) {
		extractedValueFromLine := strings.TrimSpace(strings.TrimPrefix(line, headerCompleteTag))
		headerTagNameInEmailStructureFormat := strings.Replace(headerTag, "-", "", -1)
		src.SetField(email, headerTagNameInEmailStructureFormat, extractedValueFromLine)
		*currentLineIndex++
	} else {
		// This is special for handling cases such as file
		// api/indexer/test_files/reitmeyer-j/inbox/7.
		// this file has a multiline at 'To:' header flag.
		pendingHeaderTag := models.HeadersList[*currentLineIndex-1]
		interfaceValue, _ := src.GetFieldValue(email, pendingHeaderTag)
		if currentStrValue, ok := interfaceValue.(string); ok {
			newValue := currentStrValue + " " + strings.TrimSpace(line)
			src.SetField(email, pendingHeaderTag, newValue)
		}
	}
}

func parseEmail(filename string) (*models.Email, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	email := &models.Email{}
	scanner := bufio.NewScanner(file)

	totalHeadersItems := len(models.HeadersList)
	currentLineIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		if totalHeadersItems > currentLineIndex {
			parseHeader(email, line, &currentLineIndex)
		} else {
			email.Content += "\n" + line
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return email, nil
}
