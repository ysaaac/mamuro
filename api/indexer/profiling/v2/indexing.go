package v1

import (
	"api/models"
	"api/src"
	"api/zincsearch"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type EmailTypeBoxHelper struct {
	EmailQty   int
	EmailsList []*models.Email
}

func IndexData(folderPath string) {
	fmt.Printf("=================\n")
	// I'm defining a counter because I'm going to send a bulk every 1000 records due to bulk Zincsearch
	// specification, you can see more here -> https://zincsearch-docs.zinc.dev/api/document/bulk/#request-action
	inboxTypeHelper := EmailTypeBoxHelper{
		EmailQty: 0,
	}
	sentItemsTypeHelper := EmailTypeBoxHelper{
		EmailQty: 0,
	}

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// parent dir is the index type that can be inbox type or sent_items type
			parentDir := filepath.Base(filepath.Dir(path))

			if !strings.HasPrefix(info.Name(), ".") {
				//fmt.Printf("File: %s, Index Type: %s\n", path, parentDir)
				if parentDir == "inbox" {
					handleEmailsTypes(path, &inboxTypeHelper, "inbox")
				} else if parentDir == "sent_items" {
					handleEmailsTypes(path, &sentItemsTypeHelper, "sent_items")
				}
			}
		}
		return nil
	})

	// Send the remaining documents
	if inboxTypeHelper.EmailQty > 0 {
		zincsearch.BulkV2("inbox", inboxTypeHelper.EmailsList)
	}
	if sentItemsTypeHelper.EmailQty > 0 {
		zincsearch.BulkV2("sent_items", sentItemsTypeHelper.EmailsList)
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("=================\n")
	return
}

func handleEmailsTypes(path string, typeHelper *EmailTypeBoxHelper, indexName string) {
	resLock := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	parseEmail(path, typeHelper, wg, resLock)
	wg.Wait()
	// for Zincsearch spec I split the qty sent by 1000 items into bulk
	// https://zincsearch-docs.zinc.dev/api/document/bulkv2/
	//fmt.Printf("EmailQty: %d EmailsListLength: %d \n", typeHelper.EmailQty, len(typeHelper.EmailsList))
	if typeHelper.EmailQty == 1000 {
		// I sent all the data to Zincsearch bulk
		zincsearch.BulkV2(indexName, typeHelper.EmailsList)
		// I restart the helper with initial values
		typeHelper.EmailsList = []*models.Email{}
		typeHelper.EmailQty = 0
	}
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

// parseEmail gets all data of the documents
func parseEmail(filename string, result *EmailTypeBoxHelper, wg *sync.WaitGroup, resLock *sync.Mutex) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer file.Close()

		email := &models.Email{}
		scanner := bufio.NewScanner(file)

		// Initial buffer size
		const initialBufferSize = 64 * 1024 // 64 KB
		maxBufferSize := initialBufferSize

		// Attempt to scan with increasing buffer sizes because sometimes I'm getting
		// bufio.Scanner: token too long and set a too long size can be a problem for
		// memory corrupted data and also performance can be downgraded. So, It will
		// increase only when necessary
		for {
			// Set the buffer size
			buf := make([]byte, maxBufferSize)
			scanner.Buffer(buf, maxBufferSize)
			// Reset line index for each attempt
			currentLineIndex := 0
			// Reset email content for each attempt
			email.Content = ""

			// Reset scanner error
			var scannerErr error

			for scanner.Scan() {
				line := scanner.Text()
				resLock.Lock()

				if len(models.HeadersList) > currentLineIndex {
					parseHeader(email, line, &currentLineIndex)
				} else {
					email.Content += "\n" + line
				}
				resLock.Unlock()
			}

			// Check scanner error
			if scannerErr == nil {
				// No error, return successfully
				result.EmailsList = append(result.EmailsList, email)
				result.EmailQty++
				return
			}

			// If error is related to token too long, double the buffer size and retry
			if strings.Contains(scannerErr.Error(), "token too long") {
				maxBufferSize *= 2
				continue
			}

			// If it's another error, return the error
			fmt.Println("Error:", scannerErr)
			return
		}
	}()
}
