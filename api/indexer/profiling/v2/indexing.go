package v2

import (
	"api/models"
	"api/src"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func IndexData(filesPath string) {
	fmt.Printf("=================\n")
	filename := "../test_files/allen-p/inbox/1." // Normal File
	//filename := "../test_files/reitmeyer-j/inbox/7." // MultiLine File
	email, err := parseEmail(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert email struct to JSON for demonstration purposes
	emailJSON, err := json.MarshalIndent(email, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(emailJSON))
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
	//originalMessages := []models.OriginalMessage{}

	scanner := bufio.NewScanner(file)
	//currentMessage := &models.OriginalMessage{}

	totalHeadersItems := len(models.HeadersList)
	currentLineIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		if totalHeadersItems > currentLineIndex {
			parseHeader(email, line, &currentLineIndex)
		} else if !strings.Contains(line, models.OriginalMessageSeparator) {
			// Content starts one line after header ends
			email.Content += "\n" + line
		} else {
			// If joins here is because it finds the message separator
		}
	}

	//if currentMessage != nil {
	//	// Add the last original message
	//	originalMessages = append(originalMessages, *currentMessage)
	//}

	//email.OriginalMessage = &originalMessages

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return email, nil
}

//func parseEmail(filename string) (*models.Email, error) {
//	file, err := os.Open(filename)
//	if err != nil {
//		return nil, err
//	}
//	defer file.Close()
//
//	email := &models.Email{}
//	originalMessages := []models.OriginalMessage{}
//
//	scanner := bufio.NewScanner(file)
//	currentMessage := &models.OriginalMessage{}
//
//	for scanner.Scan() {
//		line := scanner.Text()
//
//		// Extracting header information
//		if strings.HasPrefix(line, "Message-ID: ") {
//			email.MessageID = strings.TrimSpace(strings.TrimPrefix(line, "Message-ID: "))
//		} else if strings.HasPrefix(line, "Date: ") {
//			email.Date = strings.TrimSpace(strings.TrimPrefix(line, "Date: "))
//		} else if strings.HasPrefix(line, "From: ") {
//			email.From = strings.TrimSpace(strings.TrimPrefix(line, "From: "))
//		} else if strings.HasPrefix(line, "To: ") {
//			email.To = strings.TrimSpace(strings.TrimPrefix(line, "To: "))
//		} else if strings.HasPrefix(line, "Subject: ") {
//			email.Subject = strings.TrimSpace(strings.TrimPrefix(line, "Subject: "))
//		} else if strings.HasPrefix(line, "Mime-Version: ") {
//			email.MimeVersion = strings.TrimSpace(strings.TrimPrefix(line, "Mime-Version: "))
//		} else if strings.HasPrefix(line, "Content-Type: ") {
//			email.ContentType = strings.TrimSpace(strings.TrimPrefix(line, "Content-Type: "))
//		} else if strings.HasPrefix(line, "Content-Transfer-Encoding: ") {
//			email.ContentTransferEncoding = strings.TrimSpace(strings.TrimPrefix(line, "Content-Transfer-Encoding: "))
//		} else if strings.HasPrefix(line, "X-From: ") {
//			email.XFrom = strings.TrimSpace(strings.TrimPrefix(line, "X-From: "))
//		} else if strings.HasPrefix(line, "X-To: ") {
//			email.XTo = strings.TrimSpace(strings.TrimPrefix(line, "X-To: "))
//		} else if strings.HasPrefix(line, "X-cc: ") {
//			email.Xcc = strings.TrimSpace(strings.TrimPrefix(line, "X-cc: "))
//		} else if strings.HasPrefix(line, "X-bcc: ") {
//			email.Xbcc = strings.TrimSpace(strings.TrimPrefix(line, "X-bcc: "))
//		} else if strings.HasPrefix(line, "X-Folder: ") {
//			email.XFolder = strings.TrimSpace(strings.TrimPrefix(line, "X-Folder: "))
//		} else if strings.HasPrefix(line, "X-Origin: ") {
//			email.XOrigin = strings.TrimSpace(strings.TrimPrefix(line, "X-Origin: "))
//		} else if strings.HasPrefix(line, "X-FileName: ") {
//			email.XFileName = strings.TrimSpace(strings.TrimPrefix(line, "X-FileName: "))
//		} else if line == "" && currentMessage != nil {
//			// Blank line indicates the end of headers and start of content
//			email.Content = strings.TrimSpace(email.Content)
//			originalMessages = append(originalMessages, *currentMessage)
//			currentMessage = nil
//		} else if currentMessage != nil {
//			// Inside the content of the original message
//			currentMessage.Content += line + "\n"
//		}
//	}
//
//	if currentMessage != nil {
//		// Add the last original message
//		originalMessages = append(originalMessages, *currentMessage)
//	}
//
//	email.OriginalMessage = &originalMessages
//
//	if err := scanner.Err(); err != nil {
//		return nil, err
//	}
//
//	return email, nil
//}
